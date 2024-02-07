package mocks

import "dealls/core"

type XenditApicallMock struct {
	QRCreate    map[string]interface{}
	QRCreateErr *core.CustomError
	QrCheck     map[string]interface{}
	QrCheckErr  *core.CustomError
}
