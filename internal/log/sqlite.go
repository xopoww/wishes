package log

import (
	"time"

	"github.com/rs/zerolog"
	"github.com/xopoww/wishes/internal/repository/sqlite"
)

func Sqlite(l zerolog.Logger) (t sqlite.Trace) {
	l = l.With().Str(zerolog.CallerFieldName, "sqlite").Logger()

	t.OnQuery = func(si sqlite.OnQueryStartInfo) func(sqlite.OnQueryDoneInfo) {
		e := l.Trace().
			Str("method", si.Method).
			Str("query", si.Query)
		if len(si.Args) > 0 {
			arr := zerolog.Arr()
			for _, arg := range si.Args {
				arr = arr.Interface(arg)
			}
			e = e.Array("args", arr)
		}
		e.Msg("sql query start")
		start := time.Now()
		return func(di sqlite.OnQueryDoneInfo) {
			ee := l.Trace().Dur("latency", time.Since(start))
			if di.Error != nil {
				ee = ee.Err(di.Error)
			}
			ee.Msg("sql query done")
		}
	}

	t.OnExec = func(si sqlite.OnExecStartInfo) func(sqlite.OnExecDoneInfo) {
		e := l.Trace().
			Str("query", si.Query)
		if len(si.Args) > 0 {
			arr := zerolog.Arr()
			for _, arg := range si.Args {
				arr = arr.Interface(arg)
			}
			e = e.Array("args", arr)
		}
		e.Msg("sql exec start")
		start := time.Now()
		return func(di sqlite.OnExecDoneInfo) {
			ee := l.Trace().Dur("latency", time.Since(start))
			if di.Error != nil {
				ee = ee.Err(di.Error)
			} else {
				if liid, err := di.Result.LastInsertId(); err != nil {
					ee = ee.AnErr("liid_error", err)
				} else {
					ee = ee.Int64("liid", liid)
				}
				if ra, err := di.Result.RowsAffected(); err != nil {
					ee = ee.AnErr("ra_error", err)
				} else {
					ee = ee.Int64("ra", ra)
				}
			}
			ee.Msg("sql exec done")
		}
	}

	t.OnConnect = func(si sqlite.OnConnectStartInfo) func(sqlite.OnConnectDoneInfo) {
		l.Info().
			Str("dbs", si.DBS).
			Msg("connect start")
		start := time.Now()
		return func(di sqlite.OnConnectDoneInfo) {
			l.Debug().
				Dur("latency", time.Since(start)).
				Msg("connect finished")
		}
	}

	t.OnTxBegin = func(si sqlite.OnTxBeginStartInfo) func(sqlite.OnTxBeginDoneInfo) {
		return func(di sqlite.OnTxBeginDoneInfo) {
			if di.Error != nil {
				l.Error().Err(di.Error).Msg("tx begin error")
			} else {
				l.Debug().Msg("tx begin done")
			}
		}
	}

	t.OnTxEnd = func(si sqlite.OnTxEndStartInfo) func(sqlite.OnTxEndDoneInfo) {
		return func(di sqlite.OnTxEndDoneInfo) {
			var result string
			if si.Commit {
				result = "commit"
			} else {
				result = "rollback"
			}
			if di.Error != nil {
				l.Error().Err(di.Error).Str("result", result).Msg("tx end error")
			} else {
				l.Debug().Str("result", result).Msg("tx end done")
			}
		}
	}

	return t
}
