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
	loggerWithCaller = logger.With().Caller().Logger().With().Timestamp().Logger()
}

func GetLogger() zerolog.Logger {
	return logger
}

func Info(msg string) {
	loggerWithCaller.Info().Msg(msg)
}

func Debug(msg string) {
	loggerWithCaller.Debug().Msg(msg)
}
