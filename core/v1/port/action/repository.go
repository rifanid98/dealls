package action

import (
	"dealls/core"
	"dealls/core/v1/entity"
)

//go:generate mockery --name ActionRepository --filename action_repository.go --output ./mocks
type ActionRepository interface {
	InsertAction(ic *core.InternalContext, action *entity.Action) (*entity.Action, *core.CustomError)
	UpdateAction(ic *core.InternalContext, action *entity.Action) (*entity.Action, *core.CustomError)
	FindActionByTargetId(ic *core.InternalContext, targetId string) (*entity.Action, *core.CustomError)
}
