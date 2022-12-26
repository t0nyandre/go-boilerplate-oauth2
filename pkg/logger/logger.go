package logger

import (
	"encoding/json"
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() *zap.SugaredLogger {
	dateString := time.Now().Format("20060102")
	rawJSON := []byte(`{
		"level": "debug",
		"encoding": "json",
		"errorOutputPaths": ["stderr"],
		"encoderConfig": {
			"messageKey": "msg",
			"levelKey": "level",
			"levelEncoder": "lowercase"
		}
	}`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}

	cfg.OutputPaths = []string{"stdout", fmt.Sprintf("./logs/%s.json", dateString)}
	cfg.EncoderConfig.TimeKey = "timestamp"
	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder

	l, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	logger := l.Sugar()

	return logger
}
