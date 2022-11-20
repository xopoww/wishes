package service_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/xopoww/wishes/internal/models"
	"github.com/xopoww/wishes/internal/models/fixtures"
	"github.com/xopoww/wishes/internal/service"
	"github.com/xopoww/wishes/internal/testutil"
)

func TestGetUserLists(t *testing.T) {
	a, b := fixtures.TwoUsers()
	rets := []struct {
		lids []int64
		err  error
	}{
		{lids: []int64{1, 2, 3}},
		{lids: []int64{}},
		{err: service.ErrNotFound},
	}

	for _, client := range []*models.User{a, b} {
		t.Run(fmt.Sprintf("client=%s", client.Name), func(t *testing.T) {
			for _, ret := range rets {
				t.Run(fmt.Sprintf("ret=%+v", ret), func(t *testing.T) {
					ctrl, ctx := gomock.WithContext(context.Background(), t)
					r := NewMockRepository(ctrl)

					r.EXPECT().
						GetUserLists(gomock.Any(), gomock.Eq(a.ID), gomock.Eq(client.ID != a.ID)).
						Return(ret.lids, ret.err)

					s := service.NewService(r, NewMockListTokenProvider(ctrl))
					lids, err := s.GetUserLists(ctx, a.ID, client)
					if !errors.Is(err, ret.err) {
						t.Fatalf("err: want %+v, got %+v", ret.err, err)
					}
					if err != nil {
						return
					}
					if len(lids) != len(ret.lids) {
						t.Fatalf("lids: want %v, got %v", ret.lids, lids)
					}
				})
			}
		})
	}
}

func TestGetListToken(t *testing.T) {
	a, b := fixtures.TwoUsers()

	list := fixtures.List()
	list.ID = 42
	list.OwnerID = a.ID

	token := "awesome-token"

	for _, client := range []*models.User{a, b} {
		t.Run(client.Name, func(t *testing.T) {
			ctrl, ctx := gomock.WithContext(context.Background(), t)
			r := NewMockRepository(ctrl)

			r.EXPECT().
				GetList(gomock.Any(), gomock.Eq(list.ID)).
				Return(list, nil)

			ltp := NewMockListTokenProvider(ctrl)
			if client.ID == a.ID {
				ltp.EXPECT().
					GenerateToken(gomock.Eq(service.ListClaims{ListID: list.ID})).
					Return(token, nil)
			}

			s := service.NewService(r, ltp)
			got, err := s.GetListToken(ctx, list.ID, client)
			if client.ID == a.ID {
				if got != token {
					t.Fatalf("want %q, got %q", token, got)
				}
			} else {
				if !errors.Is(err, service.ErrAccessDenied) {
					t.Fatalf("want %+v, got %+v", service.ErrAccessDenied, err)
				}
			}
		})
	}
}

