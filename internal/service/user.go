package service

import (
	"context"

	"github.com/xopoww/wishes/internal/models"
)

func (s *service) Register(ctx context.Context, username string, password string) (int64, error) {
	hash, err := s.hashPassword(password)
	if err != nil {
		return 0, err
	}
	user := &models.User{
		Name:     username,
		PassHash: hash,
	}
	user, err = s.r.AddUser(ctx, user)
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (s *service) GetUser(ctx context.Context, id int64, client *models.User) (*models.User, error) {
	return s.r.GetUser(ctx, id)
}

func (s *service) EditUser(ctx context.Context, user *models.User, client *models.User) error {
	if user.ID != client.ID {
		return ErrAccessDenied
	}
	return s.r.EditUser(ctx, user)
}
