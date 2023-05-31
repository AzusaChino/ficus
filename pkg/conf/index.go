package conf

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

var (
	Config = &FicusConfig{
		App:      AppConfig{},
		Server:   ServerConfig{},
		Database: DatabaseConfig{},
	}
)

// FicusConfig the application config structure
type FicusConfig struct {
	App      AppConfig      `mapstructure:"app" json:"app" yaml:"app"`
	Server   ServerConfig   `mapstructure:"server" json:"server" yaml:"server"`
	Database DatabaseConfig `mapstructure:"database" json:"database" yaml:"database"`
}

type AppConfig struct {
	RuntimeRootPath string `mapstructure:"runtimeRootPath" json:"runtimeRootPath" yaml:"runtimeRootPath"`
	LogFileLocation string `mapstructure:"logFileLocation" json:"logFileLocation" yaml:"logFileLocation"`
	LogFileSaveName string `mapstructure:"logFileSaveName" json:"logFileSaveName" yaml:"logFileSaveName"`
	LogFileExt      string `mapstructure:"logFileExt" json:"logFileExt" yaml:"logFileExt"`
	TimeFormat      string `mapstructure:"timeFormat" json:"timeFormat" yaml:"timeFormat"`
}

type ServerConfig struct {
	RunMode      string        `mapstructure:"runMode" json:"runMode" yaml:"runMode"`
	HttpPort     int           `mapstructure:"httpPort" json:"httpPort" yaml:"httpPort"`
	ReadTimeout  time.Duration `mapstructure:"readTimeout" json:"readTimeout" yaml:"readTimeout"`
	WriteTimeout time.Duration `mapstructure:"writeTimeout" json:"writeTimeout" yaml:"writeTimeout"`
}

type DatabaseConfig struct {
	DbName      string `mapstructure:"dbName" json:"dbName" yaml:"dbName"`
	DbHost      string `mapstructure:"dbHost" json:"dbHost" yaml:"dbHost"`
	DbPort      int    `mapstructure:"dbPort" json:"dbPort" yaml:"dbPort"`
	DbUser      string `mapstructure:"dbUser" json:"dbUser" yaml:"dbUser"`
	DbPass      string `mapstructure:"dbPass" json:"dbPass" yaml:"dbPass"`
	TablePrefix string `mapstructure:"tablePrefix" json:"tablePrefix" yaml:"tablePrefix"`
}

func init() {
	var err error
	vp := viper.New()
	vp.SetConfigName("ficus")
	vp.SetConfigType("yaml")
	vp.AddConfigPath("configs")
	if err = vp.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error when reading config file: %w", err))
	}
	err = vp.Unmarshal(&Config)
	if err != nil {
		panic(fmt.Errorf("fatal error when unmarshal config file: %w", err))
	}
}
