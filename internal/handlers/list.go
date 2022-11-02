package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/xopoww/wishes/internal/models"
	"github.com/xopoww/wishes/restapi/operations"
)

type (
	OnGetListStartInfo struct {
		ListID    int64
		Principal *models.Principal
	}
	OnGetListDoneInfo struct {
		List  *models.List
		Error error
	}
)

func GetList(t Trace) operations.GetListHandler {
	return operations.GetListHandlerFunc(func(glp operations.GetListParams, p *models.Principal) middleware.Responder {
		onDone := traceOnGetList(t, glp.ID, p)
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
		List      models.List
		Principal *models.Principal
	}
	OnPostListDoneInfo struct {
		ListID int64
		Error  error
	}
)

func PostList(t Trace) operations.PostListHandler {
	return operations.PostListHandlerFunc(func(plp operations.PostListParams, p *models.Principal) middleware.Responder {
		return nil
	})
}

type (
	OnPatchListStartInfo struct {
		ListID    int64
		List      models.List
		Principal *models.Principal
	}
	OnPatchListDoneInfo struct {
		Error error
	}
)

func PatchList(t Trace) operations.PatchListHandler {
	return operations.PatchListHandlerFunc(func(plp operations.PatchListParams, p *models.Principal) middleware.Responder {
		return nil
	})
}

type (
	OnDeleteListStartInfo struct {
		ListID    int64
		Principal *models.Principal
	}
	OnDeleteListDoneInfo struct {
		Error error
	}
)

func DeleteList(t Trace) operations.DeleteListHandler {
	return operations.DeleteListHandlerFunc(func(dlp operations.DeleteListParams, p *models.Principal) middleware.Responder {
		return nil
	})
}

type (
	OnGetUserListsStartInfo struct {
		UserID    int64
		Principal *models.Principal
	}
	OnGetUserListsDoneInfo struct {
		ListIDs []int64
		Error   error
	}
)

func GetUserLists(t Trace) operations.GetUserListsHandler {
	return operations.GetUserListsHandlerFunc(func(gulp operations.GetUserListsParams, p *models.Principal) middleware.Responder {
		return nil
	})
}
