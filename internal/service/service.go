package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/xopoww/wishes/internal/models"
	"github.com/xopoww/wishes/internal/service/repository"
)

var (
	ErrAccessDenied = errors.New("access denied")

	ErrNotFound = repository.ErrNotFound
	ErrConflict = repository.ErrConflict
	ErrOutdated = fmt.Errorf("%w: outdated revision", ErrConflict)

	ErrOAuth    	= errors.New("invalid oauth credentials")
	ErrNoProvider	= fmt.Errorf("%w: unknown provider", ErrOAuth)
)

type ErrAlreadyTaken struct {
	TakenBy int64
}

func (err ErrAlreadyTaken) Error() string {
	return fmt.Sprintf("%s: already taken by %d", ErrConflict, err.TakenBy)
}

func (err ErrAlreadyTaken) Unwrap() error {
	return ErrConflict
}

type Service interface {
	Auth(ctx context.Context, token string) (*models.User, error)

	// Login returns ErrAccessDenied on any problem with login credentials (i.e. it
	// does not distinguish between bad username or bad password). Other errors indicate
	// internal server-side error.
	Login(ctx context.Context, username, password string) (token string, err error)

	Register(ctx context.Context, username, password string) (int64, error)

	GetUser(ctx context.Context, id int64, client *models.User) (*models.User, error)

	EditUser(ctx context.Context, user, client *models.User) error

	GetUserLists(ctx context.Context, id int64, client *models.User) ([]int64, error)

	GetList(ctx context.Context, id int64, client *models.User, token *string) (*models.List, error)

	GetListItems(ctx context.Context, list *models.List, client *models.User, token *string) (*models.List, error)

	EditList(ctx context.Context, list *models.List, client *models.User) (*models.List, error)

	AddList(ctx context.Context, list *models.List, client *models.User) (*models.List, error)

	AddListItems(ctx context.Context, list *models.List, items []models.ListItem, client *models.User) (*models.List, error)

	GetListToken(ctx context.Context, id int64, client *models.User) (string, error)

	DeleteList(ctx context.Context, list *models.List, client *models.User) error

	DeleteListItems(ctx context.Context, list *models.List, ids []int64, client *models.User) (*models.List, error)

	TakeItem(ctx context.Context, list *models.List, itemId int64, client *models.User, token *string) error

	UntakeItem(ctx context.Context, list *models.List, itemId int64, client *models.User, token *string) error

	// AddOAuthProvider registers new provider. It is not safe for concurrent use with other methods
	// and should be called only during configuration.
	AddOAuthProvider(providerId string, op OAuthProvider)

	OAuthRegister(ctx context.Context, username, provider, oauthToken string) (int64, error)

	OAuthLogin(ctx context.Context, provider, oauthToken string) (token string, err error)
}

type service struct {
	r   repository.Repository
	ltp ListTokenProvider
	ops map[string]OAuthProvider
}

func NewService(r repository.Repository, ltp ListTokenProvider) Service {
	return &service{r: r, ltp: ltp, ops: make(map[string]OAuthProvider)}
}

//go:generate mockgen -destination mock_oauth_test.go -package service_test . OAuthProvider
type OAuthProvider interface {
	Validate(ctx context.Context, token string) (extId string, err error)
}