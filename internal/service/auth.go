package service

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/xopoww/wishes/internal/models"
	"golang.org/x/crypto/bcrypt"
)

var ErrTokenExpired = errors.New("token expired")

func (s *service) Auth(ctx context.Context, raw string) (*models.User, error) {
	token, err := jwt.ParseWithClaims(raw, &jwtClaims{},
		func(t *jwt.Token) (interface{}, error) { return getKey(), nil },
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
		jwt.WithoutClaimsValidation(),
	)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*jwtClaims)
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

	return &models.User{
		ID:   claims.UserID,
		Name: claims.Subject,
	}, nil
}

func (s *service) Login(ctx context.Context, username string, password string) (token string, ok bool) {
	id, err := s.r.CheckUsername(ctx, username)
	if err != nil {
		return "", false
	}
	user, err := s.r.GetUser(ctx, id)
	if err != nil {
		return "", false
	}
	if !s.comparePassword(password, user.PassHash) {
		return "", false
	}
	token, err = s.generateToken(user)
	return token, err == nil
}

func (s *service) hashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
}

func (s *service) comparePassword(password string, hash []byte) bool {
	return bcrypt.CompareHashAndPassword(hash, []byte(password)) == nil
}

type jwtClaims struct {
	jwt.RegisteredClaims
	UserID int64 `json:"user_id"`
}

func NewJwtClaims(base jwt.RegisteredClaims, userId int64) jwt.Claims {
	return &jwtClaims{
		RegisteredClaims: base,
		UserID:           userId,
	}
}

func (s *service) generateToken(user *models.User) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, NewJwtClaims(
		jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * 24)),
			Subject:   user.Name,
		},
		user.ID,
	))
	return token.SignedString(getKey())
}

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
