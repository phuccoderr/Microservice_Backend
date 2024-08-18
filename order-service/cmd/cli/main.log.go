package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

func main() {

	encoder := getEncoder()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("hello world")
}

// Format Logs
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()

	// 1716714967 -> 2024-05-26T16:16:07.877+0700
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// ts -> timestamp
	encoderConfig.TimeKey = "timestamp"

	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriterSync() zapcore.WriteSyncer {
	file, err := os.OpenFile("./log/log.txt", os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
