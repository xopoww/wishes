package handlers

import (
	"context"
	"errors"

	"github.com/go-openapi/runtime/middleware"
	"github.com/xopoww/wishes/internal/models"
	"github.com/xopoww/wishes/internal/service"
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
		
		payload, err = ac.s.GetList(context.TODO(), glp.ID, client)
		if errors.Is(err, service.ErrNotFound) {
			return operations.NewGetListNotFound()
		}
		if errors.Is(err, service.ErrAccessDenied) {
			return operations.NewGetListForbidden()
		}
		if err != nil {
			return operations.NewGetListInternalServerError()
		}
		return operations.NewGetListOK().WithPayload(conv.SwagList(payload))
	})
}

type (
	OnPostListStartInfo struct {
		List   *models.List
		Client *models.User
	}
	OnPostListDoneInfo struct {
		ListID int64
		Error  error
	}
)

func (ac *ApiController) PostList() operations.PostListHandler {
	return operations.PostListHandlerFunc(func(plp operations.PostListParams, p *apimodels.Principal) middleware.Responder {
		client := conv.Client(p)
		list := conv.List(plp.List)
		
		onDone := traceOnPostList(ac.t, list, client)
		var err error
		defer func() { onDone(list.ID, err) }()

		list, err = ac.s.AddList(context.TODO(), list, client)
		if err != nil {
			return operations.NewPostListInternalServerError()
		}

		return operations.NewPostListCreated().WithPayload(conv.SwagID(list.ID))
	})
}

type (
	OnPatchListStartInfo struct {
		List   *models.List
		Client *models.User
	}
	OnPatchListDoneInfo struct {
		Error error
	}
)

func (ac *ApiController) PatchList() operations.PatchListHandler {
	return operations.PatchListHandlerFunc(func(plp operations.PatchListParams, p *apimodels.Principal) middleware.Responder {
		client := conv.Client(p)
		list := conv.List(plp.List)
		list.ID = plp.ID

		onDone := traceOnPatchList(ac.t, list, client)
		var err error
		defer func() { onDone(err) }()

		err = ac.s.EditList(context.TODO(), list, client)
		if errors.Is(err, service.ErrNotFound) {
			return operations.NewPatchListNotFound()
		}
		if errors.Is(err, service.ErrAccessDenied) {
			return operations.NewPatchListForbidden()
		}
		if err != nil {
			return operations.NewPatchListInternalServerError()
		}
		return operations.NewPatchListNoContent()
	})
}

type (
	OnDeleteListStartInfo struct {
		List   *models.List
		Client *models.User
	}
	OnDeleteListDoneInfo struct {
		Error error
	}
)

func (ac *ApiController) DeleteList() operations.DeleteListHandler {
	return operations.DeleteListHandlerFunc(func(dlp operations.DeleteListParams, p *apimodels.Principal) middleware.Responder {
		client := conv.Client(p)
		list := &models.List{ID: dlp.ID}

		onDone := traceOnDeleteList(ac.t, list, client)
		var err error
		defer func() { onDone(err) }()

		err = ac.s.DeleteList(context.TODO(), list, client)
		if errors.Is(err, service.ErrNotFound) {
			return operations.NewDeleteListNotFound()
		}
		if errors.Is(err, service.ErrAccessDenied) {
			return operations.NewDeleteListForbidden()
		}
		if err != nil {
			return operations.NewDeleteListInternalServerError()
		}
		return operations.NewDeleteListNoContent()
	})
}

type (
	OnGetUserListsStartInfo struct {
		UserID int64
		Client *models.User
	}
	OnGetUserListsDoneInfo struct {
		ListIDs []int64
		Error   error
	}
)

func (ac *ApiController) GetUserLists() operations.GetUserListsHandler {
	return operations.GetUserListsHandlerFunc(func(gulp operations.GetUserListsParams, p *apimodels.Principal) middleware.Responder {
		client := conv.Client(p)

		onDone := traceOnGetUserLists(ac.t, gulp.ID, client)
		var (
			lids []int64
			err  error
		)
		defer func() { onDone(lids, err) }()

		lids, err = ac.s.GetUserLists(context.TODO(), gulp.ID, client)
		if errors.Is(err, service.ErrNotFound) {
			return operations.NewGetUserListsNotFound()
		}
		if err != nil {
			return operations.NewGetUserListsInternalServerError()
		}
		return operations.NewGetUserListsOK().WithPayload(lids)
	})
}
