package models

import (
	"errors"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var db *sqlx.DB

var ErrNotConnected = errors.New("database not connected")

func Connect(dbs string) (err error) {
	db, err = sqlx.Open("sqlite3", dbs)
	if err != nil {
		return err
	}
	return nil
}