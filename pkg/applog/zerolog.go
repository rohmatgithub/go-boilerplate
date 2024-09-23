package applog

import (
	"boilerplate/pkg/configs"
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

func InitLogger() (file *os.File, err error) {
	file, err = os.OpenFile(configs.App.LogPath,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		return
	}
	multi := io.MultiWriter(file, os.Stdout)
	logger = zerolog.New(multi).With().Timestamp().Logger()
	loggerWithCaller = logger.With().Caller().Logger()

	return
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
