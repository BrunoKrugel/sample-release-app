package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {

	cfg := zapcore.EncoderConfig{
		TimeKey:          "dt",
		LevelKey:         "level",
		MessageKey:       "message",
		CallerKey:        "caller",
		ConsoleSeparator: " ",
		EncodeLevel:      zapcore.LowercaseLevelEncoder,
		EncodeTime:       zapcore.RFC3339NanoTimeEncoder,
		EncodeCaller:     zapcore.ShortCallerEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg),
		zapcore.AddSync(os.Stdout),
		zap.InfoLevel,
	)

	logger := zap.New(core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	)

	logger.Info("Hello, World!")
}
