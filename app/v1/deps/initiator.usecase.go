package deps

import (
	"dealls/core/v1/usecase/account"
	"dealls/core/v1/usecase/auth"
	"dealls/core/v1/usecase/subscriber"
)

func (d *dependency) initService() {
	d.initAuthUsecase()
	d.initAccountUsecase()
	d.initSubscriberUsecase()
}

func (d *dependency) initAuthUsecase() {
	d.usecase.AuthUsecase = auth.NewAuthUsecase(
		d.repo.AccountRepository,
		d.repo.CacheRepository,
		d.base.Cfg,
	)
}

func (d *dependency) initAccountUsecase() {
	d.usecase.AccountUsecase = account.NewAccountUsecase(
		d.usecase.AuthUsecase,
		d.repo.AccountRepository,
		d.repo.ActionRepository,
		d.repo.CacheRepository,
		d.apicall.XenditApiCall,
		d.base.Mdbt,
	)
}

func (d *dependency) initSubscriberUsecase() {
	d.usecase.SubscriberUsecase = subscriber.NewSubscriberUsecase()
}
