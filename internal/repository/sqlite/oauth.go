package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/mattn/go-sqlite3"
	"github.com/xopoww/wishes/internal/models"
	"github.com/xopoww/wishes/internal/service"
)

// CheckOAuth checks for OAuth record and returns corresponding user ID on success
func (r *handle) CheckOAuth(ctx context.Context, provider string, uid string) (id int64, err error) {
	row := r.tracer().QueryRowxContext(ctx, `SELECT user_id FROM OAuth WHERE provider = $1 AND external_id = $2`, provider, uid)
	err = row.Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		err = fmt.Errorf("provider=%q, external_id=%q: %w", provider, uid, service.ErrNotFound)
	}
	return id, err
}

// AddOAuth adds OAuth record for a user
func (r *handle) AddOAuth(ctx context.Context, provider string, uid string, user *models.User) error {
	_, err := r.tracer().ExecContext(ctx,
		`INSERT INTO OAuth (provider, external_id, user_id) VALUES ($1, $2, $3)`, provider, uid, user.ID,
	)
	var serr sqlite3.Error
	if errors.As(err, &serr) {
		switch serr.ExtendedCode {
		case sqlite3.ErrConstraintUnique:
			err = fmt.Errorf("provider=%q, external_id=%q: %w", provider, uid, service.ErrConflict)
		case sqlite3.ErrConstraintForeignKey:
			err = fmt.Errorf("user_id=%d: %w", user.ID, service.ErrNotFound)
		}
	}
	return err
}
