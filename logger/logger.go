package logger

import (
	"log/slog"
	"os"
	"sync"

	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	log  *slog.Logger
	once sync.Once
)

func Init() {
	once.Do(func() {
		_ = os.MkdirAll("logs", 0755)

		writer := &lumberjack.Logger{
			Filename:   "logs/app.log",
			MaxSize:    10, // MB
			MaxBackups: 5,
			MaxAge:     7, // days
			Compress:   true,
		}

		handler := NewTextHandler(writer, slog.LevelDebug)

		log = slog.New(handler)

		slog.SetDefault(log)
	})
}

func Info(msg string) {
	log.Info(msg)
}

func Debug(msg string) {
	log.Debug(msg)
}

func Warn(msg string) {
	log.Warn(msg)
}

func Error(msg string, err error) {
	log.Error(msg, err)
}
