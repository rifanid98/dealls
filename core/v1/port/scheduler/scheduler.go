package scheduler

import (
	"dealls/core"
)

//go:generate mockery --name Scheduler --filename scheduler.go --output ./mocks
type Scheduler interface {
	Start(ic *core.InternalContext) *core.CustomError
}
