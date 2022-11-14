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
	rets := []struct{
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
					repo := NewMockRepository(ctrl)

					repo.EXPECT().
						GetUserLists(gomock.Any(), gomock.Eq(a.ID)).
						Return(ret.lids, ret.err)
					
					s := service.NewService(repo)
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

func TestGetList(t *testing.T) {
	client := fixtures.User()
	rets := []struct{
		name string
		list *models.List
		err  error
	}{
		{name: "no items", list: fixtures.List()},
		{name: "with items",list: fixtures.List(fixtures.Items(3)...)},
		{name: "not found",err: service.ErrNotFound},
	}
	lid := int64(42)

	for _, ret := range rets {
		t.Run(ret.name, func(t *testing.T) {
			ctrl, ctx := gomock.WithContext(context.Background(), t)
			repo := NewMockRepository(ctrl)

			repo.EXPECT().
				GetList(gomock.Any(), gomock.Eq(lid)).
				DoAndReturn(func(_, _ interface{}) (*models.List, error) {
					if ret.err != nil {
						return ret.list, ret.err
					}
					return &models.List{
						ID: ret.list.ID,
						OwnerID: ret.list.OwnerID,
						Title: ret.list.Title,
					}, nil
				})
			
			if ret.err == nil {
				repo.EXPECT().
					GetListItems(gomock.Any(), testutil.MatcherFunc(func(x interface{}) error {
						list, ok := x.(*models.List)
						if !ok || list == nil {
							return fmt.Errorf("type: want %T", list)
						}
						if list.ID != ret.list.ID {
							return fmt.Errorf("ID: want %d", ret.list.ID)
						}
						if list.OwnerID != ret.list.OwnerID {
							return fmt.Errorf("OwnerID: want %d", ret.list.OwnerID)
						}
						if list.Title != ret.list.Title {
							return fmt.Errorf("Title: want %q", ret.list.Title)
						}
						return nil
					})).
					DoAndReturn(func(_ interface{}, _ *models.List) (*models.List, error) {
						return ret.list, nil
					})
			}
			
			s := service.NewService(repo)
			list, err := s.GetList(ctx, lid, client)
			if !errors.Is(err, ret.err) {
				t.Fatalf("err: want %+v, got %+v", ret.err, err)
			}
			if err != nil {
				return
			}
			if list != ret.list {
				t.Fatalf("list: want %+v, got %+v", ret.list, list)
			}
		})
	}
}

func TestAddList(t *testing.T) {
	client := fixtures.User()
	list := fixtures.List()
	lid := int64(42)

	ctrl, ctx := gomock.WithContext(context.Background(), t)
	repo := NewMockRepository(ctrl)

	repo.EXPECT().
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
		DoAndReturn(func (_ interface{}, l *models.List) (*models.List, error) {
			ll := &models.List{}
			*ll = *l
			ll.ID = lid
			return ll, nil
		})
	
	s := service.NewService(repo)
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

	badId := int64(404)

	tcs := []struct{
		client  *models.User
		lid     int64
		wantErr error
	}{
		{
			client: a,
			lid: old.ID,
		},
		{
			client: b,
			lid: old.ID,
			wantErr: service.ErrAccessDenied,
		},
		{
			client: a,
			lid: badId,
			wantErr: service.ErrNotFound,
		},
		{
			client: b,
			lid: badId,
			wantErr: service.ErrNotFound,
		},
	}

	for _, tc := range tcs {
		t.Run(fmt.Sprintf("client=%s,bad_id=%t", tc.client.Name, tc.lid == badId), func(t *testing.T) {
			new := fixtures.List()
			new.ID = tc.lid
			new.Title = "new title"
			
			ctrl, ctx := gomock.WithContext(context.Background(), t)
			repo := NewMockRepository(ctrl)

			var (
				rerr error
				rlist *models.List
			)
			if tc.lid == badId {
				rerr = service.ErrNotFound
			} else {
				rlist = old
			}
			repo.EXPECT().
				GetList(gomock.Any(), gomock.Eq(new.ID)).
				Return(rlist, rerr)
			
			if tc.wantErr == nil {
				repo.EXPECT().
					EditList(gomock.Any(), gomock.Eq(new)).
					Return(nil)
			}

			s := service.NewService(repo)
			err := s.EditList(ctx, new, tc.client)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("err: want %+v, got %+v", tc.wantErr, err)
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

	tcs := []struct{
		client  *models.User
		lid     int64
		wantErr error
	}{
		{
			client: a,
			lid: old.ID,
		},
		{
			client: b,
			lid: old.ID,
			wantErr: service.ErrAccessDenied,
		},
		{
			client: a,
			lid: badId,
			wantErr: service.ErrNotFound,
		},
		{
			client: b,
			lid: badId,
			wantErr: service.ErrNotFound,
		},
	}

	for _, tc := range tcs {
		t.Run(fmt.Sprintf("client=%s,bad_id=%t", tc.client.Name, tc.lid == badId), func(t *testing.T) {
			new := fixtures.List()
			new.ID = tc.lid
			
			ctrl, ctx := gomock.WithContext(context.Background(), t)
			repo := NewMockRepository(ctrl)

			var (
				rerr error
				rlist *models.List
			)
			if tc.lid == badId {
				rerr = service.ErrNotFound
			} else {
				rlist = old
			}
			repo.EXPECT().
				GetList(gomock.Any(), gomock.Eq(new.ID)).
				Return(rlist, rerr)
			
			if tc.wantErr == nil {
				repo.EXPECT().
					DeleteList(gomock.Any(), gomock.Eq(old)).
					Return(nil)
			}

			s := service.NewService(repo)
			err := s.DeleteList(ctx, new, tc.client)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("err: want %+v, got %+v", tc.wantErr, err)
			}
		})
	}
}