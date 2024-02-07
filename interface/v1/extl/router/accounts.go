package router

import (
	"github.com/labstack/echo/v4"

	"dealls/app/v1/deps"
	"dealls/interface/v1/extl/handler/accounts"
	"dealls/interface/v1/general/middleware"
)

func accountRouter(
	e *echo.Group,
	deps deps.IDependency,
) {
	service := deps.GetServices()
	handler := accounts.New(service)
	cfg := deps.GetBase().Cfg

	accounts := e.Group("/accounts")
	accounts.Use(middleware.InternalContext)
	accounts.Use(middleware.JwtAccessTokenMiddleware(service.AuthUsecase, cfg.JwtSecretKeys...))
	accounts.Use(middleware.JwtAccessTokenParseMiddleware)
	accounts.GET("", handler.List)
	accounts.POST("/:id/action", handler.Action)
	accounts.GET("/:id", handler.Get)

	// TODO:
	// - check apakah sudah verified atau belum
	// - generate qr
	// - jika belum maka akan update field metadata dengan status pending, simpan juga id transaksi dari xendit + qr code
	// - return qr
	accounts.POST("/:id/activate", handler.Activate)
}
