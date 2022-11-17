package service

import (
	"context"
	"errors"

	"github.com/xopoww/wishes/internal/models"
)

var (
	ErrAccessDenied = errors.New("access denied")
)

type Service interface {
	Auth(ctx context.Context, token string) (*models.User, error)

	Login(ctx context.Context, username, password string) (token string, ok bool)

	Register(ctx context.Context, username, password string) (int64, error)

	GetUser(ctx context.Context, id int64, client *models.User) (*models.User, error)

	EditUser(ctx context.Context, user, client *models.User) error

	GetUserLists(ctx context.Context, id int64, client *models.User) ([]int64, error)

	GetList(ctx context.Context, id int64, client *models.User, token *string) (*models.List, error)

	GetListItems(ctx context.Context, list *models.List, client *models.User, token *string) (*models.List, error)

	EditList(ctx context.Context, list *models.List, client *models.User) error

	AddList(ctx context.Context, list *models.List, client *models.User) (*models.List, error)

	GetListToken(ctx context.Context, id int64, client *models.User) (string, error)

	DeleteList(ctx context.Context, list *models.List, client *models.User) error
}

type service struct {
	r  Repository
	ltp ListTokenProvider
}

func NewService(r Repository, ltp ListTokenProvider) Service {
	return &service{r: r, ltp: ltp}
}
