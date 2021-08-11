package logging

import (
	"fmt"
	"github.com/AzusaChino/ficus/pkg/conf"
	"time"
)

// getLogFilePath return log file path
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", conf.AppConfig.RuntimeRootPath, conf.AppConfig.LogFileLocation)
}

// getLogFileName return today's log name
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		conf.AppConfig.LogFileSaveName,
		time.Now().Format(conf.AppConfig.TimeFormat),
		conf.AppConfig.LogFileExt)
}