func TestGetList(t *testing.T) {
	a, b := fixtures.TwoUsers()

	list := *fixtures.List()
	list.ID = 42
	list.OwnerID = a.ID

	var (
		goodToken  = "good token"
		wrongToken = "wrong token"
		badToken   = "bad token"
	)

	for _, client := range []*models.User{a, b} {
		for _, access := range []models.ListAccess{models.PublicAccess, models.LinkAccess, models.PrivateAccess} {
			list.Access = access
			for _, token := range []*string{nil, &goodToken, &wrongToken, &badToken} {
				name := fmt.Sprintf("%s,%s", client.Name, access)
				if token != nil {
					name += fmt.Sprintf(",%s", *token)
				}

				var wantErr error
				if token != nil {
					if *token != goodToken || access == models.PrivateAccess {
						wantErr = service.ErrAccessDenied
					}
				} else {
					if access != models.PublicAccess && client.ID != a.ID {
						wantErr = service.ErrAccessDenied
					}
				}

				t.Run(name, func(t *testing.T) {
					ctrl, ctx := gomock.WithContext(context.Background(), t)
					r := NewMockRepository(ctrl)
					ltp := NewMockListTokenProvider(ctrl)

					r.EXPECT().
						GetList(gomock.Any(), gomock.Eq(list.ID)).
						Return(&list, nil)

					if token != nil {
						ltp.EXPECT().
							ValidateToken(gomock.Eq(*token)).
							DoAndReturn(func(t string) (service.ListClaims, error) {
								switch t {
								case goodToken:
									return service.ListClaims{ListID: list.ID}, nil
								case wrongToken:
									return service.ListClaims{ListID: list.ID + 1}, nil
								default:
									return service.ListClaims{}, errors.New("bad token")
								}
							}).AnyTimes()
					}

					s := service.NewService(r, ltp)
					_, err := s.GetList(ctx, list.ID, client, token)
					if !errors.Is(err, wantErr) {
						t.Fatalf("want %+v, got %+v", wantErr, err)
					}
				})

				t.Run(fmt.Sprintf("(items)%s", name), func(t *testing.T) {
					ctrl, ctx := gomock.WithContext(context.Background(), t)
					r := NewMockRepository(ctrl)
					tx := NewMockTransaction(ctrl)
					r.EXPECT().Begin().Return(tx, nil)
					ltp := NewMockListTokenProvider(ctrl)

					tx.EXPECT().
						GetList(gomock.Any(), gomock.Eq(list.ID)).
						Return(&list, nil)

					if token != nil {
						ltp.EXPECT().
							ValidateToken(gomock.Eq(*token)).
							DoAndReturn(func(t string) (service.ListClaims, error) {
								switch t {
								case goodToken:
									return service.ListClaims{ListID: list.ID}, nil
								case wrongToken:
									return service.ListClaims{ListID: list.ID + 1}, nil
								default:
									return service.ListClaims{}, errors.New("bad token")
								}
							}).AnyTimes()
					}

					if wantErr == nil {
						items := fixtures.Items(1)
						taken := int64(42)
						items[0].TakenBy = &taken
						tx.EXPECT().GetListItems(gomock.Any(), gomock.Eq(&list)).Return(items, nil)
						tx.EXPECT().Commit().Return(nil)
					} else {
						tx.EXPECT().Rollback().Return(nil)
					}

					s := service.NewService(r, ltp)
					l, err := s.GetListItems(ctx, &models.List{ID: list.ID}, client, token)
					if !errors.Is(err, wantErr) {
						t.Fatalf("want %+v, got %+v", wantErr, err)
					}
					if err == nil && client.ID == a.ID && l.Items[0].TakenBy != nil {
						t.Fatalf("taken_by not obscured")
					}
				})
			}
		}
	}
}

func TestAddList(t *testing.T) {
	client := fixtures.User()
	list := fixtures.List()
	lid := int64(42)

	ctrl, ctx := gomock.WithContext(context.Background(), t)
	r := NewMockRepository(ctrl)

	r.EXPECT().
		AddList(gomock.Any(), testutil.MatcherFunc(func(x interface{}) error {
			l, ok := x.(*models.List)
			if !ok || l == nil {
				return fmt.Errorf("type: want %T", list)
			}
			if l.OwnerID != client.ID {
				return fmt.Errorf("OwnerID: want %d", client.ID)
			}
			return nil
		})).
		DoAndReturn(func(_ interface{}, l *models.List) (*models.List, error) {
			ll := &models.List{}
			*ll = *l
			ll.ID = lid
			return ll, nil
		})

	s := service.NewService(r, NewMockListTokenProvider(ctrl))
	got, err := s.AddList(ctx, list, client)
	if err != nil {
		t.Fatalf("add list: %s", err)
	}
	if got == nil {
		t.Fatalf("got nil")
	}
	if got.ID != lid {
		t.Fatalf("got.ID: want %d, got %d", lid, got.ID)
	}
}

func TestEditList(t *testing.T) {
	a, b := fixtures.TwoUsers()

	old := fixtures.List()
	old.ID = 42
	old.OwnerID = a.ID

	for _, client := range []*models.User{a, b} {
		t.Run(client.Name, func(t *testing.T) {
			new := fixtures.List()
			new.ID = old.ID
			new.Title = "new title"

			ctrl, ctx := gomock.WithContext(context.Background(), t)
			r := NewMockRepository(ctrl)
			tx := NewMockTransaction(ctrl)
			r.EXPECT().Begin().Return(tx, nil)

			tx.EXPECT().
				GetList(gomock.Any(), gomock.Eq(new.ID)).
				Return(old, nil)

			var wantErr error
			if client.ID == a.ID {
				tx.EXPECT().
					EditList(gomock.Any(), gomock.Eq(new)).
					Return(new, nil)
				tx.EXPECT().Commit().Return(nil)
			} else {
				tx.EXPECT().Rollback().Return(nil)
				wantErr = service.ErrAccessDenied
			}

			s := service.NewService(r, NewMockListTokenProvider(ctrl))
			new, err := s.EditList(ctx, new, client)
			if !errors.Is(err, wantErr) {
				t.Fatalf("err: want %+v, got %+v", wantErr, err)
			}
			if err == nil && new.RevisionID != old.RevisionID {
				t.Fatalf("rev: want %d, got %d", old.RevisionID, new.RevisionID)
			}
		})
	}
}

