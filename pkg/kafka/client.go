package kafka

import (
	"github.com/AzusaChino/ficus/pkg/conf"
	"github.com/Shopify/sarama"
)

var KafkaClient sarama.Client

func Setup() {
	var err error
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Producer.Partitioner = sarama.NewRoundRobinPartitioner

	KafkaClient,err = sarama.NewClient(conf.KafkaConfig.Locations, kafkaConfig)
	if err != nil {
		panic("failed to setup kafka client")
	}
	producer, err := sarama.New

}

func Close() {
	_ = KafkaClient.Close()
}