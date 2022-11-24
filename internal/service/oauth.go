package service

import (
	"context"
	"fmt"

	"github.com/xopoww/wishes/internal/models"
)

func (s *service) AddOAuthProvider(providerId string, op OAuthProvider) {
	if s.ops[providerId] != nil {
		panic(fmt.Sprintf("wishes: duplicate OAuth provider: %q", providerId))
	}
	s.ops[providerId] = op
}

func (s *service) OAuthRegister(ctx context.Context, username, provider, oauthToken string) (int64, error) {
	op := s.ops[provider]
	if op == nil {
		return 0, ErrNoProvider
	}
	eid, err := op.Validate(ctx, oauthToken)
	if err != nil {
		return 0, fmt.Errorf("%w: %s", ErrOAuth, err)
	}

	tx, err := s.r.Begin()
	if err != nil {
		return 0, err
	}
	user, err := tx.AddUser(ctx, &models.User{Name: username})
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}
	err = tx.AddOAuth(ctx, provider, eid, user)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}
	err = tx.Commit()
	return user.ID, err
}

func (s *service) OAuthLogin(ctx context.Context, provider, oauthToken string) (token string, err error) {
	op := s.ops[provider]
	if op == nil {
		return "", ErrNoProvider
	}
	eid, err := op.Validate(ctx, oauthToken)
	if err != nil {
		return "", fmt.Errorf("%w: %s", ErrOAuth, err)
	}

	tx, err := s.r.Begin()
	if err != nil {
		return "", err
	}
	uid, err := tx.CheckOAuth(ctx, provider, eid)
	if err != nil {
		_ = tx.Rollback()
		return "", err
	}
	user, err := tx.GetUser(ctx, uid)
	if err != nil {
		_ = tx.Rollback()
		return "", err
	}
	_ = tx.Commit()
	return s.generateToken(user)
}
