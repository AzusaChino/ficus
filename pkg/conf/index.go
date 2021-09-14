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

type Grpc struct {
	Server string
	Port   string
}

var GrpcConfig *Grpc

func Setup() {
	var err error
	// using local viper, not global one
	vp := viper.New()
	vp.SetConfigName("application")
	vp.SetConfigType("yml")
	vp.AddConfigPath("conf")
	if err = vp.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error when reading config file: %w \n", err))
	}

	AppConfig = &App{
		RuntimeRootPath: vp.GetString("app.runtimeRootPath"),
		LogFileLocation: vp.GetString("app.logFileLocation"),
		LogFileSaveName: vp.GetString("app.logFileSaveName"),
		LogFileExt:      vp.GetString("app.logFileExt"),
		TimeFormat:      vp.GetString("app.timeFormat"),
	}

	ServerConfig = &Server{
		RunMode:      vp.GetString("server.runMode"),
		HttpPort:     vp.GetInt("server.httpPort"),
		ReadTimeout:  time.Duration(vp.GetInt("server.readTimeout")),
		WriteTimeout: time.Duration(vp.GetInt("server.writeTimeout")),
	}

	KafkaConfig = &Kafka{
		Locations: vp.GetStringSlice("kafka.locations"),
	}

	GrpcConfig = &Grpc{
		Server: vp.GetString("grpc.server"),
		Port:   vp.GetString("grpc.port"),
	}
}
