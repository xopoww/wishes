package service

import (
	"context"

	"github.com/xopoww/wishes/internal/models"
)

func (s *service) GetUserLists(ctx context.Context, id int64, client *models.User) ([]int64, error) {
	return s.r.GetUserLists(ctx, id)
}

func (s *service) GetList(ctx context.Context, id int64, client *models.User) (*models.List, error) {
	list, err := s.r.GetList(ctx, id)
	if err != nil {
		return nil, err
	}
	return s.r.GetListItems(ctx, list)
}

func (s *service) EditList(ctx context.Context, list *models.List, client *models.User) error {
	if err := s.checkWriteAccess(ctx, list, client); err != nil {
		return err
	}
	return s.r.EditList(ctx, list)
}

func (s *service) AddList(ctx context.Context, list *models.List, client *models.User) (*models.List, error) {
	list.OwnerID = client.ID
	return s.r.AddList(ctx, list)
}

func (s *service) DeleteList(ctx context.Context, list *models.List, client *models.User) error {
	if err := s.checkWriteAccess(ctx, list, client); err != nil {
		return err
	}
	return s.r.DeleteList(ctx, list)
}

func (s *service) checkWriteAccess(ctx context.Context, list *models.List, client *models.User) error {
	l, err := s.r.GetList(ctx, list.ID)
	if err != nil {
		return err
	}
	list.OwnerID = l.OwnerID
	if l.OwnerID != client.ID {
		return ErrAccessDenied
	}
	return nil
}