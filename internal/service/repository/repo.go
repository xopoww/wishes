package repository

import (
	"context"
	"errors"

	"github.com/xopoww/wishes/internal/models"

	_ "github.com/golang/mock/mockgen/model"
)

var (
	ErrNotFound = errors.New("not found")
	ErrConflict = errors.New("conflict")
)

type Handle interface {
	CheckUsername(ctx context.Context, username string) (int64, error)

	GetUser(ctx context.Context, id int64) (*models.User, error)

	AddUser(ctx context.Context, user *models.User) (*models.User, error)

	EditUser(ctx context.Context, user *models.User) error

	GetUserLists(ctx context.Context, id int64, publicOnly bool) ([]int64, error)

	// GetList gets only List header (i.e. it does not get ListItems)
	GetList(ctx context.Context, id int64) (*models.List, error)

	GetListItems(ctx context.Context, list *models.List) (*models.List, error)

	AddList(ctx context.Context, list *models.List) (*models.List, error)

	EditList(ctx context.Context, list *models.List) (*models.List, error)

	DeleteList(ctx context.Context, list *models.List) error
}

type Transaction interface {
	Handle

	Commit() error
	Rollback() error
}

//go:generate mockgen -destination ../mock_repo_test.go -package service_test . Repository

type Repository interface {
	Handle

	Begin() (Transaction, error)

	Close() error
}
