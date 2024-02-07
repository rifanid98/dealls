package mocks

import "dealls/core/v1/port/retrier"

type RetrierMock struct {
	Retry retrier.Effector
}
