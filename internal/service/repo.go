package service

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

//go:generate mockgen -destination mock_repo_test.go -package service_test . Repository

type Repository interface {
	CheckUsername(ctx context.Context, username string) (int64, error)

	GetUser(ctx context.Context, id int64) (*models.User, error)

	AddUser(ctx context.Context, user *models.User) (*models.User, error)

	EditUser(ctx context.Context, user *models.User) error

	GetUserLists(ctx context.Context, id int64) ([]int64, error)

	GetList(ctx context.Context, id int64) (*models.List, error)

	AddList(ctx context.Context, list *models.List) (*models.List, error)

	EditList(ctx context.Context, list *models.List) error

	DeleteList(ctx context.Context, list *models.List) error

	Close() error
}
