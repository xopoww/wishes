package sqlite_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/xopoww/wishes/internal/models"
	"github.com/xopoww/wishes/internal/models/fixtures"
	"github.com/xopoww/wishes/internal/repository/sqlite"
	"github.com/xopoww/wishes/internal/service"
)

func TestGetUserLists(t *testing.T) {
	tcs := []struct {
		name          string
		migs          []*migrate.Migration
		wantLen       int
		wantLenPublic int
	}{
		{
			name: "no lists",
			migs: []*migrate.Migration{
				upMigrationFromString(t,
					`INSERT INTO Users (user_name, pwd_hash) VALUES ("user", "cGFzc3dvcmQ=")`,
					testMigrationVersionStart,
				),
			},
			wantLen:       0,
			wantLenPublic: 0,
		},
		{
			name: "public list",
			migs: []*migrate.Migration{
				upMigrationFromString(t,
					`INSERT INTO Users (user_name, pwd_hash) VALUES ("user", "cGFzc3dvcmQ=")`,
					testMigrationVersionStart,
				),
				upMigrationFromString(t,
					`INSERT INTO Lists (title, owner_id, access) SELECT title, Users.id AS owner_id, ListAccessEnum.N as access FROM `+
						`(SELECT "list" AS title) JOIN Users ON Users.user_name = "user" JOIN ListAccessEnum on ListAccessEnum.S = "public"`,
					testMigrationVersionStart+1,
				),
			},
			wantLen:       1,
			wantLenPublic: 1,
		},
		{
			name: "private lists",
			migs: []*migrate.Migration{
				upMigrationFromString(t,
					`INSERT INTO Users (user_name, pwd_hash) VALUES ("user", "cGFzc3dvcmQ=")`,
					testMigrationVersionStart,
				),
				upMigrationFromString(t,
					`INSERT INTO Lists (title, owner_id, access) SELECT title, Users.id AS owner_id, ListAccessEnum.N as access FROM `+
						`(SELECT "list1" AS title) JOIN Users ON Users.user_name = "user" JOIN ListAccessEnum on ListAccessEnum.S = "private"`,
					testMigrationVersionStart+1,
				),
				upMigrationFromString(t,
					`INSERT INTO Lists (title, owner_id, access) SELECT title, Users.id AS owner_id, ListAccessEnum.N as access FROM `+
						`(SELECT "list2" AS title) JOIN Users ON Users.user_name = "user" JOIN ListAccessEnum on ListAccessEnum.S = "link"`,
					testMigrationVersionStart+1,
				),
			},
			wantLen:       2,
			wantLenPublic: 0,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			dbs := newTestDatabase(t, tc.migs...)
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

			lids, err := r.GetUserLists(ctx, uid, true)
			if err != nil {
				t.Errorf("get public user lists: %s", err)
			}
			if len(lids) != tc.wantLenPublic {
				t.Errorf("public lids len: want %d, got %d", tc.wantLenPublic, len(lids))
			}

			lids, err = r.GetUserLists(ctx, uid, false)
			if err != nil {
				t.Errorf("get user lists: %s", err)
			}
			if len(lids) != tc.wantLen {
				t.Errorf("lids len: want %d, got %d", tc.wantLen, len(lids))
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
			`INSERT INTO Lists (title, owner_id, access, revision) SELECT title, Users.id AS owner_id, ListAccessEnum.N as access, revision FROM `+
				`(SELECT "list" AS title, 42 as revision) JOIN Users ON Users.user_name = "user" JOIN ListAccessEnum on ListAccessEnum.S = "link"`,
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

	lids, err := r.GetUserLists(ctx, uid, false)
	if err != nil {
		t.Fatalf("get user lists: %s", err)
	}
	if len(lids) != 1 {
		t.Fatalf("get user lists: wrong len (%d)", len(lids))
	}
	lid := lids[0]

	want := &models.List{
		ID:         lid,
		OwnerID:    uid,
		Title:      "list",
		Access:     models.LinkAccess,
		RevisionID: 42,
	}
	got, err := r.GetList(ctx, lid)
	if err != nil {
		t.Fatalf("get list: %s", err)
	}
	assertListsEq(t, want, got)

	_, err = r.GetList(ctx, lid+50)
	if !errors.Is(err, service.ErrNotFound) {
		t.Fatalf("get wrong list: want %#v, got %#v", service.ErrNotFound, err)
	}
}

func TestGetItemTaken(t *testing.T) {
	dbs := newTestDatabase(t,
		upMigrationFromString(t,
			`INSERT INTO Users (user_name, pwd_hash) VALUES ("user1", "cGFzc3dvcmQ=");`+
				`INSERT INTO Users (user_name, pwd_hash) VALUES ("user2", "cGFzc3dvcmQ=")`,
			testMigrationVersionStart,
		),
		upMigrationFromString(t,
			`INSERT INTO Lists (title, owner_id, access, revision) SELECT title, Users.id AS owner_id, ListAccessEnum.N as access, revision FROM `+
				`(SELECT "list1" AS title, 0 as revision) JOIN Users ON Users.user_name = "user1" JOIN ListAccessEnum on ListAccessEnum.S = "link"; `+

				`INSERT INTO Items (title, list_id, id) SELECT item_title as title, Lists.id AS list_id, item_id as id FROM `+
				`(SELECT "item" AS item_title, 1 as item_id) JOIN Lists ON Lists.title = "list1";`+

				`INSERT INTO Items (title, list_id, id, taken_by) `+
				`SELECT item_title as title, Lists.id AS list_id, item_id as id, Users.id as taken_by FROM `+
				`(SELECT "item" AS item_title, 2 as item_id) JOIN Lists ON Lists.title = "list1" JOIN Users ON Users.user_name == "user2";`+

				`INSERT INTO Lists (title, owner_id, access, revision) SELECT title, Users.id AS owner_id, ListAccessEnum.N as access, revision FROM `+
				`(SELECT "list2" AS title, 0 as revision) JOIN Users ON Users.user_name = "user1" JOIN ListAccessEnum on ListAccessEnum.S = "link"`,
			testMigrationVersionStart+2,
		),
	)
	r, err := sqlite.NewRepository(dbs, trace(t))
	if err != nil {
		t.Fatalf("new repo: %s", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	uid1, err := r.CheckUsername(ctx, "user1")
	if err != nil {
		t.Fatalf("check user1: %s", err)
	}

	lids, err := r.GetUserLists(ctx, uid1, false)
	if err != nil {
		t.Fatalf("get lists: %s", err)
	}
	if len(lids) != 2 {
		t.Fatalf("get user lists: wrong len (%d)", len(lids))
	}
	lid1 := lids[0]
	lid2 := lids[1]

	taken, err := r.GetItemTaken(ctx, lid1, 1)
	if err != nil {
		t.Fatalf("item1: %s", err)
	}
	assertInt64PtrEq(t, nil, taken)

	uid2, err := r.CheckUsername(ctx, "user2")
	if err != nil {
		t.Fatalf("check user2: %s", err)
	}

	taken, err = r.GetItemTaken(ctx, lid1, 2)
	if err != nil {
		t.Fatalf("item2: %s", err)
	}
	assertInt64PtrEq(t, &uid2, taken)

	_, err = r.GetItemTaken(ctx, lid2, 1)
	if !errors.Is(err, service.ErrNotFound) {
		t.Fatalf("wrong lid: want %+v, got %+v", service.ErrNotFound, err)
	}
	_, err = r.GetItemTaken(ctx, lid1, 3)
	if !errors.Is(err, service.ErrNotFound) {
		t.Fatalf("wrong lid: want %+v, got %+v", service.ErrNotFound, err)
	}
}

func TestSetItemTaken(t *testing.T) {
	dbs := newTestDatabase(t,
		upMigrationFromString(t,
			`INSERT INTO Users (user_name, pwd_hash) VALUES ("user1", "cGFzc3dvcmQ=")`,
			testMigrationVersionStart,
		),
		upMigrationFromString(t,
			`INSERT INTO Lists (title, owner_id, access, revision) SELECT title, Users.id AS owner_id, ListAccessEnum.N as access, revision FROM `+
				`(SELECT "list1" AS title, 0 as revision) JOIN Users ON Users.user_name = "user1" JOIN ListAccessEnum on ListAccessEnum.S = "link"; `+

				`INSERT INTO Items (title, list_id, id) SELECT item_title as title, Lists.id AS list_id, item_id as id FROM `+
				`(SELECT "item" AS item_title, 1 as item_id) JOIN Lists ON Lists.title = "list1"; `+

				`INSERT INTO Lists (title, owner_id, access, revision) SELECT title, Users.id AS owner_id, ListAccessEnum.N as access, revision FROM `+
				`(SELECT "list2" AS title, 0 as revision) JOIN Users ON Users.user_name = "user1" JOIN ListAccessEnum on ListAccessEnum.S = "link"`,
			testMigrationVersionStart+2,
		),
	)
	r, err := sqlite.NewRepository(dbs, trace(t))
	if err != nil {
		t.Fatalf("new repo: %s", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	uid1, err := r.CheckUsername(ctx, "user1")
	if err != nil {
		t.Fatalf("check user1: %s", err)
	}

	lids, err := r.GetUserLists(ctx, uid1, false)
	if err != nil {
		t.Fatalf("get lists: %s", err)
	}
	if len(lids) != 2 {
		t.Fatalf("get user lists: wrong len (%d)", len(lids))
	}
	lid1 := lids[0]
	lid2 := lids[1]

	u2, err := r.AddUser(ctx, &models.User{
		Name:     "user2",
		PassHash: []byte("password"),
	})
	if err != nil {
		t.Fatalf("add user: %s", err)
	}

	for _, taken := range []*int64{&u2.ID, nil} {
		err := r.SetItemTaken(ctx, lid1, 1, taken)
		if err != nil {
			t.Fatalf("err: %s", err)
		}
		got, err := r.GetItemTaken(ctx, lid1, 1)
		if err != nil {
			t.Fatalf("get item taken: %s", err)
		}
		assertInt64PtrEq(t, taken, got)
	}

	err = r.SetItemTaken(ctx, lid2, 1, &u2.ID)
	if !errors.Is(err, service.ErrNotFound) {
		t.Fatalf("wrong lid: want %+v, got %+v", service.ErrNotFound, err)
	}
	err = r.SetItemTaken(ctx, lid1, 3, &u2.ID)
	if !errors.Is(err, service.ErrNotFound) {
		t.Fatalf("wrong iid: want %+v, got %+v", service.ErrNotFound, err)
	}
	badUid := u2.ID + 50
	err = r.SetItemTaken(ctx, lid1, 1, &badUid)
	if !errors.Is(err, service.ErrNotFound) {
		t.Fatalf("wrong uid: want %+v, got %+v", service.ErrNotFound, err)
	}
}

func TestGetListItems(t *testing.T) {
	dbs := newTestDatabase(t,
		upMigrationFromString(t,
			`INSERT INTO Users (user_name, pwd_hash) VALUES ("user1", "cGFzc3dvcmQ=")`,
			testMigrationVersionStart,
		),
		upMigrationFromString(t,
			`INSERT INTO Lists (title, owner_id, access, revision) SELECT title, Users.id AS owner_id, ListAccessEnum.N as access, revision FROM `+
				`(SELECT "list1" AS title, 0 as revision) JOIN Users ON Users.user_name = "user1" JOIN ListAccessEnum on ListAccessEnum.S = "link"`,
			testMigrationVersionStart+1,
		),
		upMigrationFromString(t,
			`INSERT INTO Lists (title, owner_id, access, revision) SELECT title, Users.id AS owner_id, ListAccessEnum.N as access, revision FROM `+
				`(SELECT "list2" AS title, 0 as revision) JOIN Users ON Users.user_name = "user1" JOIN ListAccessEnum on ListAccessEnum.S = "link"; `+

				`INSERT INTO Items (title, list_id, id) SELECT item_title as title, Lists.id AS list_id, item_id as id FROM `+
				`(SELECT "item" AS item_title, 42 as item_id) JOIN Lists ON Lists.title = "list2";`,
			testMigrationVersionStart+2,
		),
	)
	r, err := sqlite.NewRepository(dbs, trace(t))
	if err != nil {
		t.Fatalf("new repo: %s", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	uid1, err := r.CheckUsername(ctx, "user1")
	if err != nil {
		t.Fatalf("check user1: %s", err)
	}

	lids, err := r.GetUserLists(ctx, uid1, false)
	if err != nil {
		t.Fatalf("get user lists: %s", err)
	}
	if len(lids) != 2 {
		t.Fatalf("get user lists: wrong len (%d)", len(lids))
	}

	u2, err := r.AddUser(ctx, &models.User{
		Name:     "user2",
		PassHash: []byte("password"),
	})
	if err != nil {
		t.Fatalf("add user: %s", err)
	}

	for i, lid := range lids {
		t.Run(fmt.Sprintf("list #%d", i), func(t *testing.T) {
			cctx, cancel := context.WithCancel(context.Background())
			t.Cleanup(cancel)

			list, err := r.GetList(cctx, lid)
			if err != nil {
				t.Fatalf("get list: %s", err)
			}

			want := *list
			switch want.Title {
			case "list1":
				want.Items = nil
			case "list2":
				want.Items = append(want.Items, models.ListItem{
					ID:    42,
					Title: "item",
				})
			default:
				t.Fatalf("unexpected list.Title: %q", list.Title)
			}

			list.Items, err = r.GetListItems(cctx, list)
			if err != nil {
				t.Fatalf("err: %s", err)
			}
			assertListsEq(t, &want, list)

			if want.Title == "list1" {
				return
			}
			err = r.SetItemTaken(ctx, list.ID, list.Items[0].ID, &u2.ID)
			if err != nil {
				t.Fatalf("set item taken: %s", err)
			}

			list.Items, err = r.GetListItems(cctx, list)
			if err != nil {
				t.Fatalf("err: %s", err)
			}
			want.Items[0].TakenBy = &u2.ID
			assertListsEq(t, &want, list)

		})
	}
}

func TestAddList(t *testing.T) {
	dbs := newTestDatabase(t)
	r, err := sqlite.NewRepository(dbs, trace(t))
	if err != nil {
		t.Fatalf("new repo: %s", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	user, err := r.AddUser(ctx, &models.User{
		Name:     "user",
		PassHash: []byte("password"),
	})
	if err != nil {
		t.Fatalf("add: %s", err)
	}

	tcs := []struct {
		name    string
		owner   int64
		access  models.ListAccess
		rev     int64
		wantErr error
	}{
		{
			name:   "public",
			owner:  user.ID,
			access: models.PublicAccess,
		},
		{
			name:   "private",
			owner:  user.ID,
			access: models.PrivateAccess,
		},
		{
			name:  "custom revision",
			owner: user.ID,
			rev:   42,
		},
		{
			name:    "wrong owner",
			owner:   user.ID + 50,
			wantErr: service.ErrNotFound,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			want := &models.List{
				Title:      "list",
				OwnerID:    tc.owner,
				Access:     tc.access,
				RevisionID: tc.rev,
			}
			cctx, cancel := context.WithCancel(ctx)
			t.Cleanup(cancel)

			got, err := r.AddList(cctx, want)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("err: want %#v, got %#v", tc.wantErr, err)
			}
			if err != nil {
				return
			}
			assertListsEq(t, want, got)

			got, err = r.GetList(cctx, want.ID)
			if err != nil {
				t.Fatalf("get list: %s", err)
			}
			got.Items, err = r.GetListItems(cctx, got)
			if err != nil {
				t.Fatalf("get list items: %s", err)
			}
			assertListsEq(t, want, got)
		})
	}
}

func TestEditList(t *testing.T) {
	dbs := newTestDatabase(t)
	r, err := sqlite.NewRepository(dbs, trace(t))
	if err != nil {
		t.Fatalf("new repo: %s", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	user, err := r.AddUser(ctx, &models.User{
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
			name: "change access",
			a: models.List{
				Title:  "list",
				Access: models.PublicAccess,
			},
			b: models.List{
				Title:  "list",
				Access: models.PrivateAccess,
			},
		},
		{
			name: "change revision",
			a: models.List{
				Title:      "list",
				RevisionID: 0,
			},
			b: models.List{
				Title:      "list",
				RevisionID: 5,
			},
		},
	}
	var lastLid int64
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			old := tc.a
			new := tc.b
			old.OwnerID = user.ID

			cctx, cancel := context.WithCancel(ctx)
			t.Cleanup(cancel)

			list, err := r.AddList(cctx, &old)
			if err != nil {
				t.Fatalf("add list: %s", err)
			}
			lastLid = list.ID

			new.ID = list.ID
			_, err = r.EditList(cctx, &new)
			if err != nil {
				t.Fatalf("edit list: %s", err)
			}

			got, err := r.GetList(cctx, new.ID)
			if err != nil {
				t.Fatalf("get list: %s", err)
			}
			got.Items, err = r.GetListItems(cctx, got)
			if err != nil {
				t.Fatalf("get list items: %s", err)
			}
			new.OwnerID = old.OwnerID
			assertListsEq(t, &new, got)
		})
	}

	t.Run("not found", func(t *testing.T) {
		cctx, cancel := context.WithCancel(ctx)
		t.Cleanup(cancel)

		_, err := r.EditList(cctx, &models.List{
			ID:      lastLid + 50,
			OwnerID: user.ID,
			Title:   "list",
		})
		if !errors.Is(err, service.ErrNotFound) {
			t.Fatalf("want %#v, got %#v", service.ErrNotFound, err)
		}
	})
}

func TestAddListItems(t *testing.T) {
	dbs := newTestDatabase(t)
	r, err := sqlite.NewRepository(dbs, trace(t))
	if err != nil {
		t.Fatalf("new repo: %s", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	user, err := r.AddUser(ctx, &models.User{
		Name:     "user",
		PassHash: []byte("password"),
	})
	if err != nil {
		t.Fatalf("add user: %s", err)
	}

	list := fixtures.List()
	list.OwnerID = user.ID
	list, err = r.AddList(ctx, list)
	if err != nil {
		t.Fatalf("add list: %s", err)
	}

	items := fixtures.Items(3)
	list.Items, err = r.AddListItems(ctx, list, items)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	new, err := r.GetList(ctx, list.ID)
	if err != nil {
		t.Fatalf("get list: %s", err)
	}
	new.Items, err = r.GetListItems(ctx, new)
	if err != nil {
		t.Fatalf("get list items: %s", err)
	}
	assertListsEq(t, list, new)
}

func TestDeleteListItems(t *testing.T) {
	dbs := newTestDatabase(t)
	r, err := sqlite.NewRepository(dbs, trace(t))
	if err != nil {
		t.Fatalf("new repo: %s", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	user, err := r.AddUser(ctx, &models.User{
		Name:     "user",
		PassHash: []byte("password"),
	})
	if err != nil {
		t.Fatalf("add user: %s", err)
	}

	list := fixtures.List(fixtures.Items(3)...)
	list.OwnerID = user.ID
	list, err = r.AddList(ctx, list)
	if err != nil {
		t.Fatalf("add list: %s", err)
	}

	ids := []int64{list.Items[0].ID, list.Items[2].ID}
	err = r.DeleteListItems(ctx, list, ids)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	new, err := r.GetList(ctx, list.ID)
	if err != nil {
		t.Fatalf("get list: %s", err)
	}
	new.Items, err = r.GetListItems(ctx, new)
	if err != nil {
		t.Fatalf("get list items: %s", err)
	}
	list.Items = list.Items[1:2]
	assertListsEq(t, list, new)
}

func TestDeleteList(t *testing.T) {
	dbs := newTestDatabase(t)
	r, err := sqlite.NewRepository(dbs, trace(t))
	if err != nil {
		t.Fatalf("new repo: %s", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	user, err := r.AddUser(ctx, &models.User{
		Name:     "user",
		PassHash: []byte("password"),
	})
	if err != nil {
		t.Fatalf("add: %s", err)
	}

	list, err := r.AddList(ctx, &models.List{
		OwnerID: user.ID,
		Title:   "list",
	})
	if err != nil {
		t.Fatalf("add list: %s", err)
	}

	err = r.DeleteList(ctx, list)
	if err != nil {
		t.Fatalf("delete list: %s", err)
	}

	_, err = r.GetList(ctx, list.ID)
	if !errors.Is(err, service.ErrNotFound) {
		t.Fatalf("get list: want %#v, got %#v", service.ErrNotFound, err)
	}

	err = r.DeleteList(ctx, list)
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
	if want.Access != got.Access {
		t.Fatalf("want %+v, got %+v", want, got)
	}
	if want.RevisionID != got.RevisionID {
		t.Fatalf("want %+v, got %+v", want, got)
	}
	if len(want.Items) != len(got.Items) {
		t.Fatalf("want %+v, got %+v", want, got)
	}
	for i := range want.Items {
		if want.Items[i].ID != got.Items[i].ID {
			t.Fatalf("want %+v, got %+v", want, got)
		}
		if want.Items[i].Title != got.Items[i].Title {
			t.Fatalf("want %+v, got %+v", want, got)
		}
		if want.Items[i].Desc != got.Items[i].Desc {
			t.Fatalf("want %+v, got %+v", want, got)
		}
		assertInt64PtrEq(t, want.Items[i].TakenBy, got.Items[i].TakenBy)
	}
}

func assertInt64PtrEq(t *testing.T, want, got *int64) {
	p2s := func(p *int64) string {
		if p == nil {
			return "<nil>"
		}
		return fmt.Sprint(*p)
	}

	if want == nil && got == nil {
		return
	}
	if want == nil || got == nil || *want != *got {
		t.Fatalf("want %s, got %s", p2s(want), p2s(got))
	}
}
