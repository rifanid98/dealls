package mocks

import (
	"dealls/core"
	"dealls/core/v1/entity"
)

type AuthUsecaseMock struct {
	ChangePassword  *core.CustomError
	IsActiveToken   *core.CustomError
	Login           *entity.Jwt
	LoginErr        *core.CustomError
	RefreshToken    *entity.Jwt
	RefreshTokenErr *core.CustomError
	Register        *entity.Account
	RegisterErr     *core.CustomError
	RevokeToken     *core.CustomError
}
