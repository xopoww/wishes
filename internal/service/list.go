package service

import (
	"context"

	"github.com/xopoww/wishes/internal/models"
)

func (s *service) GetUserLists(ctx context.Context, id int64, client *models.User) ([]int64, error) {
	panic("not implemented") // TODO: Implement
}

func (s *service) GetList(ctx context.Context, id int64, client *models.User) (*models.List, error) {
	panic("not implemented") // TODO: Implement
}

func (s *service) EditList(ctx context.Context, list *models.List, client *models.User) error {
	panic("not implemented") // TODO: Implement
}

func (s *service) AddList(ctx context.Context, list *models.List, client *models.User) (*models.List, error) {
	panic("not implemented") // TODO: Implement
}

func (s *service) DeleteList(ctx context.Context, list *models.List, client *models.User) error {
	panic("not implemented") // TODO: Implement
}