func TestAddListItems(t *testing.T) {
	a, b := fixtures.TwoUsers()

	old := fixtures.List()
	old.ID = 42
	old.OwnerID = a.ID
	old.RevisionID = 5

	for _, client := range []*models.User{a, b} {
		t.Run(client.Name, func(t *testing.T) {
			items := fixtures.Items(3)

			ctrl, ctx := gomock.WithContext(context.Background(), t)
			r := NewMockRepository(ctrl)
			tx := NewMockTransaction(ctrl)
			r.EXPECT().Begin().Return(tx, nil)

			tx.EXPECT().
				GetList(gomock.Any(), gomock.Eq(old.ID)).
				Return(old, nil)

			var wantErr error
			if client.ID == a.ID {
				tx.EXPECT().
					AddListItems(gomock.Any(), gomock.Any(), gomock.Eq(items)).
					Return(items, nil)
				tx.EXPECT().EditList(gomock.Any(), gomock.Any())
				tx.EXPECT().Commit().Return(nil)
			} else {
				tx.EXPECT().Rollback().Return(nil)
				wantErr = service.ErrAccessDenied
			}

			s := service.NewService(r, NewMockListTokenProvider(ctrl))
			new, err := s.AddListItems(ctx, &models.List{ID: old.ID, RevisionID: old.RevisionID}, items, client)
			if !errors.Is(err, wantErr) {
				t.Fatalf("err: want %+v, got %+v", wantErr, err)
			}
			if err == nil && new.RevisionID != old.RevisionID+1 {
				t.Fatalf("rev: want %d, got %d", old.RevisionID+1, new.RevisionID)
			}
			old.RevisionID++

			r.EXPECT().Begin().Return(tx, nil)
			tx.EXPECT().
				GetList(gomock.Any(), gomock.Eq(old.ID)).
				Return(old, nil)
			tx.EXPECT().Rollback().Return(nil)

			_, err = s.AddListItems(ctx, &models.List{ID: old.ID, RevisionID: old.RevisionID - 1}, items, client)
			if wantErr == nil {
				wantErr = service.ErrOutdated
			}
			if !errors.Is(err, wantErr) {
				t.Fatalf("err: want %+v, got %+v", wantErr, err)
			}
		})
	}
}

func TestDeleteListItems(t *testing.T) {
	a, b := fixtures.TwoUsers()

	old := fixtures.List()
	old.ID = 42
	old.OwnerID = a.ID
	old.RevisionID = 5

	for _, client := range []*models.User{a, b} {
		t.Run(client.Name, func(t *testing.T) {
			ids := []int64{1, 3, 5}

			ctrl, ctx := gomock.WithContext(context.Background(), t)
			r := NewMockRepository(ctrl)
			tx := NewMockTransaction(ctrl)
			r.EXPECT().Begin().Return(tx, nil)

			tx.EXPECT().
				GetList(gomock.Any(), gomock.Eq(old.ID)).
				Return(old, nil)

			var wantErr error
			if client.ID == a.ID {
				tx.EXPECT().
					DeleteListItems(gomock.Any(), gomock.Any(), gomock.Eq(ids)).
					Return(nil)
				tx.EXPECT().EditList(gomock.Any(), gomock.Any())
				tx.EXPECT().Commit().Return(nil)
			} else {
				tx.EXPECT().Rollback().Return(nil)
				wantErr = service.ErrAccessDenied
			}

			s := service.NewService(r, NewMockListTokenProvider(ctrl))
			new, err := s.DeleteListItems(ctx, &models.List{ID: old.ID, RevisionID: old.RevisionID}, ids, client)
			if !errors.Is(err, wantErr) {
				t.Fatalf("err: want %+v, got %+v", wantErr, err)
			}
			if err == nil && new.RevisionID != old.RevisionID+1 {
				t.Fatalf("rev: want %d, got %d", old.RevisionID+1, new.RevisionID)
			}
			old.RevisionID++

			r.EXPECT().Begin().Return(tx, nil)
			tx.EXPECT().
				GetList(gomock.Any(), gomock.Eq(old.ID)).
				Return(old, nil)
			tx.EXPECT().Rollback().Return(nil)

			_, err = s.DeleteListItems(ctx, &models.List{ID: old.ID, RevisionID: old.RevisionID - 1}, ids, client)
			if wantErr == nil {
				wantErr = service.ErrOutdated
			}
			if !errors.Is(err, wantErr) {
				t.Fatalf("err: want %+v, got %+v", wantErr, err)
			}
		})
	}
}

