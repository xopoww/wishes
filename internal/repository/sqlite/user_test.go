package sqlite_test

import (
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/xopoww/wishes/internal/models"
	"github.com/xopoww/wishes/internal/repository/sqlite"
	"github.com/xopoww/wishes/internal/service"
)

func TestCheckUsername(t *testing.T) {
	dbs := newTestDatabase(t,
		upMigrationFromString(t,
			`INSERT INTO Users (user_name, pwd_hash) VALUES ("user", "cGFzc3dvcmQ=")`,
			testMigrationVersionStart,
		),
	)
	repo, err := sqlite.NewRepository(dbs, trace(t))
	if err != nil {
		t.Fatalf("new repo: %s", err)
	}

	usernames := []string{
		"user",
		"User",
		"admin",
	}

	for _, username := range usernames {
		t.Run(username, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			t.Cleanup(cancel)

			_, err := repo.CheckUsername(ctx, username)
			if username == "user" {
				if err != nil {
					t.Fatalf("want %#v, got %#v", nil, err)
				}
			} else {
				if !errors.Is(err, service.ErrNotFound) {
					t.Fatalf("want %#v, got %#v", service.ErrNotFound, err)
				}
			}
		})
	}
}

func TestAddUser(t *testing.T) {
	dbs := newTestDatabase(t)
	repo, err := sqlite.NewRepository(dbs, trace(t))
	if err != nil {
		t.Fatalf("new repo: %s", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	user, err := repo.AddUser(ctx, &models.User{
		Name:     "user",
		PassHash: []byte("password"),
	})
	if err != nil {
		t.Fatalf("err: got %s", err)
	}
	if user == nil || user.Name != "user" {
		t.Fatalf("user: got %+v", user)
	}

	id, err := repo.CheckUsername(ctx, "user")
	if err != nil {
		t.Fatalf("check user error: want %#v, got %#v", nil, err)
	}
	if id != user.ID {
		t.Fatalf("check user id: want %d, got %d", user.ID, id)
	}

	_, err = repo.AddUser(ctx, &models.User{
		Name:     "user",
		PassHash: []byte("password"),
	})
	if !errors.Is(err, service.ErrConflict) {
		t.Fatalf("register dupe: want %#v, got %#v", service.ErrConflict, err)
	}
}

func TestGetUser(t *testing.T) {
	dbs := newTestDatabase(t)
	repo, err := sqlite.NewRepository(dbs, trace(t))
	if err != nil {
		t.Fatalf("new repo: %s", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	want, err := repo.AddUser(ctx, &models.User{
		Name:     "user",
		PassHash: []byte("password"),
	})
	if err != nil {
		t.Fatalf("add: %s", err)
	}
	if want == nil {
		t.Fatalf("add: got nil user")
	}

	got, err := repo.GetUser(ctx, want.ID)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	assertUserEq(t, want, got)

	_, err = repo.GetUser(ctx, want.ID+50)
	if !errors.Is(err, service.ErrNotFound) {
		t.Fatalf("get wrong id: want %#v, got %#v", service.ErrNotFound, err)
	}
}

func TestEditUser(t *testing.T) {
	dbs := newTestDatabase(t)
	repo, err := sqlite.NewRepository(dbs, trace(t))
	if err != nil {
		t.Fatalf("new repo: %s", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	want, err := repo.AddUser(ctx, &models.User{
		Name:     "user",
		PassHash: []byte("password"),
	})
	if err != nil {
		t.Fatalf("add: %s", err)
	}
	if want == nil {
		t.Fatalf("add: got nil user")
	}

	want.Fname = "John"
	want.Lname = "Doe"
	err = repo.EditUser(ctx, want)
	if err != nil {
		t.Fatalf("edit user: %s", err)
	}

	got, err := repo.GetUser(ctx, want.ID)
	if err != nil {
		t.Fatalf("get user: %s", err)
	}
	assertUserEq(t, want, got)
}

func assertUserEq(t *testing.T, want, got *models.User) {
	if want == nil && got == nil {
		return
	}
	if want == nil || got == nil {
		t.Fatalf("want %+v, got %+v", want, got)
	}
	if want.ID != got.ID {
		t.Fatalf("want %+v, got %+v", want, got)
	}
	if want.Name != got.Name {
		t.Fatalf("want %+v, got %+v", want, got)
	}
	if want.Fname != got.Fname {
		t.Fatalf("want %+v, got %+v", want, got)
	}
	if want.Lname != got.Lname {
		t.Fatalf("want %+v, got %+v", want, got)
	}
	if !bytes.Equal(want.PassHash, got.PassHash) {
		t.Fatalf("want %+v, got %+v", want, got)
	}
}
