package global

import (
	"github.com/azusachino/ficus/pkg/conf"
	"github.com/azusachino/ficus/pkg/logger"
	"github.com/panjf2000/ants/v2"
	"gorm.io/gorm"
)

var (
	// Global Config
	Config = &conf.FicusConfig{
		App:      conf.AppConfig{},
		Server:   conf.ServerConfig{},
		Database: conf.DatabaseConfig{},
	}

	// Db Handler
	DbEngine *gorm.DB

	// Async Working Pool
	Pool *ants.Pool

	// Global Logger
	Logger *logger.Logger
)
