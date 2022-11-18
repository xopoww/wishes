package sqlite

import (
	"github.com/jmoiron/sqlx"
	repo "github.com/xopoww/wishes/internal/service/repository"
)

type handle struct {
	ext sqlx.ExtContext
	t   Trace
}

var _ repo.Handle = (*handle)(nil)

func (h *handle) db() *sqlx.DB {
	db, ok := h.ext.(*sqlx.DB)
	if !ok || db == nil {
		panic("sqlite: handle.db() called on wrong handle")
	}
	return db
}

func (h *handle) tx() *sqlx.Tx {
	tx, ok := h.ext.(*sqlx.Tx)
	if !ok || tx == nil {
		panic("sqlite: handle.tx() called on wrong handle")
	}
	return tx
}

type transaction struct {
	handle
}

func (t *transaction) Commit() error {
	return t.handle.tx().Commit()
}

func (t *transaction) Rollback() error {
	return t.handle.tx().Rollback()
}

type repository struct {
	handle
}

func NewRepository(dbs string, t Trace) (repo.Repository, error) {
	onDone := traceOnConnect(t, dbs)
	defer onDone()

	dbs += "?_foreign_keys=1"
	db, err := sqlx.Connect("sqlite3", dbs)
	if err != nil {
		return nil, err
	}

	return &repository{handle: handle{ext: db, t: t}}, nil
}

func (r *repository) Begin() (repo.Transaction, error) {
	tx, err := r.handle.db().Beginx()
	if err != nil {
		return nil, err
	}
	return &transaction{handle: handle{ext: tx, t: r.t}}, nil
}

func (r *repository) Close() error {
	return r.handle.db().Close()
}
