package deps

import (
	"dealls/infrastructure/v1/broker/gcp/adapter"
)

func (d *dependency) initBroker() {
	// [NEED GCP CREDENTIALS ACCESS]
	//d.initGcpPubsub()
}

func (d *dependency) initGcpPubsub() {
	d.broker.Pubsub = adapter.NewGcpPubsub(d.base.Gcpc, d.base.Cfg.GcpPubsub, nil)
}
