package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var Log *logrus.Logger

func InitLogger() {
	file, err := os.OpenFile("application.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("failed open log file %e", err)
	}

	Log = logrus.New()
	Log.SetOutput(io.MultiWriter(os.Stdout, file))
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// IN PROD NAIKAN KE WARN
	Log.SetLevel(logrus.DebugLevel)
}
