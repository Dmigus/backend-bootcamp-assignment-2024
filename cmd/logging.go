package main

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func getLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	// TODO: параметризировать уровень логирования
	level, err := zap.ParseAtomicLevel("INFO")
	if err != nil {
		return nil, err
	}
	cfg.Level = level
	cfg.OutputPaths = []string{"stdout"}
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	cfg.DisableStacktrace = true
	cfg.DisableCaller = true
	logger, err := cfg.Build()
	return logger, err
}
