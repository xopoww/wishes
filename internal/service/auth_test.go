package service_test

import (
	"context"
	"encoding/base64"
	"errors"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/golang/mock/gomock"
	"github.com/xopoww/wishes/internal/models"
	"github.com/xopoww/wishes/internal/service"
	"github.com/xopoww/wishes/internal/testutil"
)

func TestAuth(t *testing.T) {
	now := time.Now().Add(time.Second * -2)

	secret := []byte("my-awesome-secret")
	t.Setenv(service.TokenSecretVariableName, base64.RawStdEncoding.EncodeToString(secret))

	mustSign := func(token *jwt.Token, key interface{}) string {
		s, err := token.SignedString(key)
		if err != nil {
			t.Fatalf("mustSign: %s", err)
		}
		return s
	}

	user := "user"
	id := int64(42)

	tcs := []struct {
		name    string
		token   string
		wantErr error
	}{
		{
			name: "ok token",
			token: mustSign(jwt.NewWithClaims(jwt.SigningMethodHS256, service.NewJwtClaims(jwt.RegisteredClaims{
				IssuedAt:  jwt.NewNumericDate(now),
				NotBefore: jwt.NewNumericDate(now),
				ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour)),
				Subject:   user,
			}, id)), secret),
		},
		{
			name: "expired token",
			token: mustSign(jwt.NewWithClaims(jwt.SigningMethodHS256, service.NewJwtClaims(jwt.RegisteredClaims{
				IssuedAt:  jwt.NewNumericDate(now.Add(time.Hour * -2)),
				NotBefore: jwt.NewNumericDate(now.Add(time.Hour * -2)),
				ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * -2).Add(time.Hour)),
				Subject:   user,
			}, id)), secret),
			wantErr: service.ErrTokenExpired,
		},
		{
			name: "not active token",
			token: mustSign(jwt.NewWithClaims(jwt.SigningMethodHS256, service.NewJwtClaims(jwt.RegisteredClaims{
				IssuedAt:  jwt.NewNumericDate(now.Add(time.Hour * 1)),
				NotBefore: jwt.NewNumericDate(now.Add(time.Hour * 1)),
				ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * 1).Add(time.Hour)),
				Subject:   user,
			}, id)), secret),
			wantErr: testutil.AnyError{},
		},
		{
			name: "missing exp",
			token: mustSign(jwt.NewWithClaims(jwt.SigningMethodHS256, service.NewJwtClaims(jwt.RegisteredClaims{
				IssuedAt:  jwt.NewNumericDate(now),
				NotBefore: jwt.NewNumericDate(now),
				Subject:   user,
			}, id)), secret),
			wantErr: service.ErrTokenExpired,
		},
		{
			name: "missing nbf",
			token: mustSign(jwt.NewWithClaims(jwt.SigningMethodHS256, service.NewJwtClaims(jwt.RegisteredClaims{
				IssuedAt:  jwt.NewNumericDate(now),
				ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour)),
				Subject:   user,
			}, id)), secret),
			wantErr: testutil.AnyError{},
		},
		{
			name: "wrong method",
			token: mustSign(jwt.NewWithClaims(jwt.SigningMethodNone, service.NewJwtClaims(jwt.RegisteredClaims{
				IssuedAt:  jwt.NewNumericDate(now),
				NotBefore: jwt.NewNumericDate(now),
				ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * 24)),
				Subject:   user,
			}, id)), jwt.UnsafeAllowNoneSignatureType),
			wantErr: testutil.AnyError{},
		},
		{
			name: "wrong secret",
			token: mustSign(jwt.NewWithClaims(jwt.SigningMethodHS256, service.NewJwtClaims(jwt.RegisteredClaims{
				IssuedAt:  jwt.NewNumericDate(now),
				NotBefore: jwt.NewNumericDate(now),
				ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * 24)),
				Subject:   user,
			}, id)), []byte("fake-not-awesome-secret")),
			wantErr: testutil.AnyError{},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			ctrl, ctx := gomock.WithContext(context.Background(), t)
			r := NewMockRepository(ctrl)
			s := service.NewService(r, NewMockListTokenProvider(ctrl))

			client, err := s.Auth(ctx, tc.token)

			if !errors.Is(tc.wantErr, err) {
				t.Fatalf("error: want %#v, got %#v", tc.wantErr, err)
			}
			if err != nil {
				return
			}

			wantClient := &models.User{
				ID:   id,
				Name: user,
			}
			if client == nil || client.ID != wantClient.ID || client.Name != wantClient.Name {
				t.Fatalf("client: want %v, got %v", wantClient, client)
			}
		})
	}
}
