package logrus

import (
	logger "github.com/sirupsen/logrus"
	"log"

	"os"
)

func LogsInit() *logger.Entry {

	logger.SetFormatter(&logger.JSONFormatter{})
	standardFields := logger.Fields{}
	hlog := logger.WithFields(standardFields)
	f, err := os.OpenFile("logs/error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	err = f.Close()
	if err != nil {
		logger.Errorf("Файл логов не закрылся %s", err)
	}
	log.SetOutput(f)
	return hlog
}
