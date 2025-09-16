package logger

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Init() {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	logDir := filepath.Join(homeDir, ".alistcli")
	logFilePath := filepath.Join(logDir, "alistcli.log")

	zap.ReplaceGlobals(zap.New(
		zapcore.NewCore(
			getEncoder(),
			getWriteSyncer(logFilePath),
			getLogLevel(),
		),
		zap.AddCaller(),
	))

}

func Sync(l *zap.Logger) {
	err := l.Sync()
	if err != nil {
		log.Println("Error when flushing zap logger")
	}
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogLevel() zapcore.Level {
	return zapcore.ErrorLevel
}

func getWriteSyncer(logFilePath string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    5,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
	}

	return zapcore.AddSync(io.MultiWriter(os.Stdout, lumberJackLogger))
}
