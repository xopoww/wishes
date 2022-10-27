package log

import (
	"errors"
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
			if errors.Is(di.Error, db.ErrNotConnected) {
				l.Error().
					Err(di.Error).
					Msg("check user failed")
			} else {
				l.Debug().
					Dur("latency", time.Since(start)).
					Err(di.Error).
					Msg("check user done")
			}
		}
	}
	return t
}