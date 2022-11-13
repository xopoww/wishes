package sqlite

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type extTracer struct {
	sqlx.ExtContext
	t Trace
}

func (r *repository) tracer(ext sqlx.ExtContext) sqlx.ExtContext {
	return &extTracer{ext, r.t}
}

func (et *extTracer) QueryContext(ctx context.Context, query string, args ...interface{}) (rows *sql.Rows, err error) {
	onDone := traceOnQuery(et.t, "QueryContext", query, args)
	defer func() { onDone(err) }()
	rows, err = et.ExtContext.QueryContext(ctx, query, args...)
	return
}

func (et *extTracer) QueryxContext(ctx context.Context, query string, args ...interface{}) (rows *sqlx.Rows, err error) {
	onDone := traceOnQuery(et.t, "QueryxContext", query, args)
	defer func() { onDone(err) }()
	rows, err = et.ExtContext.QueryxContext(ctx, query, args...)
	return
}

func (et *extTracer) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	onDone := traceOnQuery(et.t, "QueryRowxContext", query, args)
	defer func() { onDone(nil) }()
	return et.ExtContext.QueryRowxContext(ctx, query, args...)
}

type result struct {
	inner sql.Result
	liid  *int64
	ra    *int64
}

func (r *result) LastInsertId() (int64, error) {
	if r.liid != nil {
		return *r.liid, nil
	}
	liid, err := r.inner.LastInsertId()
	if err == nil {
		r.liid = &liid
	}
	return liid, err
}

func (r *result) RowsAffected() (int64, error) {
	if r.ra != nil {
		return *r.ra, nil
	}
	ra, err := r.inner.RowsAffected()
	if err == nil {
		r.ra = &ra
	}
	return ra, err
}

func (et *extTracer) ExecContext(ctx context.Context, query string, args ...interface{}) (r sql.Result, err error) {
	onDone := traceOnExec(et.t, query, args)
	defer func() { onDone(r, err) }()
	r, err = et.ExtContext.ExecContext(ctx, query, args...)
	r = &result{inner: r}
	return
}
