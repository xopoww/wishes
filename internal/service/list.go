package service

import (
	"context"
	"fmt"

	"github.com/xopoww/wishes/internal/models"
)

func (s *service) GetUserLists(ctx context.Context, id int64, client *models.User) ([]int64, error) {
	publicOnly := client.ID != id
	return s.r.GetUserLists(ctx, id, publicOnly)
}

func (s *service) GetList(ctx context.Context, id int64, client *models.User, token *string) (*models.List, error) {
	list := &models.List{ID: id}
	err := s.checkReadAccess(ctx, list, client, token)
	return list, err
}

func (s *service) GetListItems(ctx context.Context, list *models.List, client *models.User, token *string) (*models.List, error) {
	if err := s.checkReadAccess(ctx, list, client, token); err != nil {
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

func (s *service) GetListToken(ctx context.Context, id int64, client *models.User) (string, error) {
	if err := s.checkWriteAccess(ctx, &models.List{ID: id}, client); err != nil {
		return "", err
	}
	return s.ltp.GenerateToken(ListClaims{ListID: id})
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

func (s *service) checkReadAccess(ctx context.Context, list *models.List, client *models.User, token *string) error {
	l, err := s.r.GetList(ctx, list.ID)
	if err != nil {
		return err
	}
	list.Title = l.Title
	list.OwnerID = l.OwnerID
	list.Access = l.Access
	if token != nil {
		if l.Access == models.PrivateAccess {
			return ErrAccessDenied
		}
		claims, err := s.ltp.ValidateToken(*token)
		if err != nil {
			return fmt.Errorf("%w (%s)", ErrAccessDenied, err)
		}
		if claims.ListID != l.ID {
			return fmt.Errorf("%w (token for %d)", ErrAccessDenied, l.ID)
		}
		return nil
	}
	if l.Access != models.PublicAccess && l.OwnerID != client.ID {
		return ErrAccessDenied
	}
	return nil
}