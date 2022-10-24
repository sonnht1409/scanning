package workers

import (
	"context"

	"github.com/go-redis/redis/v9"

	"github.com/sonnht1409/scanning/service/common"
	"github.com/sonnht1409/scanning/service/config"
	"github.com/sonnht1409/scanning/service/connectors/cache"
)

type Worker struct {
	cache *redis.Client
}

func NewServiceWorker() Worker {
	return Worker{cache: cache.NewCache()}
}

func (w Worker) Subscribe() {
	log := common.GetLogger("Worker", config.Values.Env)
	subscriber := w.cache.Subscribe(context.Background(), config.Values.Worker.Topic)
	for {
		mess, err := subscriber.ReceiveMessage(context.Background())
		if err != nil {
			log.Errorf("ReceiveMessage Error :s \n", err.Error())
		}
		log.Infof("ReceiveMessage %+v", mess.Payload)
	}
}
