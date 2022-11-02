package db_test

import (
	"errors"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/xopoww/wishes/internal/db"
)

func TestGetUserLists(t *testing.T) {
	withTrace(t)

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
			if err := db.Connect(dbs); err != nil {
				t.Fatalf("connect: %s", err)
			}

			uid, err := db.CheckUser("user")
			if err != nil {
				t.Fatalf("check user: %s", err)
			}

			lids, err := db.GetUserLists(uid, uid)
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
	withTrace(t)

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
	if err := db.Connect(dbs); err != nil {
		t.Fatalf("connect: %s", err)
	}

	uid, err := db.CheckUser("user")
	if err != nil {
		t.Fatalf("check user: %s", err)
	}

	lids, err := db.GetUserLists(uid, uid)
	if err != nil {
		t.Fatalf("get user lists: %s", err)
	}
	if len(lids) != 1 {
		t.Fatalf("get user lists: wrong len (%d)", len(lids))
	}
	lid := lids[0]

	list, err := db.GetList(lid, uid)
	if err != nil {
		t.Fatalf("get list: %s", err)
	}
	if list == nil {
		t.Fatalf("get list: nil list")
	}
	if list.Title != "list" {
		t.Fatalf("list title: want %q, got %q", "list", list.Title)
	}
	if list.OwnerID != uid {
		t.Fatalf("list owner_id: want %d, got %d", uid, list.OwnerID)
	}

	_, err = db.GetList(lid+50, uid)
	if !errors.Is(err, db.ErrNotFound) {
		t.Fatalf("get wrong list: want %#v, got %#v", db.ErrNotFound, err)
	}
}

func TestAddList(t *testing.T) {
	withTrace(t)

	dbs := newTestDatabase(t)
	if err := db.Connect(dbs); err != nil {
		t.Fatalf("connect: %s", err)
	}

	user, err := db.AddUser("user", []byte("password"))
	if err != nil {
		t.Fatalf("register: %s", err)
	}
	if user == nil {
		t.Fatalf("register: nil user")
	}

	tcs := []struct {
		name    string
		owner   int64
		items   []db.ListItem
		wantErr error
	}{
		{
			name:  "no items",
			owner: user.ID,
		},
		{
			name:  "one item",
			owner: user.ID,
			items: []db.ListItem{
				{Title: "foo"},
			},
		},
		{
			name:  "one item with desc",
			owner: user.ID,
			items: []db.ListItem{
				{Title: "foo", Desc: "description of an item"},
			},
		},
		{
			name:  "many items",
			owner: user.ID,
			items: []db.ListItem{
				{Title: "foo", Desc: "description of an item"},
				{Title: "bar"},
				{Title: "baz", Desc: "another description of an item"},
			},
		},
		{
			name:    "wrong owner",
			owner:   user.ID + 50,
			wantErr: db.ErrNotFound,
		},
		{
			name:  "wrong owner with items",
			owner: user.ID + 50,
			items: []db.ListItem{
				{Title: "foo"},
			},
			wantErr: db.ErrNotFound,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			want := &db.List{
				Title:   "list",
				OwnerID: tc.owner,
				Items:   tc.items,
			}

			lid, err := db.AddList(want.Title, want.Items, want.OwnerID)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("err: want %#v, got %#v", tc.wantErr, err)
			}
			if err != nil {
				return
			}
			want.ID = lid

			got, err := db.GetList(lid, want.ID)
			if err != nil {
				t.Fatalf("get list: %s", err)
			}
			assertListsEq(t, want, got)
		})
	}
}

