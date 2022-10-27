package db

import (
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/mattn/go-sqlite3"
)

type User struct {
	ID   int
	Name string

	FirstName	string
	LastName	string
}

var (
	ErrNameTaken = errors.New("username already taken")
	ErrNotFound  = errors.New("not found")
)

// CheckUser returns nil if username is not found in the database.
// It returns ErrNameTaken if username is found in the database. All other
// return values indicate internal error during check.
func CheckUser(username string) error {
	onDone := traceOnCheckUser(t, username)
	defer func() { onDone() }()

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

func AddUser(username string, passHash []byte) (*User, error) {
	if db == nil {
		return nil, ErrNotConnected
	}

	hashString := base64.RawStdEncoding.EncodeToString(passHash)
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
		ID:   int(id),
		Name: username,
	}, nil
}

// GetFullUser retrieves user and their passHash from database.
func GetFullUser(username string) (user *User, passHash []byte, err error) {
	if db == nil {
		return nil, nil, ErrNotConnected
	}

	var (
		id         int64
		hashString string
	)
	user = &User{Name: username}
	row := db.QueryRow(`SELECT user_id, fname, lname, pwd_hash FROM Users WHERE user_name = $1`, username)
	err = row.Scan(&id, &user.FirstName, &user.LastName, &hashString)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil, ErrNotFound
	}
	if err != nil {
		return nil, nil, err
	}

	passHash, err = base64.RawStdEncoding.DecodeString(hashString)
	if err != nil {
		return nil, nil, fmt.Errorf("invalid base64 in database: %w", err)
	}

	user.ID = int(id)
	return user, passHash, nil
}

func EditUser(user *User) error {
	if db == nil {
		return ErrNotConnected
	}

	r, err := db.Exec(
		`UPDATE Users SET fname = $1, lname = $2 WHERE user_name = $3`,
		user.FirstName, user.LastName, user.Name,
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
