package logging

import (
	"fmt"
	"os"
	"time"

	"github.com/azusachino/ficus/pkg/conf"
)

// getLogFilePath return log file path
func getLogFilePath() string {
	appConfig := conf.Config.App
	return fmt.Sprintf("%s%s%s", appConfig.RuntimeRootPath, string(os.PathSeparator), appConfig.LogFileLocation)
}

// getLogFileName return today's log name
func getLogFileName() string {
	appConfig := conf.Config.App
	return fmt.Sprintf("%s-%s.%s",
		appConfig.LogFileSaveName,
		time.Now().Format(appConfig.TimeFormat),
		appConfig.LogFileExt,
	)
}
