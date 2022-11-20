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
			items []models.ListItem
			err   error
		)
		defer func() { onDone(items, err) }()

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
		items = list.Items
		payload := make([]*operations.GetListItemsOKBodyItemsItems0, len(items))
		for i := range items {
			payload[i] = &operations.GetListItemsOKBodyItemsItems0{
				ID:       *conv.SwagID(items[i].ID),
				ListItem: *conv.SwagItem(items[i]),
			}
		}
		return operations.NewGetListItemsOK().WithPayload(&operations.GetListItemsOKBody{
			Items:    payload,
			Revision: *conv.SwagRevision(list.RevisionID),
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

		return operations.NewPostListCreated().WithPayload(&operations.PostListCreatedBody{
			ID:       *conv.SwagID(list.ID),
			Revision: *conv.SwagRevision(list.RevisionID),
		})
	})
}

type (
	OnPostListItemsStartInfo struct {
		List   *models.List
		Items  []models.ListItem
		Client *models.User
	}
	OnPostListItemsDoneInfo struct {
		Error error
	}
)

func (ac *ApiController) PostListItems() operations.PostListItemsHandler {
	return operations.PostListItemsHandlerFunc(func(params operations.PostListItemsParams, p *apimodels.Principal) middleware.Responder {
		client := conv.Client(p)
		list := &models.List{ID: params.ID, RevisionID: conv.Revision(&params.Items.Revision)}
		items := conv.Items(params.Items.Items)

		onDone := traceOnPostListItems(ac.t, list, items, client)
		var err error
		defer func() { onDone(err) }()

		list, err = ac.s.AddListItems(context.TODO(), list, items, client)
		if errors.Is(err, service.ErrNotFound) {
			return operations.NewPostListItemsNotFound()
		}
		if errors.Is(err, service.ErrAccessDenied) {
			return operations.NewPostListItemsForbidden()
		}
		if errors.Is(err, service.ErrConflict) {
			return operations.NewPostListItemsConflict()
		}
		if err != nil {
			return operations.NewPostListItemsInternalServerError()
		}

		return operations.NewPostListItemsCreated().WithPayload(conv.SwagRevision(list.RevisionID))
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
	OnDeleteListItemsStartInfo struct {
		List    *models.List
		ItemIDs []int64
		Client  *models.User
	}
	OnDeleteListItemsDoneInfo struct {
		Error error
	}
)

func (ac *ApiController) DeleteListItems() operations.DeleteListItemsHandler {
	return operations.DeleteListItemsHandlerFunc(func(params operations.DeleteListItemsParams, p *apimodels.Principal) middleware.Responder {
		client := conv.Client(p)
		list := &models.List{ID: params.ID, RevisionID: params.Rev}
		ids := params.Ids

		onDone := traceOnDeleteListItems(ac.t, list, ids, client)
		var err error
		defer func() { onDone(err) }()

		list, err = ac.s.DeleteListItems(context.TODO(), list, ids, client)
		if errors.Is(err, service.ErrNotFound) {
			return operations.NewDeleteListItemsNotFound()
		}
		if errors.Is(err, service.ErrAccessDenied) {
			return operations.NewDeleteListItemsForbidden()
		}
		if errors.Is(err, service.ErrConflict) {
			return operations.NewDeleteListItemsConflict()
		}
		if err != nil {
			return operations.NewDeleteListItemsInternalServerError()
		}

		return operations.NewDeleteListItemsOK().WithPayload(conv.SwagRevision(list.RevisionID))
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

type (
	OnPostItemTakenStartInfo struct {
		List   *models.List
		ItemID int64
		Client *models.User
		Token  *string
	}
	OnPostItemTakenDoneInfo struct {
		Error error
	}
)

func (ac *ApiController) PostItemTaken() operations.PostItemTakenHandler {
	return operations.PostItemTakenHandlerFunc(func(params operations.PostItemTakenParams, p *apimodels.Principal) middleware.Responder {
		client := conv.Client(p)
		list := &models.List{ID: params.ID, RevisionID: conv.Revision(params.Body)}
		itemId := params.ItemID
		token := params.AccessToken

		onDone := traceOnPostItemTaken(ac.t, list, itemId, client, token)
		var err error
		defer func() { onDone(err) }()

		err = ac.s.TakeItem(context.TODO(), list, itemId, client, token)
		if errors.Is(err, service.ErrNotFound) {
			return operations.NewPostItemTakenNotFound()
		}
		if errors.Is(err, service.ErrAccessDenied) {
			return operations.NewPostItemTakenForbidden()
		}
		if errors.Is(err, service.ErrConflict) {
			payload := &operations.PostItemTakenConflictBody{}
			var errat service.ErrAlreadyTaken
			if errors.As(err, &errat) {
				reason := "already taken"
				payload.Reason = &reason
				payload.TakenBy = errat.TakenBy
			} else {
				reason := "outdated revision"
				payload.Reason = &reason
			}
			return operations.NewPostItemTakenConflict().WithPayload(payload)
		}
		if err != nil {
			return operations.NewPostItemTakenInternalServerError()
		}
		return operations.NewPostItemTakenNoContent()
	})
}

type (
	OnDeleteItemTakenStartInfo struct {
		List   *models.List
		ItemID int64
		Client *models.User
		Token  *string
	}
	OnDeleteItemTakenDoneInfo struct {
		Error error
	}
)

func (ac *ApiController) DeleteItemTaken() operations.DeleteItemTakenHandler {
	return operations.DeleteItemTakenHandlerFunc(func(params operations.DeleteItemTakenParams, p *apimodels.Principal) middleware.Responder {
		client := conv.Client(p)
		list := &models.List{ID: params.ID, RevisionID: params.Rev}
		itemId := params.ItemID
		token := params.AccessToken

		onDone := traceOnDeleteItemTaken(ac.t, list, itemId, client, token)
		var err error
		defer func() { onDone(err) }()

		err = ac.s.UntakeItem(context.TODO(), list, itemId, client, token)
		if errors.Is(err, service.ErrNotFound) {
			return operations.NewDeleteItemTakenNotFound()
		}
		if errors.Is(err, service.ErrAccessDenied) {
			return operations.NewDeleteItemTakenForbidden()
		}
		if errors.Is(err, service.ErrConflict) {
			var reason string
			if errors.Is(err, service.ErrOutdated) {
				reason = "outdated revision"
			} else {
				reason = "not taken"
			}
			return operations.NewDeleteItemTakenConflict().WithPayload(&operations.DeleteItemTakenConflictBody{Reason: &reason})
		}
		if err != nil {
			return operations.NewDeleteItemTakenInternalServerError()
		}
		return operations.NewDeleteItemTakenNoContent()
	})
}
