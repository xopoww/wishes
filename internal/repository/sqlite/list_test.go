package sqlite_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/xopoww/wishes/internal/models"
	"github.com/xopoww/wishes/internal/repository/sqlite"
	"github.com/xopoww/wishes/internal/service"
)

func TestGetUserLists(t *testing.T) {
	tcs := []struct {
		name    string
		migs    []*migrate.Migration
		wantLen int
	}{
		{
			name: "no lists",
			migs: []*migrate.Migration{
				upMigrationFromString(t,
					`INSERT INTO Users (user_name, pwd_hash) VALUES ("user", "cGFzc3dvcmQ=")`,
					testMigrationVersionStart,
				),
			},
			wantLen: 0,
		},
		{
			name: "one list",
			migs: []*migrate.Migration{
				upMigrationFromString(t,
					`INSERT INTO Users (user_name, pwd_hash) VALUES ("user", "cGFzc3dvcmQ=")`,
					testMigrationVersionStart,
				),
				upMigrationFromString(t,
					`INSERT INTO Lists (title, owner_id) SELECT title, Users.id AS owner_id FROM `+
						`(SELECT "list" AS title) JOIN Users ON Users.user_name = "user"`,
					testMigrationVersionStart+1,
				),
			},
			wantLen: 1,
		},
		{
			name: "two lists",
			migs: []*migrate.Migration{
				upMigrationFromString(t,
					`INSERT INTO Users (user_name, pwd_hash) VALUES ("user", "cGFzc3dvcmQ=")`,
					testMigrationVersionStart,
				),
				upMigrationFromString(t,
					`INSERT INTO Lists (title, owner_id) SELECT title, Users.id AS owner_id FROM `+
						`(SELECT "list" AS title) JOIN Users ON Users.user_name = "user"`,
					testMigrationVersionStart+1,
				),
				upMigrationFromString(t,
					`INSERT INTO Lists (title, owner_id) SELECT title, Users.id AS owner_id FROM `+
						`(SELECT "list" AS title) JOIN Users ON Users.user_name = "user"`,
					testMigrationVersionStart+1,
				),
			},
			wantLen: 2,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			dbs := newTestDatabase(t, tc.migs...)
			repo, err := sqlite.NewRepository(dbs, trace(t))
			if err != nil {
				t.Fatalf("new repo: %s", err)
			}

			ctx, cancel := context.WithCancel(context.Background())
			t.Cleanup(cancel)

			uid, err := repo.CheckUsername(ctx, "user")
			if err != nil {
				t.Fatalf("check user: %s", err)
			}

			lids, err := repo.GetUserLists(ctx, uid)
			if err != nil {
				t.Fatalf("get user lists: %s", err)
			}
			if len(lids) != tc.wantLen {
				t.Fatalf("lids len: want %d, got %d", tc.wantLen, len(lids))
			}
		})
	}
}

func TestGetList(t *testing.T) {
	dbs := newTestDatabase(t,
		upMigrationFromString(t,
			`INSERT INTO Users (user_name, pwd_hash) VALUES ("user", "cGFzc3dvcmQ=")`,
			testMigrationVersionStart,
		),
		upMigrationFromString(t,
			`INSERT INTO Lists (title, owner_id) SELECT title, Users.id AS owner_id FROM `+
				`(SELECT "list" AS title) JOIN Users ON Users.user_name = "user"`,
			testMigrationVersionStart+1,
		),
	)
	repo, err := sqlite.NewRepository(dbs, trace(t))
	if err != nil {
		t.Fatalf("new repo: %s", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	uid, err := repo.CheckUsername(ctx, "user")
	if err != nil {
		t.Fatalf("check user: %s", err)
	}

	lids, err := repo.GetUserLists(ctx, uid)
	if err != nil {
		t.Fatalf("get user lists: %s", err)
	}
	if len(lids) != 1 {
		t.Fatalf("get user lists: wrong len (%d)", len(lids))
	}
	lid := lids[0]

	want := &models.List{
		ID:      lid,
		OwnerID: uid,
		Title:   "list",
	}
	got, err := repo.GetList(ctx, lid)
	if err != nil {
		t.Fatalf("get list: %s", err)
	}
	assertListsEq(t, want, got)

	_, err = repo.GetList(ctx, lid+50)
	if !errors.Is(err, service.ErrNotFound) {
		t.Fatalf("get wrong list: want %#v, got %#v", service.ErrNotFound, err)
	}
}

func TestAddList(t *testing.T) {
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
		t.Fatalf("add: %s", err)
	}

	tcs := []struct {
		name    string
		owner   int64
		items   []models.ListItem
		wantErr error
	}{
		{
			name:  "no items",
			owner: user.ID,
		},
		{
			name:  "one item",
			owner: user.ID,
			items: []models.ListItem{
				{Title: "foo"},
			},
		},
		{
			name:  "one item with desc",
			owner: user.ID,
			items: []models.ListItem{
				{Title: "foo", Desc: "description of an item"},
			},
		},
		{
			name:  "many items",
			owner: user.ID,
			items: []models.ListItem{
				{Title: "foo", Desc: "description of an item"},
				{Title: "bar"},
				{Title: "baz", Desc: "another description of an item"},
			},
		},
		{
			name:    "wrong owner",
			owner:   user.ID + 50,
			wantErr: service.ErrNotFound,
		},
		{
			name:  "wrong owner with items",
			owner: user.ID + 50,
			items: []models.ListItem{
				{Title: "foo"},
			},
			wantErr: service.ErrNotFound,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			want := &models.List{
				Title:   "list",
				OwnerID: tc.owner,
				Items:   tc.items,
			}
			cctx, cancel := context.WithCancel(ctx)
			t.Cleanup(cancel)

			got, err := repo.AddList(cctx, want)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("err: want %#v, got %#v", tc.wantErr, err)
			}
			if err != nil {
				return
			}
			want.ID = got.ID
			assertListsEq(t, want, got)

			got, err = repo.GetList(cctx, want.ID)
			if err != nil {
				t.Fatalf("get list: %s", err)
			}
			assertListsEq(t, want, got)
		})
	}
}

