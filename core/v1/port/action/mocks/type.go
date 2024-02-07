package mocks

import (
	"dealls/core"
	"dealls/core/v1/entity"
)

type ActionRepositoryMock struct {
	FindActionByTargetId    *entity.Action
	FindActionByTargetIdErr *core.CustomError
	InsertAction            *entity.Action
	InsertActionErr         *core.CustomError
	UpdateAction            *entity.Action
	UpdateActionErr         *core.CustomError
}
