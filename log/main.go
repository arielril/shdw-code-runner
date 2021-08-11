package log

import (
	"os"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger
var once sync.Once

const debugLogEnv = "DEBUG"

func init() {
	once.Do(func() {
		logger = logrus.New()
		logger.SetFormatter(&logrus.TextFormatter{
			QuoteEmptyFields: true,
			TimestampFormat:  time.StampMilli,
		})

		if os.Getenv("LOG_LEVEL") == debugLogEnv {
			logger.SetLevel(logrus.DebugLevel)
		} else {
			logger.SetLevel(logrus.InfoLevel)
		}

		logger.SetOutput(os.Stdout)
	})
}

func GetInstance() *logrus.Logger {
	return logger
}
