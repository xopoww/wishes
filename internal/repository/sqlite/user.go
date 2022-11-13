package sqlite

import (
	"context"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mattn/go-sqlite3"
	"github.com/xopoww/wishes/internal/models"
	"github.com/xopoww/wishes/internal/service"
)

func (r *repository) CheckUsername(ctx context.Context, username string) (id int64, err error) {
	err = sqlx.GetContext(ctx, r.tracer(r.db), &id, `SELECT id FROM Users WHERE user_name = $1`, username)
	if errors.Is(err, sql.ErrNoRows) {
		err = service.ErrNotFound
	}
	return id, err
}

func (r *repository) GetUser(ctx context.Context, id int64) (*models.User, error) {
	row := r.tracer(r.db).QueryRowxContext(ctx, `SELECT user_name, fname, lname, pwd_hash FROM Users WHERE id = $1`, id)

	user := &models.User{ID: id}
	var b64 string
	err := row.Scan(&user.Name, &user.Fname, &user.Lname, &b64)
	if errors.Is(err, sql.ErrNoRows) {
		err = fmt.Errorf("user_id %d: %w", id, service.ErrNotFound)
	}
	if err != nil {
		return nil, err
	}

	user.PassHash, err = base64.RawStdEncoding.DecodeString(b64)
	if err != nil {
		return nil, fmt.Errorf("base64: %w", err)
	}
	return user, nil
}

func (r *repository) AddUser(ctx context.Context, user *models.User) (*models.User, error) {
	b64 := base64.RawStdEncoding.EncodeToString(user.PassHash)
	res, err := r.tracer(r.db).ExecContext(ctx, `INSERT INTO Users (user_name, pwd_hash) VALUES ($1, $2)`, user.Name, b64)
	var serr sqlite3.Error
	if errors.As(err, &serr) && serr.ExtendedCode == sqlite3.ErrConstraintUnique {
		err = fmt.Errorf("user_name %q: %w", user.Name, service.ErrConflict)
	}
	if err != nil {
		return nil, fmt.Errorf("insert: %w", err)
	}
	user.ID, err = res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("last insert id: %w", err)
	}
	return user, nil
}

func (r *repository) EditUser(ctx context.Context, user *models.User) error {
	res, err := r.tracer(r.db).ExecContext(ctx,
		`UPDATE Users SET fname = $1, lname = $2 WHERE id = $3`,
		user.Fname, user.Lname, user.ID,
	)
	if err != nil {
		return fmt.Errorf("update: %w", err)
	}
	ra, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %w", err)
	}
	if ra == 0 {
		return fmt.Errorf("user_id %d: %w", user.ID, service.ErrNotFound)
	}
	return nil
}
