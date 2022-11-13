package log

import (
	"time"

	"github.com/rs/zerolog"
	"github.com/xopoww/wishes/internal/repository/sqlite"
)

func Sqlite(l zerolog.Logger) (t sqlite.Trace) {
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

	return t
}
