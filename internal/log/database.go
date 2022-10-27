package log

import (
	"time"

	"github.com/rs/zerolog"
	"github.com/xopoww/wishes/internal/db"
)

func Database(l zerolog.Logger) (t db.Trace) {
	t.OnCheckUser = func(si db.OnCheckUserStartInfo) func(db.OnCheckUserDoneInfo) {
		l.Debug().
			Str("username", si.Username).
			Msg("check user start")
		start := time.Now()
		return func(di db.OnCheckUserDoneInfo) {
			l.Debug().
				TimeDiff("latency", time.Now(), start).
				Msg("check user done")
		}
	}
	return t
}
