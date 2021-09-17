package kafka

import (
	"github.com/AzusaChino/ficus/pkg/conf"
	"github.com/Shopify/sarama"
	"log"
	"time"
)

var dataCollector sarama.AsyncProducer

func Setup() {
	endpoints := conf.KafkaConfig.Locations
	if len(endpoints) < 1 {
		return
	}
	var err error
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.RequiredAcks = sarama.WaitForLocal
	kafkaConfig.Producer.Compression = sarama.CompressionGZIP
	kafkaConfig.Producer.Flush.Frequency = 500 * time.Millisecond

	client, err := sarama.NewClient(endpoints, kafkaConfig)
	if err != nil {
		log.Fatalf("failed to setup kafka client, error: %v\n", err)
	}
	dataCollector, err = sarama.NewAsyncProducerFromClient(client)
	log.Println("kafka client initialized...")
}

func Close() {
	if err := dataCollector.Close(); err != nil {
		log.Println("Failed to shut down data collector cleanly", err)
	}
}

func SendMessage(topic string, key string, value string) {
	dataCollector.Input() <- &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(value),
	}
}
