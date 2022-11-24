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

func TestOAuthRegister(t *testing.T) {
	ctrl, ctx := gomock.WithContext(context.Background(), t)
	r := NewMockRepository(ctrl)
	tx := NewMockTransaction(ctrl)
	op := NewMockOAuthProvider(ctrl)

	s := service.NewService(r, NewMockListTokenProvider(ctrl))
	s.AddOAuthProvider("provider", op)

	u := fixtures.User()

	// happy
	op.EXPECT().
		Validate(gomock.Any(), gomock.Eq("oauth_token")).
		Return("ext_id", nil)
	r.EXPECT().Begin().Return(tx, nil)
	tx.EXPECT().
		AddUser(gomock.Any(), testutil.MatcherFunc(func(x interface{}) error {
			user, ok := x.(*models.User)
			if !ok || user == nil {
				return fmt.Errorf("type: %T", user)
			}
			if user.Name != u.Name {
				return fmt.Errorf("name: %q", u.Name)
			}
			return nil
		})).
		Return(u, nil)
	tx.EXPECT().
		AddOAuth(gomock.Any(), gomock.Eq("provider"), gomock.Eq("ext_id"), gomock.Eq(u)).
		Return(nil)
	tx.EXPECT().Commit().Return(nil)

	got, err := s.OAuthRegister(ctx, u.Name, "provider", "oauth_token")
	if err != nil {
		t.Errorf("happy: %s", err)
	}
	if got != u.ID {
		t.Errorf("happy: want %d, got %d", u.ID, got)
	}

	// validate err
	op.EXPECT().
		Validate(gomock.Any(), gomock.Eq("oauth_token")).
		Return("", errors.New("some validate error"))
	_, err = s.OAuthRegister(ctx, u.Name, "provider", "oauth_token")
	if !errors.Is(err, service.ErrOAuth) {
		t.Errorf("validate err: want %+v, got %+v", service.ErrOAuth, err)
	}

	// unknown provider
	_, err = s.OAuthRegister(ctx, u.Name, "wrong_provider", "oauth_token")
	if !errors.Is(err, service.ErrNoProvider) {
		t.Errorf("unknown provider: want %+v, got %+v", service.ErrNoProvider, err)
	}
}
