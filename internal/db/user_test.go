package db_test

import (
	"errors"
	"testing"

	"github.com/xopoww/wishes/internal/db"
)

func TestCheckUser(t *testing.T) {
	withTrace(t)

	dbs := newTestDatabase(t,
		upMigrationFromString(t,
			`INSERT INTO Users (user_name, pwd_hash) VALUES ("user", "cGFzc3dvcmQ=")`,
			testMigrationVersionStart,
		),
	)
	if err := db.Connect(dbs); err != nil {
		t.Fatalf("connect: %s", err)
	}

	usernames := []string{
		"user",
		"User",
		"admin",
	}

	for _, username := range usernames {
		t.Run(username, func(t *testing.T) {
			err := db.CheckUser(username)
			if username == "user" {
				if !errors.Is(err, db.ErrNameTaken) {
					t.Fatalf("want %#v, got %#v", db.ErrNameTaken, err)
				}
			} else {
				if err != nil {
					t.Fatalf("want nil, got %#v", err)
				}
			}

		})
	}

}

func TestAddUser(t *testing.T) {
	withTrace(t)

	dbs := newTestDatabase(t)
	if err := db.Connect(dbs); err != nil {
		t.Fatalf("connect: %s", err)
	}

	user, err := db.AddUser("user", []byte("password"))
	if err != nil {
		t.Fatalf("register: got %s", err)
	}
	if user == nil || user.Name != "user" {
		t.Fatalf("register user: got %+v", user)
	}

	if err := db.CheckUser("user"); !errors.Is(err, db.ErrNameTaken) {
		t.Fatalf("check username: want %#v, got %#v", db.ErrNameTaken, err)
	}

	_, err = db.AddUser("user", []byte("password"))
	if !errors.Is(err, db.ErrNameTaken) {
		t.Fatalf("register dupe: want %#v, got %#v", db.ErrNameTaken, err)
	}
}

func TestGetFullUser(t *testing.T) {
	withTrace(t)

	dbs := newTestDatabase(t)
	if err := db.Connect(dbs); err != nil {
		t.Fatalf("connect: %s", err)
	}

	want, err := db.AddUser("user", []byte("password"))
	if err != nil {
		t.Fatalf("register: %s", err)
	}
	if want == nil {
		t.Fatalf("register: nil user")
	}

	got, pwd, err := db.GetFullUser("user")
	if err != nil {
		t.Errorf("get user: %s", err)
	}
	if got == nil || *want != *got {
		t.Errorf("get user: want %+v, got %+v", want, got)
	}
	if ws, gs := "password", string(pwd); gs != ws {
		t.Errorf("get user pwd: want %q, got %q", ws, gs)
	}

	_, _, err = db.GetFullUser("foo")
	if !errors.Is(err, db.ErrNotFound) {
		t.Errorf("get foo: want %#v, got %#v", db.ErrNotFound, err)
	}
}

func TestEditUser(t *testing.T) {
	withTrace(t)

	dbs := newTestDatabase(t)
	if err := db.Connect(dbs); err != nil {
		t.Fatalf("connect: %s", err)
	}

	want, err := db.AddUser("user", []byte("password"))
	if err != nil {
		t.Fatalf("register: %s", err)
	}
	if want == nil {
		t.Fatalf("register: nil user")
	}

	want.FirstName = "John"
	want.LastName = "Doe"
	err = db.EditUserInfo(want)
	if err != nil {
		t.Fatalf("edit user: %s", err)
	}

	got, _, err := db.GetFullUser(want.Name)
	if err != nil {
		t.Fatalf("get user: %s", err)
	}
	if got == nil || *got != *want {
		t.Fatalf("want %+v, got %+v", want, got)
	}
}
