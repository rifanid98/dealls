package account

import (
	"fmt"
	"sync"
	"time"

	"dealls/core"
	"dealls/core/v1/entity"
	"dealls/core/v1/port/adapter"
	"dealls/core/v1/usecase/common"
	"dealls/pkg/helper"
	"dealls/pkg/util"

	portAccount "dealls/core/v1/port/account"
	portAction "dealls/core/v1/port/action"
	portAuth "dealls/core/v1/port/auth"
	portCache "dealls/core/v1/port/cache"
	portCommon "dealls/core/v1/port/common"
)

var log = util.NewLogger()

type accountUsecaseImpl struct {
	authUsecase       portAuth.AuthUsecase
	accountRepository portAccount.AccountRepository
	actionRepository  portAction.ActionRepository
	cacheRepository   portCache.CacheRepository
	adapterXendit     adapter.XenditApiCall
	transaction       portCommon.Transaction
}

func NewAccountUsecase(
	authUsecase portAuth.AuthUsecase,
	accountRepository portAccount.AccountRepository,
	actionRepository portAction.ActionRepository,
	cacheRepository portCache.CacheRepository,
	adapterXendit adapter.XenditApiCall,
	transaction portCommon.Transaction,
) portAccount.AccountUsecase {
	return &accountUsecaseImpl{
		authUsecase:       authUsecase,
		accountRepository: accountRepository,
		actionRepository:  actionRepository,
		cacheRepository:   cacheRepository,
		adapterXendit:     adapterXendit,
		transaction:       transaction,
	}
}

type SwipeQuota struct {
	Profiles map[string]any `json:"profiles"`
	Expired  int64          `json:"expired"`
}

func (uc *accountUsecaseImpl) AccountAction(ic *core.InternalContext, accountId, targetId string, action int) *core.CustomError {
	redisKey := fmt.Sprintf("swap_quota::%v", accountId)
	get, cerr := uc.cacheRepository.Get(ic, redisKey)
	if cerr != nil {
		return cerr
	}

	// set expired for today until minute 59:59
	loc, _ := time.LoadLocation("Asia/Jakarta")
	year, month, day := time.Now().In(loc).AddDate(0, 0, 1).Date()
	expired := time.Date(year, month, day, 0, 0, 0, 0, loc).Add(-time.Second * 1)

	var quota = SwipeQuota{
		Profiles: make(map[string]any),
		Expired:  expired.Unix(),
	}
	if get != nil {
		cerr = helper.StringToStruct(*get, &quota)
		if cerr != nil {
			log.Error(ic.ToContext(), "failed helper.StringToStruct(*get, &quota)", cerr)
			return cerr
		}
	}

	// reset quota after 1 day
	exp := time.Unix(quota.Expired, 0)
	if time.Now().Sub(exp) > 0 {
		quota = SwipeQuota{
			Profiles: make(map[string]any),
			Expired:  expired.Unix(),
		}
	}

	var claims entity.JwtClaim
	ctxData := ic.GetData()
	c := ctxData["claims"]
	cerr = helper.StringToStruct(helper.DataToString(c), &claims)
	if cerr != nil {
		log.Error(ic.ToContext(), "failed helper.StringToStruct(helper.DataToString(c), &claims)", cerr)
		return &core.CustomError{
			Code:    core.BAD_REQUEST,
			Message: "quota limit reached",
		}
	}

	//if claims.Verified == core.ACCOUNT_UNVERIFIED && len(quota.Profiles) >= 10 {
	if claims.Verified == core.ACCOUNT_UNVERIFIED && len(quota.Profiles) >= 1 {
		return &core.CustomError{
			Code:    core.BAD_REQUEST,
			Message: "quota limit reached",
		}
	}
	quota.Profiles[targetId] = action

	var acted *entity.Action
	var now = time.Now()

	acted, cerr = uc.actionRepository.FindActionByTargetId(ic, targetId)
	if cerr != nil {
		return cerr
	}

	tx, txCtx, cerr := uc.transaction.StartTransaction(ic)
	if cerr != nil {
		return cerr
	}

	if acted == nil {
		acted, cerr = uc.actionRepository.InsertAction(txCtx, &entity.Action{
			UserId:   accountId,
			TargetId: targetId,
			Action:   action,
			History: []entity.ActionHistory{
				{
					Action:    action,
					Timestamp: now,
				},
			},
			Created: now,
		})
		if cerr != nil {
			return cerr
		}
	}

	if acted != nil {
		if acted.Action != action {
			acted.History = append(acted.History, entity.ActionHistory{
				Action:    action,
				Timestamp: now,
			})

			acted, cerr = uc.actionRepository.UpdateAction(txCtx, acted)
			if cerr != nil {
				return cerr
			}
		}
	}

	duration := time.Hour * 24
	cerr = uc.cacheRepository.Set(ic, redisKey, helper.DataToString(quota), &duration)
	if cerr != nil {
		log.Error(ic.ToContext(), "failed uc.cacheRepository.Set(ic, redisKey, helper.DataToString(profileIds), &duration)", cerr.Error())
		return common.AbortTransaction(txCtx, tx, cerr)
	}

	return common.CommitTransaction(txCtx, tx)
}

