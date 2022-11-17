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
						GetUserLists(gomock.Any(), gomock.Eq(a.ID), gomock.Eq(client.ID != a.ID)).
						Return(ret.lids, ret.err)
					
					s := service.NewService(repo, NewMockListTokenProvider(ctrl))
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
			repo := NewMockRepository(ctrl)

			repo.EXPECT().
				GetList(gomock.Any(), gomock.Eq(list.ID)).
				Return(list, nil)

			ltp := NewMockListTokenProvider(ctrl)
			if client.ID == a.ID {
				ltp.EXPECT().
					GenerateToken(gomock.Eq(service.ListClaims{ListID: list.ID})).
					Return(token, nil)
			}
			
			s := service.NewService(repo, ltp)
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
		goodToken = "good token"
		wrongToken = "wrong token"
		badToken = "bad token"
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
					repo := NewMockRepository(ctrl)
					ltp := NewMockListTokenProvider(ctrl)

					repo.EXPECT().
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
									return service.ListClaims{ListID: list.ID+1}, nil
								default:
									return service.ListClaims{}, errors.New("bad token")
								}
							}).AnyTimes()
					}

					s := service.NewService(repo, ltp)
					_, err := s.GetList(ctx, list.ID, client, token)
					if !errors.Is(err, wantErr) {
						t.Fatalf("want %+v, got %+v", wantErr, err)
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
	
	s := service.NewService(repo, NewMockListTokenProvider(ctrl))
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
					Return(new, nil)
			}

			s := service.NewService(repo, NewMockListTokenProvider(ctrl))
			_, err := s.EditList(ctx, new, tc.client)
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

			s := service.NewService(repo, NewMockListTokenProvider(ctrl))
			err := s.DeleteList(ctx, new, tc.client)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("err: want %+v, got %+v", tc.wantErr, err)
			}
		})
	}
}