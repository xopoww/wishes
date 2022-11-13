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
	OnGetUserStartInfo struct {
		UserID int64
		Client *models.User
	}
	OnGetUserDoneInfo struct {
		User  *models.User
		Error error
	}
)

func (ac *ApiController) GetUser() operations.GetUserHandler {
	return operations.GetUserHandlerFunc(func(gup operations.GetUserParams, p *apimodels.Principal) middleware.Responder {
		id := gup.ID
		client := conv.Client(p)

		onDone := traceOnGetUser(ac.t, id, client)
		user := &models.User{}
		var err error
		defer func() {
			onDone(user, err)
		}()

		user, err = ac.s.GetUser(context.TODO(), id, client)
		if errors.Is(err, service.ErrNotFound) {
			return operations.NewGetUserNotFound()
		}
		if err != nil {
			return operations.NewGetUserInternalServerError()
		}

		return operations.NewGetUserOK().WithPayload(conv.SwagUser(user))
	})
}

type (
	OnPatchUserStartInfo struct {
		User   *models.User
		Client *models.User
	}

	OnPatchUserDoneInfo struct {
		Error error
	}
)

func (ac *ApiController) PatchUser() operations.PatchUserHandler {
	return operations.PatchUserHandlerFunc(func(pup operations.PatchUserParams, p *apimodels.Principal) middleware.Responder {
		client := conv.Client(p)
		user := conv.User(pup.User)
		user.ID = pup.ID

		onDone := traceOnPatchUser(ac.t, user, client)
		var err error
		defer func() {
			onDone(err)
		}()

		if client.ID != user.ID {
			err = errors.New("forbidden")
			return operations.NewPatchUserForbidden()
		}

		err = ac.s.EditUser(context.TODO(), user, client)
		if err != nil {
			return operations.NewPatchUserInternalServerError()
		}

		return operations.NewPatchUserNoContent()
	})
}

type (
	OnRegisterStartInfo struct {
		Username string
	}
	OnRegisterDoneInfo struct {
		Ok    bool
		Error error
	}
)

func (ac *ApiController) Register() operations.RegisterHandler {
	return operations.RegisterHandlerFunc(func(pup operations.RegisterParams) middleware.Responder {
		username := string(*pup.Credentials.Username)
		password := string(*pup.Credentials.Password)

		onDone := traceOnRegister(ac.t, username)
		var (
			ok  bool
			err error
		)
		defer func() {
			onDone(ok, err)
		}()

		id, err := ac.s.Register(context.TODO(), username, password)

		ok = err == nil
		payload := &operations.RegisterOKBody{
			Ok: &ok,
		}
		if !ok {
			payload.Error = err.Error()
		} else {
			payload.User = conv.SwagID(id)
		}
		return operations.NewRegisterOK().WithPayload(payload)
	})
}
