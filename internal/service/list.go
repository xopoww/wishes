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
	list, err := s.r.GetList(ctx, id)
	if err != nil {
		return nil, err
	}
	err = s.checkReadAccess(list, client, token)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s *service) GetListItems(ctx context.Context, list *models.List, client *models.User, token *string) (*models.List, error) {
	tx, err := s.r.Begin()
	if err != nil {
		return nil, err
	}
	list, err = tx.GetList(ctx, list.ID)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}
	err = s.checkReadAccess(list, client, token)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}
	list.Items, err = tx.GetListItems(ctx, list)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}
	if s.checkWriteAccess(list, client) == nil {
		for i := range list.Items {
			list.Items[i].TakenBy = nil
		}
	}
	_ = tx.Commit()
	return list, nil
}

func (s *service) EditList(ctx context.Context, list *models.List, client *models.User) (*models.List, error) {
	tx, err := s.r.Begin()
	if err != nil {
		return nil, err
	}
	l, err := tx.GetList(ctx, list.ID)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}
	err = s.checkWriteAccess(l, client)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}
	list, err = tx.EditList(ctx, list)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}
	return list, tx.Commit()
}

func (s *service) AddList(ctx context.Context, list *models.List, client *models.User) (*models.List, error) {
	list.OwnerID = client.ID
	return s.r.AddList(ctx, list)
}

func (s *service) AddListItems(ctx context.Context, list *models.List, items []models.ListItem, client *models.User) (*models.List, error) {
	tx, err := s.r.Begin()
	if err != nil {
		return nil, err
	}
	l, err := tx.GetList(ctx, list.ID)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}
	err = s.checkWriteAccess(l, client)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}
	if list.RevisionID < l.RevisionID {
		_ = tx.Rollback()
		return nil, ErrOutdated
	}
	_, err = tx.AddListItems(ctx, list, items)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}
	list.RevisionID++
	_, err = tx.EditList(ctx, list)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}
	return list, tx.Commit()
}

func (s *service) DeleteList(ctx context.Context, list *models.List, client *models.User) error {
	tx, err := s.r.Begin()
	if err != nil {
		return err
	}
	list, err = tx.GetList(ctx, list.ID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	err = s.checkWriteAccess(list, client)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	err = tx.DeleteList(ctx, list)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (s *service) DeleteListItems(ctx context.Context, list *models.List, ids []int64, client *models.User) (*models.List, error) {
	tx, err := s.r.Begin()
	if err != nil {
		return nil, err
	}
	l, err := tx.GetList(ctx, list.ID)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}
	err = s.checkWriteAccess(l, client)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}
	if list.RevisionID < l.RevisionID {
		_ = tx.Rollback()
		return nil, ErrOutdated
	}
	err = tx.DeleteListItems(ctx, list, ids)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}
	list.RevisionID++
	_, err = tx.EditList(ctx, list)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}
	return list, tx.Commit()
}

func (s *service) GetListToken(ctx context.Context, id int64, client *models.User) (string, error) {
	list, err := s.r.GetList(ctx, id)
	if err != nil {
		return "", err
	}
	err = s.checkWriteAccess(list, client)
	if err != nil {
		return "", err
	}
	return s.ltp.GenerateToken(ListClaims{ListID: id})
}

func (s *service) TakeItem(ctx context.Context, list *models.List, itemId int64, client *models.User, token *string) error {
	tx, err := s.r.Begin()
	if err != nil {
		return err
	}
	l, err := tx.GetList(ctx, list.ID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	err = s.checkReadAccess(l, client, token)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	if s.checkWriteAccess(l, client) == nil {
		_ = tx.Rollback()
		return ErrAccessDenied
	}
	if list.RevisionID < l.RevisionID {
		_ = tx.Rollback()
		return ErrOutdated
	}
	taken, err := tx.GetItemTaken(ctx, list.ID, itemId)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	if taken != nil {
		_ = tx.Rollback()
		return ErrConflict
	}
	err = tx.SetItemTaken(ctx, list.ID, itemId, &client.ID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	err = tx.Commit()
	return err
}

func (s *service) UntakeItem(ctx context.Context, list *models.List, itemId int64, client *models.User, token *string) error {
	tx, err := s.r.Begin()
	if err != nil {
		return err
	}
	l, err := tx.GetList(ctx, list.ID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	err = s.checkReadAccess(l, client, token)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	if s.checkWriteAccess(l, client) == nil {
		_ = tx.Rollback()
		return ErrAccessDenied
	}
	if list.RevisionID < l.RevisionID {
		_ = tx.Rollback()
		return ErrOutdated
	}
	taken, err := tx.GetItemTaken(ctx, list.ID, itemId)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	if taken == nil || *taken != client.ID {
		_ = tx.Rollback()
		return ErrConflict
	}
	err = tx.SetItemTaken(ctx, list.ID, itemId, nil)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	err = tx.Commit()
	return err
}

func (s *service) checkWriteAccess(list *models.List, client *models.User) error {
	if list.OwnerID != client.ID {
		return ErrAccessDenied
	}
	return nil
}

func (s *service) checkReadAccess(list *models.List, client *models.User, token *string) error {
	if token != nil {
		if list.Access == models.PrivateAccess {
			return ErrAccessDenied
		}
		claims, err := s.ltp.ValidateToken(*token)
		if err != nil {
			return fmt.Errorf("%w (%s)", ErrAccessDenied, err)
		}
		if claims.ListID != list.ID {
			return fmt.Errorf("%w (token for %d)", ErrAccessDenied, list.ID)
		}
		return nil
	}
	if list.Access != models.PublicAccess && list.OwnerID != client.ID {
		return ErrAccessDenied
	}
	return nil
}
