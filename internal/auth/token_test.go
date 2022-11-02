package auth_test

import (
	"encoding/base64"
	"errors"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/xopoww/wishes/internal/auth"
	"github.com/xopoww/wishes/internal/models"
	"github.com/xopoww/wishes/internal/testutil"
)

func TestValidateToken(t *testing.T) {
	now := time.Now().Add(time.Second * -2)

	secret := []byte("my-awesome-secret")
	t.Setenv(auth.TokenSecretVariableName, base64.RawStdEncoding.EncodeToString(secret))

	mustSign := func(token *jwt.Token, key interface{}) string {
		s, err := token.SignedString(key)
		if err != nil {
			t.Fatalf("mustSign: %s", err)
		}
		return s
	}

	user := "user"
	wantPrincipal := models.Principal(user)

	tcs := []struct {
		name    string
		token   string
		wantErr error
	}{
		{
			name: "ok token",
			token: mustSign(jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
				IssuedAt:  jwt.NewNumericDate(now),
				NotBefore: jwt.NewNumericDate(now),
				ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour)),
				Subject:   user,
			}), secret),
		},
		{
			name: "expired token",
			token: mustSign(jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
				IssuedAt:  jwt.NewNumericDate(now.Add(time.Hour * -2)),
				NotBefore: jwt.NewNumericDate(now.Add(time.Hour * -2)),
				ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * -2).Add(time.Hour)),
				Subject:   user,
			}), secret),
			wantErr: auth.ErrTokenExpired,
		},
		{
			name: "not active token",
			token: mustSign(jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
				IssuedAt:  jwt.NewNumericDate(now.Add(time.Hour * 1)),
				NotBefore: jwt.NewNumericDate(now.Add(time.Hour * 1)),
				ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * 1).Add(time.Hour)),
				Subject:   user,
			}), secret),
			wantErr: testutil.AnyError{},
		},
		{
			name: "missing exp",
			token: mustSign(jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
				IssuedAt:  jwt.NewNumericDate(now),
				NotBefore: jwt.NewNumericDate(now),
				Subject:   user,
			}), secret),
			wantErr: auth.ErrTokenExpired,
		},
		{
			name: "missing nbf",
			token: mustSign(jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
				IssuedAt:  jwt.NewNumericDate(now),
				ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour)),
				Subject:   user,
			}), secret),
			wantErr: testutil.AnyError{},
		},
		{
			name: "wrong method",
			token: mustSign(jwt.NewWithClaims(jwt.SigningMethodNone, jwt.RegisteredClaims{
				IssuedAt:  jwt.NewNumericDate(now),
				NotBefore: jwt.NewNumericDate(now),
				ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * 24)),
				Subject:   user,
			}), jwt.UnsafeAllowNoneSignatureType),
			wantErr: testutil.AnyError{},
		},
		{
			name: "wrong secret",
			token: mustSign(jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
				IssuedAt:  jwt.NewNumericDate(now),
				NotBefore: jwt.NewNumericDate(now),
				ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * 24)),
				Subject:   user,
			}), []byte("fake-not-awesome-secret")),
			wantErr: testutil.AnyError{},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			principal, err := auth.ValidateToken(tc.token)

			if !errors.Is(tc.wantErr, err) {
				t.Fatalf("error: want %#v, got %#v", tc.wantErr, err)
			}
			if err != nil {
				return
			}

			if principal == nil || *principal != wantPrincipal {
				t.Fatalf("principal: want %v. got %v", wantPrincipal, principal)
			}
		})
	}
}