func (uc *accountUsecaseImpl) AccountList(ic *core.InternalContext, accountId string, meta map[string]any) ([]entity.Account, int32, *core.CustomError) {
	redisKey := fmt.Sprintf("swap_quota::%v", accountId)
	get, cerr := uc.cacheRepository.Get(ic, redisKey)
	if cerr != nil {
		return nil, 0, cerr
	}

	var quota = SwipeQuota{}
	if get != nil {
		cerr = helper.StringToStruct(*get, &quota)
		if cerr != nil {
			log.Error(ic.ToContext(), "failed helper.StringToStruct(*get, &quota)", cerr)
			return nil, 0, cerr
		}
	}

	var profileIds = []string{accountId}
	for id, _ := range quota.Profiles {
		profileIds = append(profileIds, id)
	}

	accounts, total, cerr := uc.accountRepository.GetAccountsExclude(ic, profileIds, map[string]any{
		"page":  helper.DataToInt(meta["page"]),
		"limit": helper.DataToInt(meta["limit"]),
	})
	if cerr != nil {
		return nil, 0, cerr
	}

	return accounts, total, nil
}

func (uc *accountUsecaseImpl) AccountGet(ic *core.InternalContext, accountId string) (*entity.Account, *core.CustomError) {
	return uc.accountRepository.FindAccountById(ic, accountId)
}

func (uc *accountUsecaseImpl) AccountActivate(ic *core.InternalContext, accountId string) (map[string]any, *core.CustomError) {
	account, cerr := uc.accountRepository.FindAccountById(ic, accountId)
	if cerr != nil {
		return nil, cerr
	}
	if account == nil {
		return nil, &core.CustomError{
			Code:    core.UNPROCESSABLE_ENTITY,
			Message: "account not found",
		}
	}

	if account.Verified == core.ACCOUNT_VERIFIED && account.Metadata != nil {
		return nil, &core.CustomError{
			Code:    core.UNPROCESSABLE_ENTITY,
			Message: "account already activated",
		}
	}

	result, cerr := uc.adapterXendit.QRCreate(ic, map[string]any{
		"account_id": account.Id,
		"amount":     10000,
	})
	if cerr != nil {
		return nil, cerr
	}

	account.Metadata = result
	account, cerr = uc.accountRepository.UpdateAccount(ic, account)
	if cerr != nil {
		return nil, cerr
	}

	return result, nil
}

func (uc *accountUsecaseImpl) AccountActivationCheck(ic *core.InternalContext) *core.CustomError {
	log.Info(ic.ToContext(), "AccountActivationCheck running...")

	accounts, cerr := uc.accountRepository.FindAccountsActivation(ic)
	if cerr != nil {
		return cerr
	}
	if accounts == nil {
		return &core.CustomError{
			Code:    core.UNPROCESSABLE_ENTITY,
			Message: "account not found",
		}
	}

	wg := sync.WaitGroup{}
	for i, account := range accounts {
		wg.Add(1)

		go func() {
			defer wg.Done()

			transactionId := helper.DataToString(account.Metadata["id"])

			result, cerr := uc.adapterXendit.QrCheck(ic, map[string]any{
				"id_qr": transactionId,
			})
			if cerr != nil {
				log.Error(ic.ToContext(), fmt.Sprintf("failed  uc.adapterXendit.QrCheck::%v::%v", account.Id, account.Metadata["id"]), cerr)
				return
			}

			// [START BYPASS STATUS XENDIT]
			// xendit account ini hanya bisa sampai testing saja (mode sandbox)
			// dikarenakan belum melengkapi data aktivasi, account xendit ini
			// hanya akan menampilkan data dummy, qr code yang didapatkan pun dummy
			// sehingga transaksi qr yang sudah dibuat tidak bisa dibayar.
			// jika tidak bisa dibayar, maka status akan selalu ACTIVE.
			result["status"] = core.XENDIT_STATUS_SUCCEEDED
			// [END BYPASS STATUS XENDIT]

			if helper.DataToString(result["status"]) == core.XENDIT_STATUS_SUCCEEDED {
				accounts[i].Metadata = result
				accounts[i].Verified = core.ACCOUNT_VERIFIED
				accounts[i].VerifiedDate = time.Now()
				_, cerr = uc.accountRepository.UpdateAccount(ic, &accounts[i])
				if cerr != nil {
					log.Error(ic.ToContext(), fmt.Sprintf("failed uc.accountRepository.UpdateAccount(ic, &accounts[i])::%v::%v", account.Id, account.Metadata["id"]), cerr)
					return
				}

				cerr = uc.authUsecase.RevokeToken(ic, account.Id)
				if cerr != nil {
					log.Error(ic.ToContext(), fmt.Sprintf("failed uc.authUsecase.RevokeToken(ic, account.Id)::%v::%v", account.Id, account.Metadata["id"]), cerr)
					return
				}
			}
		}()
	}

	wg.Wait()

	return nil
}
