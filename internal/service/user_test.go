package service_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/xopoww/wishes/internal/models"
	"github.com/xopoww/wishes/internal/service"
	"github.com/xopoww/wishes/internal/testutil"
)

func TestRegister(t *testing.T) {
	tcs := []struct {
		id  int64
		err error
	}{
		{1, nil},
		{0, service.ErrConflict},
	}

	for i, tc := range tcs {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			ctrl, ctx := gomock.WithContext(context.Background(), t)
			r := NewMockRepository(ctrl)
			r.EXPECT().
				AddUser(gomock.Any(), testutil.MatcherFunc(func(x interface{}) error {
					user, ok := x.(*models.User)
					if !ok || user == nil {
						return fmt.Errorf("type: want %T", user)
					}
					if user.Name != "user" {
						return fmt.Errorf("Name: want %q", "user")
					}
					if len(user.PassHash) == 0 {
						return fmt.Errorf("len(Passhash): want > 0")
					}
					return nil
				})).
				Return(&models.User{ID: tc.id}, tc.err)

			s := service.NewService(r)
			id, err := s.Register(ctx, "user", "password")
			if id != tc.id {
				t.Errorf("id: want %d, got %d", tc.id, id)
			}
			if !errors.Is(err, tc.err) {
				t.Errorf("error: want %+v, got %+v", tc.err, err)
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	tcs := []struct {
		err error
	}{
		{nil},
		{service.ErrNotFound},
	}

	client := &models.User{
		ID: 42,
	}
	id := int64(1)

	for i, tc := range tcs {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			ctrl, ctx := gomock.WithContext(context.Background(), t)
			r := NewMockRepository(ctrl)
			r.EXPECT().
				GetUser(gomock.Any(), gomock.Eq(id)).
				Return(&models.User{ID: id}, tc.err)

			s := service.NewService(r)
			user, err := s.GetUser(ctx, id, client)
			if user.ID != id {
				t.Errorf("user.ID: want %d, got %d", id, user.ID)
			}
			if !errors.Is(err, tc.err) {
				t.Errorf("error: want %+v, got %+v", tc.err, err)
			}
		})
	}
}

func TestEditUser(t *testing.T) {
	client := &models.User{
		ID: 42,
	}

	tcs := []struct {
		id  int64
		err error
	}{
		{42, nil},
		{1, service.ErrAccessDenied},
	}
	fname := "John"
	lname := "Doe"

	for i, tc := range tcs {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			ctrl, ctx := gomock.WithContext(context.Background(), t)
			r := NewMockRepository(ctrl)

			if tc.id == client.ID {
				r.EXPECT().
					EditUser(gomock.Any(), testutil.MatcherFunc(func(x interface{}) error {
						user, ok := x.(*models.User)
						if !ok {
							return fmt.Errorf("type: want %T", user)
						}
						if user.ID != tc.id {
							return fmt.Errorf("ID: want %d", tc.id)
						}
						if user.Fname != fname {
							return fmt.Errorf("Fname: want %q", fname)
						}
						if user.Lname != lname {
							return fmt.Errorf("Lname: want %q", lname)
						}
						return nil
					})).Return(nil)
			}

			s := service.NewService(r)
			err := s.EditUser(ctx, &models.User{
				ID:    tc.id,
				Fname: fname,
				Lname: lname,
			}, client)

			if !errors.Is(err, tc.err) {
				t.Errorf("error: want %+v, got %+v", tc.err, err)
			}
		})
	}
}
