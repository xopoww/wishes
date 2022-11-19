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
		Token  *string
	}
	OnGetListDoneInfo struct {
		List  *models.List
		Error error
	}
)

func (ac *ApiController) GetList() operations.GetListHandler {
	return operations.GetListHandlerFunc(func(glp operations.GetListParams, p *apimodels.Principal) middleware.Responder {
		client := conv.Client(p)
		token := glp.AccessToken

		onDone := traceOnGetList(ac.t, glp.ID, client, token)
		var (
			payload *models.List
			err     error
		)
		defer func() { onDone(payload, err) }()

		payload, err = ac.s.GetList(context.TODO(), glp.ID, client, token)
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
	OnGetListItemsStartInfo struct {
		ListID int64
		Client *models.User
		Token  *string
	}
	OnGetListItemsDoneInfo struct {
		Items []models.ListItem
		Error error
	}
)

func (ac *ApiController) GetListItems() operations.GetListItemsHandler {
	return operations.GetListItemsHandlerFunc(func(glip operations.GetListItemsParams, p *apimodels.Principal) middleware.Responder {
		client := conv.Client(p)
		token := glip.AccessToken

		onDone := traceOnGetListItems(ac.t, glip.ID, client, token)
		var (
			payload []models.ListItem
			err     error
		)
		defer func() { onDone(payload, err) }()

		list, err := ac.s.GetListItems(context.TODO(), &models.List{ID: glip.ID}, client, token)
		if errors.Is(err, service.ErrAccessDenied) {
			return operations.NewGetListItemsForbidden()
		}
		if errors.Is(err, service.ErrNotFound) {
			return operations.NewGetListItemsNotFound()
		}
		if err != nil {
			return operations.NewGetListItemsInternalServerError()
		}
		payload = list.Items
		return operations.NewGetListItemsOK().WithPayload(&operations.GetListItemsOKBody{
			// ListItems: apimodels.ListItems{Items: conv.SwagItems(payload)},
			Revision:  conv.SwagRevision(list.RevisionID),
		})
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
		list := conv.List(&plp.List.List)
		list.Items = conv.Items(plp.List.Items)

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
		list := conv.List(&plp.List.List)
		// list.Items = conv.Items(plp.List.Items)
		list.ID = plp.ID

		onDone := traceOnPatchList(ac.t, list, client)
		var err error
		defer func() { onDone(err) }()

		_, err = ac.s.EditList(context.TODO(), list, client)
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
		var UserID int64
		if gulp.UserID != nil {
			UserID = *gulp.UserID
		} else {
			UserID = client.ID
		}

		onDone := traceOnGetUserLists(ac.t, UserID, client)
		var (
			lids []int64
			err  error
		)
		defer func() { onDone(lids, err) }()

		lids, err = ac.s.GetUserLists(context.TODO(), UserID, client)
		if errors.Is(err, service.ErrNotFound) {
			return operations.NewGetUserListsNotFound()
		}
		if err != nil {
			return operations.NewGetUserListsInternalServerError()
		}
		return operations.NewGetUserListsOK().WithPayload(lids)
	})
}

type (
	OnGetListTokenStartInfo struct {
		ListID int64
		Client *models.User
	}
	OnGetListTokenDoneInfo struct {
		Error error
	}
)

func (ac *ApiController) GetListToken() operations.GetListTokenHandler {
	return operations.GetListTokenHandlerFunc(func(gltp operations.GetListTokenParams, p *apimodels.Principal) middleware.Responder {
		client := conv.Client(p)

		onDone := traceOnGetListToken(ac.t, gltp.ID, client)
		var err error
		defer func() { onDone(err) }()

		token, err := ac.s.GetListToken(context.TODO(), gltp.ID, client)
		if errors.Is(err, service.ErrNotFound) {
			return operations.NewGetListTokenNotFound()
		}
		if errors.Is(err, service.ErrAccessDenied) {
			return operations.NewGetListTokenForbidden()
		}
		if err != nil {
			return operations.NewGetListTokenInternalServerError()
		}
		return operations.NewGetListTokenOK().WithPayload(&operations.GetListTokenOKBody{Token: token})
	})
}
