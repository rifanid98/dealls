package account

import (
	"dealls/pkg/helper"
	"github.com/google/uuid"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"

	"dealls/core"

	accountMocks "dealls/core/v1/port/account/mocks"
	actionMocks "dealls/core/v1/port/action/mocks"
	adapterMocks "dealls/core/v1/port/adapter/mocks"
	authMocks "dealls/core/v1/port/auth/mocks"
	cacheMocks "dealls/core/v1/port/cache/mocks"
	commonMocks "dealls/core/v1/port/common/mocks"
)

type mocks struct {
	authUsecase       authMocks.AuthUsecaseMock
	accountRepository accountMocks.AccountRepositoryMock
	actionRepository  actionMocks.ActionRepositoryMock
	cacheRepository   cacheMocks.CacheRepositoryMock
	adapterXendit     adapterMocks.XenditApicallMock
	transaction       commonMocks.TransactionMock
}

func Test_accountUsecaseImpl_AccountAction(t *testing.T) {
	redisGet := helper.DataToString(SwipeQuota{
		Profiles: map[string]any{
			"1": 1,
		},
		Expired: 0,
	})
	type fields struct {
		authUsecase       *authMocks.AuthUsecase
		accountRepository *accountMocks.AccountRepository
		actionRepository  *actionMocks.ActionRepository
		cacheRepository   *cacheMocks.CacheRepository
		adapterXendit     *adapterMocks.XenditApiCall
		transaction       *commonMocks.Transaction
	}
	type args struct {
		ic        *core.InternalContext
		accountId string
		targetId  string
		action    int
	}
	tests := []struct {
		name   string
		fields fields
		mocks  mocks
		args   args
		want   *core.CustomError
	}{
		{
			name: "swipe left (pass) action must be success",
			fields: fields{
				authUsecase:       nil,
				accountRepository: nil,
				actionRepository:  nil,
				cacheRepository:   nil,
				adapterXendit:     nil,
				transaction:       nil,
			},
			mocks: mocks{
				authUsecase:       authMocks.AuthUsecaseMock{},
				accountRepository: accountMocks.AccountRepositoryMock{},
				actionRepository:  actionMocks.ActionRepositoryMock{},
				cacheRepository: cacheMocks.CacheRepositoryMock{
					Delete:  nil,
					Get:     &redisGet,
					GetErr:  nil,
					HSet:    nil,
					Publish: nil,
					Set:     nil,
				},
				adapterXendit: adapterMocks.XenditApicallMock{},
				transaction: commonMocks.TransactionMock{
					AbortTransaction:      nil,
					CommitTransaction:     nil,
					StartTransactionTx:    new(commonMocks.Transaction),
					StartTransactionTxCtx: core.NewInternalContext(uuid.NewString()),
					StartTransactionErr:   nil,
				},
			},
			args: args{
				ic:        core.NewInternalContext(uuid.NewString()),
				accountId: uuid.NewString(),
				targetId:  uuid.NewString(),
				action:    1,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			authUsecase := tt.fields.authUsecase
			accountRepository := tt.fields.accountRepository
			actionRepository := tt.fields.actionRepository
			cacheRepository := tt.fields.cacheRepository
			adapterXendit := tt.fields.adapterXendit
			transaction := tt.fields.transaction
			mocks := tt.mocks

			uc := &accountUsecaseImpl{
				authUsecase:       authUsecase,
				accountRepository: accountRepository,
				actionRepository:  actionRepository,
				cacheRepository:   cacheRepository,
				adapterXendit:     adapterXendit,
				transaction:       transaction,
			}

			mocks.transaction.StartTransactionTx.On("StartTransaction", mock.Anything).Return(mocks.transaction.StartTransactionTx, mocks.transaction.StartTransactionTxCtx, mocks.transaction.StartTransactionErr)
			mocks.transaction.StartTransactionTx.On("AbortTransaction", mock.Anything).Return(mocks.transaction.AbortTransaction)
			mocks.transaction.StartTransactionTx.On("CommitTransaction", mock.Anything).Return(mocks.transaction.CommitTransaction)

			cacheRepository.On("Get", mock.Anything, mock.Anything).Return(mocks.cacheRepository.Get, mocks.cacheRepository.GetErr)
			actionRepository.On("FindActionByTargetId", mock.Anything, mock.Anything).Return(mocks.actionRepository.FindActionByTargetId, mocks.actionRepository.FindActionByTargetIdErr)
			transaction.On("StartTransaction", mock.Anything).Return(mocks.transaction.StartTransactionTx, mocks.transaction.StartTransactionTxCtx, mocks.transaction.StartTransactionErr)
			actionRepository.On("InsertAction", mock.Anything, mock.Anything).Return(mocks.actionRepository.InsertAction, mocks.actionRepository.InsertActionErr)
			actionRepository.On("UpdateAction", mock.Anything, mock.Anything).Return(mocks.actionRepository.UpdateAction, mocks.actionRepository.UpdateActionErr)
			cacheRepository.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mocks.cacheRepository.Set)
			if got := uc.AccountAction(tt.args.ic, tt.args.accountId, tt.args.targetId, tt.args.action); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountAction() = %v, want %v", got, tt.want)
			}
		})
	}
}
