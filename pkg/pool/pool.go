package pool

import (
	"fmt"
	"github.com/AzusaChino/ficus/service/logging_service"
	"github.com/panjf2000/ants/v2"
	"log"
	"runtime"
	"time"
)

// Pool only for read file and send to kafka
var Pool *ants.PoolWithFunc

func Setup() {
	var err error
	size := runtime.NumCPU()

	Pool, err = ants.NewPoolWithFunc(size,
		logging_service.AsyncSend,
		ants.WithExpiryDuration(100*time.Second),
		ants.WithPanicHandler(func(i interface{}) {
			log.Fatal(i)
		}),
		ants.WithLogger(log.Default()))
	if err != nil {
		panic(fmt.Errorf("error when setup ants pool: %v", err))
	}
}

func Close() {
	if Pool != nil {
		Pool.Release()
	}
}
