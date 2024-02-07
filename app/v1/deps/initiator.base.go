package deps

import (
	"dealls/config"
	"dealls/infrastructure/v1/broker/gcp"
	"dealls/infrastructure/v1/persistence/mongodb"
	"dealls/infrastructure/v1/persistence/redisdb"
	"dealls/pkg/api"
	"dealls/pkg/util"
)

func (d *dependency) initBase() {
	d.initConfig()
	d.initMongodb()
	d.initCache()
	d.initHttpClient()
	d.initRetrier()

	// [NEED GCP CREDENTIALS ACCESS]
	//d.initGcpBroker()
}

func (d *dependency) initConfig() {
	d.base.Cfg = config.GetConfig()
}

func (d *dependency) initMongodb() {
	mdb, mdbc, err := mongodb.New(d.base.Cfg.MongoDb)
	if err != nil {
		panic(err)
	}
	tx := mongodb.NewTransaction(mdbc)
	d.base.Mdb = mdb
	d.base.Mdbc = mdbc
	d.base.Mdbt = tx
}

func (d *dependency) initCache() {
	rdb, err := redisdb.New(d.base.Cfg.Redis)
	if err != nil {
		panic(err)
	}
	d.base.Rdbc = rdb
}

func (d *dependency) initHttpClient() {
	d.base.Httpc = api.NewHttpClient()
}

func (d *dependency) initGcpBroker() {
	gcpc, err := gcp.New(d.base.Cfg.GcpPubsub)
	if err != nil {
		panic(err)
	}
	d.base.Gcpc = gcpc
}

func (d *dependency) initRetrier() {
	d.base.Rtr = util.NewRetrier()
}
