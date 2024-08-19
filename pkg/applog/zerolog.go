package applog

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

var (
	logger           zerolog.Logger
	loggerWithCaller zerolog.Logger
)

// type CustomHook struct{}

// func (h CustomHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
//     e.Str("custom_field", "custom_value")
// }

func InitLogger(file *os.File) {
	multi := io.MultiWriter(file, os.Stdout)
	logger = zerolog.New(multi).With().Timestamp().Logger()
	loggerWithCaller = logger.With().Caller().Logger()
}

func GetLogger() zerolog.Logger {
	return logger
}

func Info() *zerolog.Event {
	return loggerWithCaller.Info()
}

func Debug() *zerolog.Event {
	return loggerWithCaller.Debug()
}
