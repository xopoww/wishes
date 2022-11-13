package sqlite

import (
	"github.com/jmoiron/sqlx"
	"github.com/xopoww/wishes/internal/service"
)

type repository struct {
	db *sqlx.DB
	t  Trace
}

func NewRepository(dbs string, t Trace) (service.Repository, error) {
	onDone := traceOnConnect(t, dbs)
	defer onDone()

	dbs += "?_foreign_keys=1"
	db, err := sqlx.Connect("sqlite3", dbs)
	if err != nil {
		return nil, err
	}

	return &repository{db: db, t: t}, nil
}

func (r *repository) Close() error {
	return r.db.Close()
}
