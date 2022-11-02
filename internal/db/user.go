package db

import (
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mattn/go-sqlite3"
)

type User struct {
	ID   int64    `db:"user_id"`
	Name string `db:"user_name"`

	FirstName string `db:"fname"`
	LastName  string `db:"lname"`
}

var (
	ErrNameTaken = errors.New("username already taken")
	ErrNotFound  = errors.New("not found")
)

func CheckUser(username string) (id int64, err error) {
	if db == nil {
		return 0, ErrNotConnected
	}

	err = sqlx.Get(tracer(db), &id, `SELECT user_id FROM Users WHERE user_name = $1`, username)
	if errors.Is(err, sql.ErrNoRows) {
		err = ErrNotFound
	}
	return id, err
}

func AddUser(username string, passHash []byte) (*User, error) {
	if db == nil {
		return nil, ErrNotConnected
	}

	hashString := base64.RawStdEncoding.EncodeToString(passHash)
	r, err := sqlx.NamedExec(tracer(db),
		`INSERT INTO Users (user_name, pwd_hash) VALUES (:user_name, :pwd_hash)`,
		map[string]interface{}{
			"user_name": username,
			"pwd_hash":  hashString,
		},
	)
	var serr sqlite3.Error
	if errors.As(err, &serr) && serr.ExtendedCode == sqlite3.ErrConstraintUnique {
		return nil, ErrNameTaken
	}
	if err != nil {
		return nil, fmt.Errorf("insert: %w", err)
	}
	id, err := r.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("get id: %w", err)
	}
	return &User{
		ID:   id,
		Name: username,
	}, nil
}

// GetFullUser retrieves user and their passHash from database.
func GetFullUser(username string) (user *User, passHash []byte, err error) {
	if db == nil {
		return nil, nil, ErrNotConnected
	}

	full := &struct {
		User
		Hash string `db:"pwd_hash"`
	}{}
	err = sqlx.Get(tracer(db), full, `SELECT user_id, fname, lname, pwd_hash FROM Users WHERE user_name = $1`, username)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil, ErrNotFound
	}
	if err != nil {
		return nil, nil, err
	}

	passHash, err = base64.RawStdEncoding.DecodeString(full.Hash)
	if err != nil {
		return nil, nil, fmt.Errorf("invalid base64 in database: %w", err)
	}

	user = &full.User
	user.Name = username
	return user, passHash, nil
}

func GetUserById(id int64) (*User, error) {
	if db == nil {
		return nil, ErrNotConnected
	}

	user := &User{ID: id}
	err := sqlx.Get(tracer(db), user, `SELECT user_name, fname, lname FROM Users WHERE user_id = $1`, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}

func EditUserInfo(user *User) error {
	if db == nil {
		return ErrNotConnected
	}

	r, err := sqlx.NamedExec(tracer(db),
		`UPDATE Users SET fname = :fname, lname = :lname WHERE user_id = :user_id`,
		user,
	)
	if err != nil {
		return fmt.Errorf("update: %s", err)
	}
	ra, err := r.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %s", err)
	}
	if ra == 0 {
		return ErrNotFound
	}
	return nil
}
