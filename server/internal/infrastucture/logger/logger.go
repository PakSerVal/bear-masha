package logger

import (
	"log"

	"go.uber.org/zap"
)

var logger *zap.Logger

func InitLogger(isDev bool) {
	var err error
	if isDev {
		logger, err = zap.NewDevelopment()
	} else {
		cfg := zap.NewProductionConfig()
		cfg.DisableCaller = true
		cfg.OutputPaths = []string{"stdout"}
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
		logger, err = cfg.Build()
	}
	if err != nil {
		log.Fatal("cannot init zap", err)
	}
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}
