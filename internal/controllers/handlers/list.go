package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/xopoww/wishes/internal/models"
	"github.com/xopoww/wishes/restapi/apimodels"
	"github.com/xopoww/wishes/restapi/apimodels/conv"
	"github.com/xopoww/wishes/restapi/operations"
)

type (
	OnGetListStartInfo struct {
		ListID int64
		Client *models.User
	}
	OnGetListDoneInfo struct {
		List  *models.List
		Error error
	}
)

func (ac *ApiController) GetList() operations.GetListHandler {
	return operations.GetListHandlerFunc(func(glp operations.GetListParams, p *apimodels.Principal) middleware.Responder {
		client := conv.Client(p)
		onDone := traceOnGetList(ac.t, glp.ID, client)
		var (
			payload *models.List
			err     error
		)
		defer func() { onDone(payload, err) }()

		return nil
	})
}

type (
	OnPostListStartInfo struct {
		List   models.List
		Client models.User
	}
	OnPostListDoneInfo struct {
		ListID int64
		Error  error
	}
)

func (ac *ApiController) PostList() operations.PostListHandler {
	return operations.PostListHandlerFunc(func(plp operations.PostListParams, p *apimodels.Principal) middleware.Responder {
		return nil
	})
}

type (
	OnPatchListStartInfo struct {
		List   models.List
		Client models.User
	}
	OnPatchListDoneInfo struct {
		Error error
	}
)

func (ac *ApiController) PatchList() operations.PatchListHandler {
	return operations.PatchListHandlerFunc(func(plp operations.PatchListParams, p *apimodels.Principal) middleware.Responder {
		return nil
	})
}

type (
	OnDeleteListStartInfo struct {
		ListID int64
		Client models.User
	}
	OnDeleteListDoneInfo struct {
		Error error
	}
)

func (ac *ApiController) DeleteList() operations.DeleteListHandler {
	return operations.DeleteListHandlerFunc(func(dlp operations.DeleteListParams, p *apimodels.Principal) middleware.Responder {
		return nil
	})
}

type (
	OnGetUserListsStartInfo struct {
		UserID int64
		Client models.User
	}
	OnGetUserListsDoneInfo struct {
		ListIDs []int64
		Error   error
	}
)

func (ac *ApiController) GetUserLists() operations.GetUserListsHandler {
	return operations.GetUserListsHandlerFunc(func(gulp operations.GetUserListsParams, p *apimodels.Principal) middleware.Responder {
		return nil
	})
}
