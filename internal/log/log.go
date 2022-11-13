package log

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	output := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "01-02-2006 15:04:05"}
	log.Logger = zerolog.New(output).With().Timestamp().Logger().Level(zerolog.TraceLevel)
}

func Printf(format string, a ...interface{}) {
	log.Printf(format, a...)
}

func Logger() zerolog.Logger {
	return log.Logger
}
