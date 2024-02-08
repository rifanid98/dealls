package account

import (
	"dealls/core/v1/entity"
	"dealls/pkg/helper"
	"github.com/google/uuid"
	"reflect"
	"testing"
	"time"

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
			name: "swipe action must be success",
			fields: fields{
				authUsecase:       new(authMocks.AuthUsecase),
				accountRepository: new(accountMocks.AccountRepository),
				actionRepository:  new(actionMocks.ActionRepository),
				cacheRepository:   new(cacheMocks.CacheRepository),
				adapterXendit:     new(adapterMocks.XenditApiCall),
				transaction:       new(commonMocks.Transaction),
			},
			mocks: mocks{
				authUsecase:       authMocks.AuthUsecaseMock{},
				accountRepository: accountMocks.AccountRepositoryMock{},
				actionRepository:  actionMocks.ActionRepositoryMock{},
				cacheRepository: cacheMocks.CacheRepositoryMock{
					Delete: nil,
					Get: helper.DataToString(SwipeQuota{
						Profiles: map[string]any{
							"1": 1,
						},
						Expired: 0,
					}),
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
				ic: core.NewInternalContext(uuid.NewString()).InjectData(map[string]any{
					"claims": entity.JwtClaim{
						Id:       "",
						Verified: 0,
					},
				}),
				accountId: uuid.NewString(),
				targetId:  uuid.NewString(),
				action:    1,
			},
			want: nil,
		},
		{
			name: "swipe action must be success, but profile already acted",
			fields: fields{
				authUsecase:       new(authMocks.AuthUsecase),
				accountRepository: new(accountMocks.AccountRepository),
				actionRepository:  new(actionMocks.ActionRepository),
				cacheRepository:   new(cacheMocks.CacheRepository),
				adapterXendit:     new(adapterMocks.XenditApiCall),
				transaction:       new(commonMocks.Transaction),
			},
			mocks: mocks{
				authUsecase:       authMocks.AuthUsecaseMock{},
				accountRepository: accountMocks.AccountRepositoryMock{},
				actionRepository: actionMocks.ActionRepositoryMock{
					FindActionByTargetId: &entity.Action{},
				},
				cacheRepository: cacheMocks.CacheRepositoryMock{
					Get: helper.DataToString(SwipeQuota{
						Profiles: map[string]any{
							"1": 1,
						},
						Expired: 0,
					}),
				},
				adapterXendit: adapterMocks.XenditApicallMock{},
				transaction: commonMocks.TransactionMock{
					StartTransactionTx:    new(commonMocks.Transaction),
					StartTransactionTxCtx: core.NewInternalContext(uuid.NewString()),
				},
			},
			args: args{
				ic: core.NewInternalContext(uuid.NewString()).InjectData(map[string]any{
					"claims": entity.JwtClaim{
						Id:       "",
						Verified: 0,
					},
				}),
				accountId: uuid.NewString(),
				targetId:  uuid.NewString(),
				action:    1,
			},
			want: nil,
		},
		{
			name: "swipe limit reached",
			fields: fields{
				authUsecase:       new(authMocks.AuthUsecase),
				accountRepository: new(accountMocks.AccountRepository),
				actionRepository:  new(actionMocks.ActionRepository),
				cacheRepository:   new(cacheMocks.CacheRepository),
				adapterXendit:     new(adapterMocks.XenditApiCall),
				transaction:       new(commonMocks.Transaction),
			},
			mocks: mocks{
				authUsecase:       authMocks.AuthUsecaseMock{},
				accountRepository: accountMocks.AccountRepositoryMock{},
				actionRepository:  actionMocks.ActionRepositoryMock{},
				cacheRepository: cacheMocks.CacheRepositoryMock{
					Delete: nil,
					Get: helper.DataToString(SwipeQuota{
						Profiles: map[string]any{
							"1":  1,
							"2":  2,
							"3":  3,
							"4":  4,
							"5":  5,
							"6":  6,
							"7":  7,
							"8":  8,
							"9":  9,
							"10": 10,
						},
						Expired: time.Now().Add(time.Hour).Unix(),
					}),
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
				ic: core.NewInternalContext(uuid.NewString()).InjectData(map[string]any{
					"claims": entity.JwtClaim{
						Id:       "",
						Verified: 0,
					},
				}),
				accountId: uuid.NewString(),
				targetId:  uuid.NewString(),
				action:    1,
			},
			want: &core.CustomError{
				Code:    core.BAD_REQUEST,
				Message: "quota limit reached",
			},
		},
		{
			name: "swipe limit reached, but premium user got unlimited swipe",
			fields: fields{
				authUsecase:       new(authMocks.AuthUsecase),
				accountRepository: new(accountMocks.AccountRepository),
				actionRepository:  new(actionMocks.ActionRepository),
				cacheRepository:   new(cacheMocks.CacheRepository),
				adapterXendit:     new(adapterMocks.XenditApiCall),
				transaction:       new(commonMocks.Transaction),
			},
			mocks: mocks{
				authUsecase:       authMocks.AuthUsecaseMock{},
				accountRepository: accountMocks.AccountRepositoryMock{},
				actionRepository:  actionMocks.ActionRepositoryMock{},
				cacheRepository: cacheMocks.CacheRepositoryMock{
					Delete: nil,
					Get: helper.DataToString(SwipeQuota{
						Profiles: map[string]any{
							"1":  1,
							"2":  2,
							"3":  3,
							"4":  4,
							"5":  5,
							"6":  6,
							"7":  7,
							"8":  8,
							"9":  9,
							"10": 10,
						},
						Expired: time.Now().Add(time.Hour).Unix(),
					}),
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
				ic: core.NewInternalContext(uuid.NewString()).InjectData(map[string]any{
					"claims": entity.JwtClaim{
						Id:       "",
						Verified: 1,
					},
				}),
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

func Test_accountUsecaseImpl_AccountList(t *testing.T) {
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
		meta      map[string]any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		mocks  mocks
		want   []entity.Account
		want1  int32
		want2  *core.CustomError
	}{
		{
			name: "get list must be success",
			fields: fields{
				authUsecase:       new(authMocks.AuthUsecase),
				accountRepository: new(accountMocks.AccountRepository),
				actionRepository:  new(actionMocks.ActionRepository),
				cacheRepository:   new(cacheMocks.CacheRepository),
				adapterXendit:     new(adapterMocks.XenditApiCall),
				transaction:       new(commonMocks.Transaction),
			},
			args: args{
				ic:        core.NewInternalContext(uuid.NewString()),
				accountId: uuid.NewString(),
				meta: map[string]any{
					"page":  1,
					"limit": 10,
				},
			},
			mocks: mocks{
				cacheRepository: cacheMocks.CacheRepositoryMock{
					Get: helper.DataToString(SwipeQuota{
						Profiles: map[string]any{
							"1": 1,
						},
						Expired: time.Now().Add(time.Hour).Unix(),
					}),
				},
				accountRepository: accountMocks.AccountRepositoryMock{
					GetAccountsExclude:      []entity.Account{},
					GetAccountsExcludeTotal: 0,
					GetAccountsExcludeErr:   nil,
				},
			},
			want:  []entity.Account{},
			want1: 0,
			want2: nil,
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

			cacheRepository.On("Get", mock.Anything, mock.Anything).Return(mocks.cacheRepository.Get, mocks.cacheRepository.GetErr)
			accountRepository.On("GetAccountsExclude", mock.Anything, mock.Anything, mock.Anything).Return(mocks.accountRepository.GetAccountsExclude, mocks.accountRepository.GetAccountsExcludeTotal, mocks.accountRepository.GetAccountsExcludeErr)
			got, got1, got2 := uc.AccountList(tt.args.ic, tt.args.accountId, tt.args.meta)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountList() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("AccountList() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("AccountList() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_accountUsecaseImpl_AccountGet(t *testing.T) {
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
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		mocks  mocks
		want   *entity.Account
		want1  *core.CustomError
	}{
		{
			name: "get account detail must be success",
			fields: fields{
				authUsecase:       new(authMocks.AuthUsecase),
				accountRepository: new(accountMocks.AccountRepository),
				actionRepository:  new(actionMocks.ActionRepository),
				cacheRepository:   new(cacheMocks.CacheRepository),
				adapterXendit:     new(adapterMocks.XenditApiCall),
				transaction:       new(commonMocks.Transaction),
			},
			args: args{
				ic:        core.NewInternalContext(uuid.NewString()),
				accountId: uuid.NewString(),
			},
			mocks: mocks{
				accountRepository: accountMocks.AccountRepositoryMock{
					FindAccountById:    &entity.Account{},
					FindAccountByIdErr: nil,
				},
			},
			want:  &entity.Account{},
			want1: nil,
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

			accountRepository.On("FindAccountById", mock.Anything, mock.Anything).Return(mocks.accountRepository.FindAccountById, mocks.accountRepository.FindAccountByIdErr)

			got, got1 := uc.AccountGet(tt.args.ic, tt.args.accountId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountGet() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("AccountGet() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_accountUsecaseImpl_AccountActivate(t *testing.T) {
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
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		mocks  mocks
		want   map[string]any
		want1  *core.CustomError
	}{
		{
			name: "activate account premium must be success",
			fields: fields{
				authUsecase:       new(authMocks.AuthUsecase),
				accountRepository: new(accountMocks.AccountRepository),
				actionRepository:  new(actionMocks.ActionRepository),
				cacheRepository:   new(cacheMocks.CacheRepository),
				adapterXendit:     new(adapterMocks.XenditApiCall),
				transaction:       new(commonMocks.Transaction),
			},
			args: args{
				ic:        core.NewInternalContext(uuid.NewString()),
				accountId: uuid.NewString(),
			},
			mocks: mocks{
				accountRepository: accountMocks.AccountRepositoryMock{
					FindAccountById:    &entity.Account{},
					FindAccountByIdErr: nil,
					UpdateAccount:      &entity.Account{},
					UpdateAccountErr:   nil,
				},
				adapterXendit: adapterMocks.XenditApicallMock{
					QRCreate: map[string]interface{}{
						"qr_code": "qr code",
					},
					QRCreateErr: nil,
				},
			},
			want: map[string]interface{}{
				"qr_code": "qr code",
			},
			want1: nil,
		},
		{
			name: "activate account premium, but already activated",
			fields: fields{
				authUsecase:       new(authMocks.AuthUsecase),
				accountRepository: new(accountMocks.AccountRepository),
				actionRepository:  new(actionMocks.ActionRepository),
				cacheRepository:   new(cacheMocks.CacheRepository),
				adapterXendit:     new(adapterMocks.XenditApiCall),
				transaction:       new(commonMocks.Transaction),
			},
			args: args{
				ic:        core.NewInternalContext(uuid.NewString()),
				accountId: uuid.NewString(),
			},
			mocks: mocks{
				accountRepository: accountMocks.AccountRepositoryMock{
					FindAccountById: &entity.Account{
						Verified: 1,
						Metadata: map[string]any{},
					},
					FindAccountByIdErr: nil,
					UpdateAccount:      &entity.Account{},
					UpdateAccountErr:   nil,
				},
				adapterXendit: adapterMocks.XenditApicallMock{
					QRCreate: map[string]interface{}{
						"qr_code": "qr code",
					},
					QRCreateErr: nil,
				},
			},
			want: nil,
			want1: &core.CustomError{
				Code:    core.UNPROCESSABLE_ENTITY,
				Message: "account already activated",
			},
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

			accountRepository.On("FindAccountById", mock.Anything, mock.Anything).Return(mocks.accountRepository.FindAccountById, mocks.accountRepository.FindAccountByIdErr)
			adapterXendit.On("QRCreate", mock.Anything, mock.Anything).Return(mocks.adapterXendit.QRCreate, mocks.adapterXendit.QRCreateErr)
			accountRepository.On("UpdateAccount", mock.Anything, mock.Anything).Return(mocks.accountRepository.UpdateAccount, mocks.accountRepository.UpdateAccountErr)

			got, got1 := uc.AccountActivate(tt.args.ic, tt.args.accountId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountActivate() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("AccountActivate() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_accountUsecaseImpl_AccountActivationCheck(t *testing.T) {
	type fields struct {
		authUsecase       *authMocks.AuthUsecase
		accountRepository *accountMocks.AccountRepository
		actionRepository  *actionMocks.ActionRepository
		cacheRepository   *cacheMocks.CacheRepository
		adapterXendit     *adapterMocks.XenditApiCall
		transaction       *commonMocks.Transaction
	}
	type args struct {
		ic *core.InternalContext
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		mocks  mocks
		want   *core.CustomError
	}{
		{
			name: "account activation check must be successs",
			fields: fields{
				authUsecase:       new(authMocks.AuthUsecase),
				accountRepository: new(accountMocks.AccountRepository),
				actionRepository:  new(actionMocks.ActionRepository),
				cacheRepository:   new(cacheMocks.CacheRepository),
				adapterXendit:     new(adapterMocks.XenditApiCall),
				transaction:       new(commonMocks.Transaction),
			},
			args: args{
				ic: core.NewInternalContext(uuid.NewString()),
			},
			mocks: mocks{
				accountRepository: accountMocks.AccountRepositoryMock{
					FindAccountsActivation: []entity.Account{
						{
							Metadata: map[string]any{
								"id": uuid.NewString(),
							},
						},
					},
					FindAccountsActivationErr: nil,
					UpdateAccount:             &entity.Account{},
					UpdateAccountErr:          nil,
				},
				adapterXendit: adapterMocks.XenditApicallMock{
					QrCheck: map[string]interface{}{
						"status": core.XENDIT_STATUS_SUCCEEDED,
					},
					QrCheckErr: nil,
				},
				authUsecase: authMocks.AuthUsecaseMock{
					RevokeToken: nil,
				},
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

			accountRepository.On("FindAccountsActivation", mock.Anything).Return(mocks.accountRepository.FindAccountsActivation, mocks.accountRepository.FindAccountsActivationErr)
			adapterXendit.On("QRCheck", mock.Anything, mock.Anything).Return(mocks.adapterXendit.QrCheck, mocks.adapterXendit.QrCheckErr)
			accountRepository.On("UpdateAccount", mock.Anything, mock.Anything).Return(mocks.accountRepository.UpdateAccount, mocks.accountRepository.UpdateAccountErr)
			authUsecase.On("RevokeToken", mock.Anything, mock.Anything).Return(mocks.authUsecase.RevokeToken)

			if got := uc.AccountActivationCheck(tt.args.ic); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountActivationCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}