func TestEditList(t *testing.T) {
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
		t.Fatalf("add: %s", err)
	}

	tcs := []struct {
		name string
		a    models.List
		b    models.List
	}{
		{
			name: "rename",
			a: models.List{
				Title: "old_list",
			},
			b: models.List{
				Title: "new_list",
			},
		},
		{
			name: "rename with items",
			a: models.List{
				Title: "old_list",
				Items: []models.ListItem{{Title: "foo"}},
			},
			b: models.List{
				Title: "new_list",
				Items: []models.ListItem{{Title: "foo"}},
			},
		},

		{
			name: "add items",
			a: models.List{
				Title: "list",
			},
			b: models.List{
				Title: "list",
				Items: []models.ListItem{{Title: "foo"}},
			},
		},

		{
			name: "append items",
			a: models.List{
				Title: "list",
				Items: []models.ListItem{{Title: "foo"}, {Title: "bar"}},
			},
			b: models.List{
				Title: "list",
				Items: []models.ListItem{{Title: "foo"}, {Title: "bar"}, {Title: "baz"}},
			},
		},
		{
			name: "prepend items",
			a: models.List{
				Title: "list",
				Items: []models.ListItem{{Title: "bar"}, {Title: "baz"}},
			},
			b: models.List{
				Title: "list",
				Items: []models.ListItem{{Title: "foo"}, {Title: "bar"}, {Title: "baz"}},
			},
		},
		{
			name: "insert items",
			a: models.List{
				Title: "list",
				Items: []models.ListItem{{Title: "foo"}, {Title: "baz"}},
			},
			b: models.List{
				Title: "list",
				Items: []models.ListItem{{Title: "foo"}, {Title: "bar"}, {Title: "baz"}},
			},
		},
		{
			name: "rearrange items",
			a: models.List{
				Title: "list",
				Items: []models.ListItem{{Title: "foo"}, {Title: "bar"}},
			},
			b: models.List{
				Title: "list",
				Items: []models.ListItem{{Title: "bar"}, {Title: "foo"}},
			},
		},
	}
	var lastLid int64
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			for _, direction := range []string{"a->b", "b->a"} {
				var (
					old models.List
					new models.List
				)
				if direction == "a->b" {
					old = tc.a
					new = tc.b
				} else {
					old = tc.b
					new = tc.a
				}
				old.OwnerID = user.ID
				t.Run(direction, func(t *testing.T) {
					cctx, cancel := context.WithCancel(ctx)
					t.Cleanup(cancel)

					list, err := repo.AddList(cctx, &old)
					if err != nil {
						t.Fatalf("add list: %s", err)
					}
					lastLid = list.ID

					// only for comparison
					new.OwnerID = user.ID
					new.ID = list.ID

					err = repo.EditList(cctx, &new)
					if err != nil {
						t.Fatalf("edit list: %s", err)
					}

					got, err := repo.GetList(cctx, new.ID)
					if err != nil {
						t.Fatalf("get list: %s", err)
					}
					assertListsEq(t, &new, got)
				})
			}
		})
	}

	t.Run("not found", func(t *testing.T) {
		cctx, cancel := context.WithCancel(ctx)
		t.Cleanup(cancel)

		err := repo.EditList(cctx, &models.List{
			ID:      lastLid + 50,
			OwnerID: user.ID,
			Title:   "list",
		})
		if !errors.Is(err, service.ErrNotFound) {
			t.Fatalf("want %#v, got %#v", service.ErrNotFound, err)
		}
	})
}

func TestDeleteList(t *testing.T) {
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
		t.Fatalf("add: %s", err)
	}

	list, err := repo.AddList(ctx, &models.List{
		OwnerID: user.ID,
		Title:   "list",
	})
	if err != nil {
		t.Fatalf("add list: %s", err)
	}

	err = repo.DeleteList(ctx, list)
	if err != nil {
		t.Fatalf("delete list: %s", err)
	}

	_, err = repo.GetList(ctx, list.ID)
	if !errors.Is(err, service.ErrNotFound) {
		t.Fatalf("get list: want %#v, got %#v", service.ErrNotFound, err)
	}

	err = repo.DeleteList(ctx, list)
	if !errors.Is(err, service.ErrNotFound) {
		t.Fatalf("delete again: want %#v, got %#v", service.ErrNotFound, err)
	}
}

func assertListsEq(t *testing.T, want, got *models.List) {
	if want == nil && got == nil {
		return
	}
	if want == nil || got == nil {
		t.Fatalf("want %+v, got %+v", want, got)
	}
	if want.ID != got.ID {
		t.Fatalf("want %+v, got %+v", want, got)
	}
	if want.Title != got.Title {
		t.Fatalf("want %+v, got %+v", want, got)
	}
	if want.OwnerID != got.OwnerID {
		t.Fatalf("want %+v, got %+v", want, got)
	}
	if len(want.Items) != len(got.Items) {
		t.Fatalf("want %+v, got %+v", want, got)
	}
	for i := range want.Items {
		if want.Items[i] != got.Items[i] {
			t.Fatalf("want %+v, got %+v", want, got)
		}
	}
}