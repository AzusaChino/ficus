package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type App struct {
	LogFileLocation string
}

var AppConfig = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerConfig = &Server{}

type Kafka struct {
	Host string
	Port int
}

var KafkaConfig *Kafka

func Setup() {
	var err error
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("conf")
	if err = viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error when reading config file: %w \n", err))
	}
	kafkaMap := viper.GetStringMap("kafka")
	KafkaConfig = &Kafka{
		Host: kafkaMap["host"].(string),
	}
}