func TestEditList(t *testing.T) {
	withTrace(t)

	dbs := newTestDatabase(t)
	if err := db.Connect(dbs); err != nil {
		t.Fatalf("connect: %s", err)
	}

	user, err := db.AddUser("user", []byte("password"))
	if err != nil {
		t.Fatalf("register: %s", err)
	}
	if user == nil {
		t.Fatalf("register: nil user")
	}

	tcs := []struct {
		name string
		a    db.List
		b    db.List
	}{
		{
			name: "rename",
			a: db.List{
				Title: "old_list",
			},
			b: db.List{
				Title: "new_list",
			},
		},
		{
			name: "rename with items",
			a: db.List{
				Title: "old_list",
				Items: []db.ListItem{{Title: "foo"}},
			},
			b: db.List{
				Title: "new_list",
				Items: []db.ListItem{{Title: "foo"}},
			},
		},

		{
			name: "add items",
			a: db.List{
				Title: "list",
			},
			b: db.List{
				Title: "list",
				Items: []db.ListItem{{Title: "foo"}},
			},
		},

		{
			name: "append items",
			a: db.List{
				Title: "list",
				Items: []db.ListItem{{Title: "foo"}, {Title: "bar"}},
			},
			b: db.List{
				Title: "list",
				Items: []db.ListItem{{Title: "foo"}, {Title: "bar"}, {Title: "baz"}},
			},
		},
		{
			name: "prepend items",
			a: db.List{
				Title: "list",
				Items: []db.ListItem{{Title: "bar"}, {Title: "baz"}},
			},
			b: db.List{
				Title: "list",
				Items: []db.ListItem{{Title: "foo"}, {Title: "bar"}, {Title: "baz"}},
			},
		},
		{
			name: "insert items",
			a: db.List{
				Title: "list",
				Items: []db.ListItem{{Title: "foo"}, {Title: "baz"}},
			},
			b: db.List{
				Title: "list",
				Items: []db.ListItem{{Title: "foo"}, {Title: "bar"}, {Title: "baz"}},
			},
		},
		{
			name: "rearrange items",
			a: db.List{
				Title: "list",
				Items: []db.ListItem{{Title: "foo"}, {Title: "bar"}},
			},
			b: db.List{
				Title: "list",
				Items: []db.ListItem{{Title: "bar"}, {Title: "foo"}},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			for _, direction := range []string{"a->b", "b->a"} {
				var (
					old db.List
					new db.List
				)
				if direction == "a->b" {
					old = tc.a
					new = tc.b
				} else {
					old = tc.b
					new = tc.a
				}
				t.Run(direction, func(t *testing.T) {
					lid, err := db.AddList(old.Title, old.Items, user.ID)
					if err != nil {
						t.Fatalf("add list: %s", err)
					}

					// only for comparison
					new.OwnerID = user.ID
					new.ID = lid

					err = db.EditList(&new, user.ID)
					if err != nil {
						t.Fatalf("edit list: %s", err)
					}

					got, err := db.GetList(new.ID, user.ID)
					if err != nil {
						t.Fatalf("get list: %s", err)
					}
					assertListsEq(t, &new, got)
				})
			}
		})
	}

	t.Run("access test", func(t *testing.T) {
		other, err := db.AddUser("other", []byte("password"))
		if err != nil {
			t.Fatalf("add other user: %s", err)
		}

		lid, err := db.AddList("list", nil, user.ID)
		if err != nil {
			t.Fatalf("add list: %s", err)
		}

		list := &db.List{
			ID:    lid,
			Title: "edited_list",
			Items: nil,
		}
		err = db.EditList(list, other.ID)
		if !errors.Is(err, db.ErrAccessDenied) {
			t.Fatalf("want %#v, got %#v", db.ErrAccessDenied, err)
		}
	})
}

func TestDeleteList(t *testing.T) {
	withTrace(t)

	dbs := newTestDatabase(t)
	if err := db.Connect(dbs); err != nil {
		t.Fatalf("connect: %s", err)
	}

	user, err := db.AddUser("user", []byte("password"))
	if err != nil {
		t.Fatalf("register: %s", err)
	}
	if user == nil {
		t.Fatalf("register: nil user")
	}

	other, err := db.AddUser("other", []byte("password"))
	if err != nil {
		t.Fatalf("register: %s", err)
	}
	if user == nil {
		t.Fatalf("register: nil user")
	}

	lid, err := db.AddList("list", nil, user.ID)
	if err != nil {
		t.Fatalf("add list: %s", err)
	}

	err = db.DeleteList(lid, other.ID)
	if !errors.Is(err, db.ErrAccessDenied) {
		t.Fatalf("want %#v, got %#v", db.ErrAccessDenied, err)
	}

	_, err = db.GetList(lid, user.ID)
	if err != nil {
		t.Fatalf("get list after wrong delete: %s", err)
	}

	err = db.DeleteList(lid, user.ID)
	if err != nil {
		t.Fatalf("delete list: %s", err)
	}

	_, err = db.GetList(lid, user.ID)
	if !errors.Is(err, db.ErrNotFound) {
		t.Fatalf("want %#v, got %#v", db.ErrNotFound, err)
	}
}

func assertListsEq(t *testing.T, want, got *db.List) {
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
