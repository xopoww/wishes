package db

import (
	"errors"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sqlx.DB
	t  Trace
)

var ErrNotConnected = errors.New("database not connected")

func Connect(dbs string) (err error) {
	onDone := traceOnConnect(t, dbs)
	defer onDone()

	dbs += "?_foreign_keys=1"
	db, err = sqlx.Connect("sqlite3", dbs)
	if err != nil {
		return err
	}
	return nil
}

func Disconnect() error {
	if db == nil {
		return ErrNotConnected
	}
	return db.Close()
}

func WithTrace(trace Trace) {
	t = trace
}
