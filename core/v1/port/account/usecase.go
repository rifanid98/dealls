package account

import (
	"dealls/core"
	"dealls/core/v1/entity"
)

//go:generate mockery --name AccountUsecase --filename account_usecase.go --output ./mocks
type AccountUsecase interface {
	AccountAction(ic *core.InternalContext, accountId, targetId string, action int) *core.CustomError
	AccountList(ic *core.InternalContext, accountId string, meta map[string]any) ([]entity.Account, int32, *core.CustomError)
	AccountGet(ic *core.InternalContext, accountId string) (*entity.Account, *core.CustomError)
	AccountActivate(ic *core.InternalContext, accountId string) (map[string]any, *core.CustomError)
	AccountActivationCheck(ic *core.InternalContext) *core.CustomError
}