func TestDeleteList(t *testing.T) {
	a, b := fixtures.TwoUsers()

	old := fixtures.List()
	old.ID = 42
	old.OwnerID = a.ID

	badId := int64(404)

	tcs := []struct {
		client  *models.User
		lid     int64
		wantErr error
	}{
		{
			client: a,
			lid:    old.ID,
		},
		{
			client:  b,
			lid:     old.ID,
			wantErr: service.ErrAccessDenied,
		},
		{
			client:  a,
			lid:     badId,
			wantErr: service.ErrNotFound,
		},
		{
			client:  b,
			lid:     badId,
			wantErr: service.ErrNotFound,
		},
	}

	for _, tc := range tcs {
		t.Run(fmt.Sprintf("client=%s,bad_id=%t", tc.client.Name, tc.lid == badId), func(t *testing.T) {
			new := fixtures.List()
			new.ID = tc.lid

			ctrl, ctx := gomock.WithContext(context.Background(), t)
			r := NewMockRepository(ctrl)
			tx := NewMockTransaction(ctrl)
			r.EXPECT().Begin().Return(tx, nil)

			var (
				rerr  error
				rlist *models.List
			)
			if tc.lid == badId {
				rerr = service.ErrNotFound
			} else {
				rlist = old
			}
			tx.EXPECT().
				GetList(gomock.Any(), gomock.Eq(new.ID)).
				Return(rlist, rerr)

			if tc.wantErr == nil {
				tx.EXPECT().
					DeleteList(gomock.Any(), gomock.Eq(old)).
					Return(nil)
				tx.EXPECT().Commit().Return(nil)
			} else {
				tx.EXPECT().Rollback().Return(nil)
			}

			s := service.NewService(r, NewMockListTokenProvider(ctrl))
			err := s.DeleteList(ctx, new, tc.client)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("err: want %+v, got %+v", tc.wantErr, err)
			}
		})
	}
}

func TestTakeItem(t *testing.T) {
	a, b := fixtures.TwoUsers()
	list := &models.List{
		ID:         1,
		RevisionID: 1,
	}
	iid := int64(2)
	taken := b.ID + 1

	ctrl, ctx := gomock.WithContext(context.Background(), t)
	r := NewMockRepository(ctrl)
	tx := NewMockTransaction(ctrl)

	tx.EXPECT().
		GetList(gomock.Any(), gomock.Eq(list.ID)).
		Return(&models.List{
			ID:         list.ID,
			RevisionID: list.RevisionID,
			OwnerID:    a.ID,
			Access:     models.PublicAccess,
		}, nil).AnyTimes()

	s := service.NewService(r, NewMockListTokenProvider(ctrl))

	// happy
	r.EXPECT().Begin().Return(tx, nil)
	tx.EXPECT().
		GetItemTaken(gomock.Any(), gomock.Any(), gomock.Eq(iid)).
		Return(nil, nil)
	tx.EXPECT().
		SetItemTaken(gomock.Any(), gomock.Eq(list.ID), gomock.Eq(iid), testutil.MatcherFunc(func(x interface{}) error {
			p, ok := x.(*int64)
			if !ok {
				return fmt.Errorf("type: want %T", p)
			}
			if p == nil || *p != b.ID {
				return fmt.Errorf("want &(%d)", b.ID)
			}
			return nil
		})).
		Return(nil)
	tx.EXPECT().Commit().Return(nil)

	err := s.TakeItem(ctx, list, iid, b, nil)
	if !errors.Is(err, nil) {
		t.Errorf("happy: want %+v, got %+v", nil, err)
	}

	// already taken
	r.EXPECT().Begin().Return(tx, nil)
	tx.EXPECT().
		GetItemTaken(gomock.Any(), gomock.Any(), gomock.Eq(iid)).
		Return(&taken, nil)
	tx.EXPECT().Rollback().Return(nil)

	err = s.TakeItem(ctx, list, iid, b, nil)
	var serr service.ErrAlreadyTaken
	if !errors.As(err, &serr) || serr.TakenBy != taken {
		t.Errorf("already taken: want %+v, got %+v", service.ErrAlreadyTaken{taken}, err)
	}

	// bad revision
	r.EXPECT().Begin().Return(tx, nil)
	tx.EXPECT().Rollback().Return(nil)

	err = s.TakeItem(ctx, &models.List{ID: list.ID, RevisionID: list.RevisionID - 1}, iid, b, nil)
	if !errors.Is(err, service.ErrOutdated) {
		t.Errorf("bad revision: want %+v, got %+v", service.ErrOutdated, err)
	}

	// take by owner
	r.EXPECT().Begin().Return(tx, nil)
	tx.EXPECT().Rollback().Return(nil)

	err = s.TakeItem(ctx, list, iid, a, nil)
	if !errors.Is(err, service.ErrAccessDenied) {
		t.Errorf("take by owner: want %+v, got %+v", service.ErrAccessDenied, err)
	}
}

