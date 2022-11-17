package service

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

type ListClaims struct {
	ListID 	 int64 `json:"lid"`
	ReadOnly bool  `json:"ro"`
}

func (*ListClaims) Valid() error {
	return nil
}

//go:generate mockgen -destination mock_list_token_test.go -package service_test . ListTokenProvider
type ListTokenProvider interface {
	GenerateToken(claims ListClaims) (string, error)
	ValidateToken(token string) (ListClaims, error)
}

type listTokenProvider struct {
	secret []byte
}

func NewListTokenProvider(secret []byte) ListTokenProvider {
	ltp := &listTokenProvider{
		secret: make([]byte, len(secret)),
	}
	copy(ltp.secret, secret)
	return ltp
}

func (ltp *listTokenProvider) GenerateToken(claims ListClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	return token.SignedString(ltp.secret)
}

func (ltp *listTokenProvider) ValidateToken(raw string) (ListClaims, error) {
	token, err := jwt.ParseWithClaims(raw, &ListClaims{},
		func(t *jwt.Token) (interface{}, error) { return ltp.secret, nil },
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
	)
	if err != nil {
		return ListClaims{}, err
	}
	claims, ok := token.Claims.(*ListClaims)
	if !ok {
		return ListClaims{}, errors.New("unsupported claims")
	}
	return *claims, nil
}