package mocks

import "dealls/core"

type PubsubMock struct {
	Publish   *core.CustomError
	Subscribe *core.CustomError
}
