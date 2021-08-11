package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type App struct {
	RuntimeRootPath string
	LogFileLocation string
	LogFileSaveName string
	LogFileExt      string
	TimeFormat      string
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
	Locations []string
}

var KafkaConfig *Kafka

func Setup() {
	var err error
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath("conf")
	if err = viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error when reading config file: %w \n", err))
	}

	AppConfig = &App{
		RuntimeRootPath: viper.GetString("app.runtimeRootPath"),
		LogFileLocation: viper.GetString("app.logFileLocation"),
		LogFileSaveName: viper.GetString("app.logFileSaveName"),
		LogFileExt:      viper.GetString("app.logFileExt"),
		TimeFormat:      viper.GetString("app.timeFormat"),
	}

	ServerConfig = &Server{
		RunMode:      viper.GetString("server.runMode"),
		HttpPort:     viper.GetInt("server.httpPort"),
		ReadTimeout:  time.Duration(viper.GetInt("server.readTimeout")),
		WriteTimeout: time.Duration(viper.GetInt("server.writeTimeout")),
	}

	KafkaConfig = &Kafka{
		Locations: viper.GetStringSlice("kafka.locations"),
	}
}
