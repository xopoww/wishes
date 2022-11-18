package log

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	output := zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: "01-02-2006 15:04:05",
		FormatCaller: func(i interface{}) string {
			s, ok := i.(string)
			if !ok {
				s = "root"
			}
			if len(s) > 8 {
				s = fmt.Sprintf("%s#", s[:7])
			}
			return fmt.Sprintf("%-8s ", s)
		},
		PartsOrder: []string{
			zerolog.TimestampFieldName,
			zerolog.CallerFieldName,
			zerolog.LevelFieldName,
			zerolog.MessageFieldName,
		},
	}
	log.Logger = zerolog.New(output).With().Timestamp().Logger().Level(zerolog.TraceLevel)
}

func Printf(format string, a ...interface{}) {
	log.Printf(format, a...)
}

func Logger() zerolog.Logger {
	return log.Logger
}
