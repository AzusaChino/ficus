package etcd

import (
	"github.com/AzusaChino/ficus/pkg/conf"
	"go.etcd.io/etcd/client/v3"
	"log"
)

var Client *clientv3.Client

func Setup() {
	endpoints := conf.EtcdConfig.EndPoints
	// skip initializing
	if len(endpoints) < 1 {
		return
	}
	var err error
	config := clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: conf.EtcdConfig.Timeout,
	}
	Client, err = clientv3.New(config)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("etcd client initialized...")
}

func Close() {
	if Client != nil {
		_ = Client.Close()
	}
}
