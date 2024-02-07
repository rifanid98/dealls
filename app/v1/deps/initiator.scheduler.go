package deps

import "dealls/infrastructure/v1/scheduler/cron"

func (d *dependency) initScheduler() {
	d.initCronScheduler()
}

func (d *dependency) initCronScheduler() {
	d.base.Schlr = cron.New(d.usecase.AccountUsecase)
}
