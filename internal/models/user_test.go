package models_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/xopoww/wishes/internal/models"
)

func TestCheckUsername(t *testing.T) {
	dbs := newTestDatabase(t,
		upMigrationFromString(t,
			`INSERT INTO Users (user_name, pwd_hash) VALUES ("user", "cGFzc3dvcmQ=")`,
			testMigrationVersionStart,
		),
	)
	if err := models.Connect(dbs); err != nil {
		t.Fatalf("connect: %s", err)
	}

	usernames := []string{
		"user",
		"User",
		"admin",
	}

	for _, username := range usernames {
		t.Run(username, func(t *testing.T) {
			err := models.CheckUsername(username)
			if username == "user" {
				if !errors.Is(err, models.ErrNameTaken) {
					t.Fatalf("want %#v, got %#v", models.ErrNameTaken, err)
				}	
			} else {
				if err != nil {
					t.Fatalf("want nil, got %#v", err)
				}
			}

		})
	}

}

func TestRegister(t *testing.T) {
	dbs := newTestDatabase(t)
	if err := models.Connect(dbs); err != nil {
		t.Fatalf("connect: %s", err)
	}

	user, err := models.Register("user", "password")
	if err != nil {
		t.Fatalf("register: got %s", err)
	}
	if user == nil || user.Name != "user" {
		t.Fatalf("register user: got %+v", user)
	}

	if err := models.CheckUsername("user"); !errors.Is(err, models.ErrNameTaken) {
		t.Fatalf("check username: want %#v, got %#v", models.ErrNameTaken, err)
	}

	_, err = models.Register("user", "password")
	if !errors.Is(err, models.ErrNameTaken) {
		t.Fatalf("register dupe: want %#v, got %#v", models.ErrNameTaken, err)
	}
}

func TestLogin(t *testing.T) {
	dbs := newTestDatabase(t)
	if err := models.Connect(dbs); err != nil {
		t.Fatalf("connect: %s", err)
	}

	want, err := models.Register("user", "password")
	if err != nil {
		t.Fatalf("register: %s", err)
	}

	tcs := []struct{
		username string
		password string
	}{
		{"user", "password"},
		{"User", "password"},
		{"user", "PASSWORD"},
		{"user", "password123"},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("(%s:%s)", tc.username, tc.password), func(t *testing.T) {
			got, err := models.Login(tc.username, tc.password)
			if err != nil {
				t.Fatalf("login: %s", err)
			}
			if tc.username == "user" && tc.password == "password" {
				if got == nil || *got != *want {
					t.Fatalf("user: want %+v, got %+v", want, got)
				}
			} else {
				if got != nil {
					t.Fatalf("user: want nil, got %+v", got)
				}
			}
		})
	}
}