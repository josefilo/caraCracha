package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {
	config := zap.Config{
		Encoding: "json",
		Level:    zap.NewAtomicLevelAt(zap.InfoLevel),
		OutputPaths: []string{
			"stdout",
			"log.txt",
			//output on the route /metrics

		},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			LevelKey:     "level",
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			TimeKey:      "time",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	logger, _ := config.Build()
	zap.ReplaceGlobals(logger)
}

func Info(msg string, fields ...zap.Field) {
	zap.L().Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	zap.L().Error(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	zap.L().Warn(msg, fields...)
}
