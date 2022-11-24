package sqlite_test

import (
	"context"
	"errors"
	"testing"

	"github.com/xopoww/wishes/internal/models"
	"github.com/xopoww/wishes/internal/models/fixtures"
	"github.com/xopoww/wishes/internal/repository/sqlite"
	"github.com/xopoww/wishes/internal/service"
)

func TestCheckOAuth(t *testing.T) {
	dbs := newTestDatabase(t,
		upMigrationFromString(t,
			`INSERT INTO Users (user_name, pwd_hash) VALUES ("user", "")`,
			testMigrationVersionStart,
		),
		upMigrationFromString(t,
			`INSERT INTO OAuth (provider, external_id, user_id) SELECT provider, external_id, Users.id as user_id FROM `+
			`(SELECT "test_provider" as provider, "test_ext_id" as external_id) JOIN Users ON Users.user_name == "user"`,
			testMigrationVersionStart+1,
		),
	)
	r, err := sqlite.NewRepository(dbs, trace(t))
	if err != nil {
		t.Fatalf("new repo: %s", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	uid, err := r.CheckUsername(ctx, "user")
	if err != nil {
		t.Fatalf("check user: %s", err)
	}

	got, err := r.CheckOAuth(ctx, "test_provider", "test_ext_id")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if got != uid {
		t.Fatalf("want %d, got %d", uid, got)
	}

	_, err = r.CheckOAuth(ctx, "wrong_provider", "test_ext_id")
	if !errors.Is(err, service.ErrNotFound) {
		t.Fatalf("want %+v, got %+v", service.ErrNotFound, err)
	}

	_, err = r.CheckOAuth(ctx, "test_provider", "wrong_ext_id")
	if !errors.Is(err, service.ErrNotFound) {
		t.Fatalf("want %+v, got %+v", service.ErrNotFound, err)
	}
}

func TestAddOAuth(t *testing.T) {
	dbs := newTestDatabase(t)
	r, err := sqlite.NewRepository(dbs, trace(t))
	if err != nil {
		t.Fatalf("new repo: %s", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	a, b := fixtures.TwoUsers()

	a, err = r.AddUser(ctx, a)
	if err != nil {
		t.Fatalf("add a: %s", err)
	}

	err = r.AddOAuth(ctx, "test_provider", "test_ext_id", a)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	got, err := r.CheckOAuth(ctx, "test_provider", "test_ext_id")
	if err != nil {
		t.Fatalf("check: %s", err)
	}
	if got != a.ID {
		t.Fatalf("want %d, got %d", a.ID, got)
	}

	b, err = r.AddUser(ctx, b)
	if err != nil {
		t.Fatalf("add b: %s", err)
	}

	err = r.AddOAuth(ctx, "test_provider", "test_ext_id", b)
	if !errors.Is(err, service.ErrConflict) {
		t.Fatalf("want %+v, got %+v", service.ErrConflict, err)
	}

	err = r.AddOAuth(ctx, "test_provider", "other_ext_id", &models.User{ID: b.ID+50})
	if !errors.Is(err, service.ErrNotFound) {
		t.Fatalf("want %+v, got %+v", service.ErrNotFound, err)
	}
}