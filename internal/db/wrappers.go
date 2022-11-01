package db

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type extTracer struct {
	sqlx.Ext
}

func tracer(ext sqlx.Ext) sqlx.Ext {
	return extTracer{ext}
}

func (et extTracer) Query(query string, args ...interface{}) (rows *sql.Rows, err error) {
	onDone := traceOnQuery(t, "Query", query, args)
	defer func() { onDone(err) }()
	rows, err = et.Ext.Query(query, args...)
	return
}

func (et extTracer) Queryx(query string, args ...interface{}) (rows *sqlx.Rows, err error) {
	onDone := traceOnQuery(t, "Queryx", query, args)
	defer func() { onDone(err) }()
	rows, err = et.Ext.Queryx(query, args...)
	return
}

func (et extTracer) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	onDone := traceOnQuery(t, "QueryRowx", query, args)
	defer func() { onDone(nil) }()
	return et.Ext.QueryRowx(query, args...)
}

type result struct {
	inner sql.Result
	liid  *int64
	ra 	  *int64
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

func (et extTracer) Exec(query string, args ...interface{}) (r sql.Result, err error) {
	onDone := traceOnExec(t, query, args)
	defer func(){ onDone(r, err) }()
	r, err = et.Ext.Exec(query, args...)
	r = &result{inner: r}
	return
}



