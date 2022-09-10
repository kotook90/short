package logrus

import (
	logger "github.com/sirupsen/logrus"
	"os"
	"sync"
)


func LogInit() (*os.File, *logger.Entry) {
	logger.SetFormatter(&logger.JSONFormatter{})
	standardFields := logger.Fields{}
	logFile, err := os.OpenFile("logrus/error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logger.Warnf("Error opening file: %v", err)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	logger.SetOutput(logFile)
	wg.Done()
	wg.Wait()
	hlog := logger.WithFields(standardFields)

	return logFile, hlog
}
