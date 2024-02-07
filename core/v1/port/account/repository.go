package account

import (
	"dealls/core"
	"dealls/core/v1/entity"
)

//go:generate mockery --name AccountRepository --filename account_repository.go --output ./mocks
type AccountRepository interface {
	InsertAccount(ic *core.InternalContext, account *entity.Account) (*entity.Account, *core.CustomError)
	FindAccountByUsername(ic *core.InternalContext, username string) (*entity.Account, *core.CustomError)
	FindAccountByEmail(ic *core.InternalContext, email string) (*entity.Account, *core.CustomError)
	GetAccountsExclude(ic *core.InternalContext, profileIds []string, meta map[string]any) ([]entity.Account, int32, *core.CustomError)
	FindAccountById(ic *core.InternalContext, accountId string) (*entity.Account, *core.CustomError)
	FindAccountsActivation(ic *core.InternalContext) ([]entity.Account, *core.CustomError)
	UpdateAccount(ic *core.InternalContext, account *entity.Account) (*entity.Account, *core.CustomError)
}
