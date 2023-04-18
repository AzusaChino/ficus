package logging

import (
	"fmt"
	"os"
	"time"

	"github.com/azusachino/ficus/pkg/conf"
)

// getLogFilePath return log file path
func getLogFilePath() string {
	return fmt.Sprintf("%s%s%s", conf.AppConfig.RuntimeRootPath, string(os.PathSeparator), conf.AppConfig.LogFileLocation)
}

// getLogFileName return today's log name
func getLogFileName() string {
	return fmt.Sprintf("%s-%s.%s",
		conf.AppConfig.LogFileSaveName,
		time.Now().Format(conf.AppConfig.TimeFormat),
		conf.AppConfig.LogFileExt,
	)
}
