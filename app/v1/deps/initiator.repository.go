package deps

import (
	mdb "dealls/infrastructure/v1/persistence/mongodb/repository"
	rdb "dealls/infrastructure/v1/persistence/redisdb/repository"
)

func (d *dependency) initRepository() {
	d.initAccountRepository()
	d.initActionRepository()
	d.initCacheRepository()
}

func (d *dependency) initAccountRepository() {
	d.repo.AccountRepository = mdb.NewAccountRepository(d.base.Mdb, d.base.Cfg)
}

func (d *dependency) initActionRepository() {
	d.repo.ActionRepository = mdb.NewActionRepository(d.base.Mdb, d.base.Cfg)
}

func (d *dependency) initCacheRepository() {
	d.repo.CacheRepository = rdb.NewCacheRepository(d.base.Rdbc)
}