func TestUntakeItem(t *testing.T) {
	a, b := fixtures.TwoUsers()
	c := fixtures.User()
	c.ID = b.ID + 1
	list := &models.List{
		ID:         1,
		RevisionID: 1,
	}
	iid := int64(2)

	ctrl, ctx := gomock.WithContext(context.Background(), t)
	r := NewMockRepository(ctrl)
	tx := NewMockTransaction(ctrl)

	tx.EXPECT().
		GetList(gomock.Any(), gomock.Eq(list.ID)).
		Return(&models.List{
			ID:         list.ID,
			RevisionID: list.RevisionID,
			OwnerID:    a.ID,
			Access:     models.PublicAccess,
		}, nil).AnyTimes()

	s := service.NewService(r, NewMockListTokenProvider(ctrl))

	// happy
	r.EXPECT().Begin().Return(tx, nil)
	tx.EXPECT().
		GetItemTaken(gomock.Any(), gomock.Any(), gomock.Eq(iid)).
		Return(&b.ID, nil)
	tx.EXPECT().
		SetItemTaken(gomock.Any(), gomock.Eq(list.ID), gomock.Eq(iid), gomock.Eq((*int64)(nil))).
		Return(nil)
	tx.EXPECT().Commit().Return(nil)

	err := s.UntakeItem(ctx, list, iid, b, nil)
	if !errors.Is(err, nil) {
		t.Errorf("happy: want %+v, got %+v", nil, err)
	}

	// not taken by anyone
	r.EXPECT().Begin().Return(tx, nil)
	tx.EXPECT().
		GetItemTaken(gomock.Any(), gomock.Any(), gomock.Eq(iid)).
		Return(nil, nil)
	tx.EXPECT().Rollback().Return(nil)

	err = s.UntakeItem(ctx, list, iid, b, nil)
	if !errors.Is(err, service.ErrConflict) {
		t.Errorf("not taken by anyone: want %+v, got %+v", nil, err)
	}

	// taken by another
	r.EXPECT().Begin().Return(tx, nil)
	tx.EXPECT().
		GetItemTaken(gomock.Any(), gomock.Any(), gomock.Eq(iid)).
		Return(&c.ID, nil)
	tx.EXPECT().Rollback().Return(nil)

	err = s.UntakeItem(ctx, list, iid, b, nil)
	if !errors.Is(err, service.ErrConflict) {
		t.Errorf("taken by another: want %+v, got %+v", nil, err)
	}

	// bad revision
	r.EXPECT().Begin().Return(tx, nil)
	tx.EXPECT().Rollback().Return(nil)

	err = s.UntakeItem(ctx, &models.List{ID: list.ID, RevisionID: list.RevisionID - 1}, iid, b, nil)
	if !errors.Is(err, service.ErrOutdated) {
		t.Errorf("bad revision: want %+v, got %+v", service.ErrOutdated, err)
	}

	// untake by owner
	r.EXPECT().Begin().Return(tx, nil)
	tx.EXPECT().Rollback().Return(nil)

	err = s.UntakeItem(ctx, list, iid, a, nil)
	if !errors.Is(err, service.ErrAccessDenied) {
		t.Errorf("untake by owner: want %+v, got %+v", service.ErrAccessDenied, err)
	}
}
