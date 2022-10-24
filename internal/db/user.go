package db

import (
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID   int
	Name string
}

var ErrNameTaken = errors.New("username already taken")

// CheckUsername returns nil if username is not found in the database.
// It returns ErrNameTaken if username is found in the database. All other
// return values indicate internal error during check.
func CheckUsername(username string) error {
	if db == nil {
		return ErrNotConnected
	}

	row := db.QueryRow(`SELECT 1 FROM Users WHERE user_name = $1`, username)
	var unused int
	err := row.Scan(&unused)
	if err == nil {
		return ErrNameTaken
	}
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	}
	return err
}

func Register(username, password string) (*User, error) {
	if db == nil {
		return nil, ErrNotConnected
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return nil, fmt.Errorf("hash: %w", err)
	}
	hashString := base64.RawStdEncoding.EncodeToString(hash)

	r, err := db.Exec(
		`INSERT INTO Users (user_name, pwd_hash) VALUES ($1, $2)`,
		username, hashString,
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
		ID: int(id),
		Name: username,
	}, nil
}

// Login checks user credentials. Non-nil error indicates internal error during check.
// Credentials validity should be checked by user != nil.
func Login(username, password string) (*User, error) {
	if db == nil {
		return nil, ErrNotConnected
	}

	var (
		id int64
		hashString string
	)
	row := db.QueryRow(`SELECT user_id, pwd_hash FROM Users WHERE user_name = $1`, username)
	err := row.Scan(&id, &hashString)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	hash, err := base64.RawStdEncoding.DecodeString(hashString)
	if err != nil {
		return nil, fmt.Errorf("invalid base64 in database: %w", err)
	}

	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		return nil, nil
	}
	return &User{
		ID: int(id),
		Name: username,
	}, nil
}
