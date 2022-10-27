package auth

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/xopoww/wishes/internal/db"
	"github.com/xopoww/wishes/models"
)


const TokenSecretVariableName = "WISHES_JWT_SECRET"
var secret []byte

func getKey() []byte {
	if len(secret) > 0 {
		return secret
	}

	secretString, exists := os.LookupEnv(TokenSecretVariableName)
	if !exists {
		panic(fmt.Sprintf("%s not set", TokenSecretVariableName))
	}

	var err error
	secret, err = base64.RawStdEncoding.DecodeString(secretString)
	if err != nil {
		panic(fmt.Sprintf("base64 decode %s: %s", TokenSecretVariableName, err))
	}

	return secret
}

func GenerateToken(user *db.User) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(now),
		NotBefore: jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * 24)),
		Subject:   user.Name,
	})
	return token.SignedString(getKey())
}

var ErrTokenExpired = errors.New("token expired")

func ValidateToken(raw string) (*models.Principal, error) {
	token, err := jwt.ParseWithClaims(raw, &jwt.RegisteredClaims{},
		func(t *jwt.Token) (interface{}, error) { return getKey(), nil },
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
		jwt.WithoutClaimsValidation(),
	)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, errors.New("unsupported claims")
	}

	now := time.Now()
	if claims.NotBefore == nil || now.Before(claims.NotBefore.Time) {
		return nil, errors.New("token not active yet")
	}
	if claims.ExpiresAt == nil || claims.ExpiresAt.Before(now) {
		return nil, ErrTokenExpired
	}

	principal := models.Principal(claims.Subject)
	return &principal, nil
}
