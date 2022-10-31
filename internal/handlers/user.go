package handlers

import (
	"errors"

	"github.com/go-openapi/runtime/middleware"
	"github.com/xopoww/wishes/internal/auth"
	"github.com/xopoww/wishes/internal/db"
	"github.com/xopoww/wishes/models"
	"github.com/xopoww/wishes/restapi/operations"
)

type (
	OnGetUserStartInfo struct {
		UserID	  int
		Principal *models.Principal
	}
	OnGetUserDoneInfo struct {
		User	*models.User
		Error 	error
	}
)

func GetUser(t Trace) operations.GetUserHandler {
	return operations.GetUserHandlerFunc(func(gup operations.GetUserParams, p *models.Principal) middleware.Responder {
		id := int(gup.ID)

		onDone := traceOnGetUser(t, id, p)
		payload := &models.User{}
		var err error
		defer func() {
			onDone(payload, err)
		}()

		user, err := db.GetUserById(id)
		if errors.Is(err, db.ErrNotFound) {
			return operations.NewGetUserNotFound()
		}
		if err != nil {
			return operations.NewGetUserInternalServerError()
		}

		payload.ID = &gup.ID
		payload.Username = models.UserName(user.Name)
		payload.Fname = user.FirstName
		payload.Lname = user.LastName
		return operations.NewGetUserOK().WithPayload(payload)
	})
}

type (
	OnPatchUserStartInfo struct {
		User	  *models.User
		Principal *models.Principal
	}

	OnPatchUserDoneInfo struct {
		Error error
	}
)

func PatchUser(t Trace) operations.PatchUserHandler {
	return operations.PatchUserHandlerFunc(func(pup operations.PatchUserParams, p *models.Principal) middleware.Responder {

		onDone := traceOnPatchUser(t, pup.User, p)
		var err error
		defer func() {
			onDone(err)
		}()


		if string(pup.User.Username) != string(*p) {
			err = errors.New("forbidden")
			return operations.NewPatchUserForbidden()
		}

		user := &db.User{
			ID:		   int(*pup.User.ID),
			Name:      string(pup.User.Username),
			FirstName: pup.User.Fname,
			LastName:  pup.User.Lname,
		}
		err = db.EditUserInfo(user)
		if errors.Is(err, db.ErrNotFound) {
			return operations.NewPatchUserNotFound()
		}
		if err != nil {
			return operations.NewPatchUserInternalServerError()
		}

		return operations.NewPatchUserOK()
	})
}

type (
	OnPostUserStartInfo struct {
		Username string
	}
	OnPostUserDoneInfo struct {
		Ok    bool
		Error error
	}
)

func PostUser(t Trace) operations.PostUserHandler {
	return operations.PostUserHandlerFunc(func(pup operations.PostUserParams) middleware.Responder {
		username := string(*pup.Credentials.Username)
		password := string(*pup.Credentials.Password)

		onDone := traceOnPostUser(t, username)
		var (
			ok  bool
			err error
		)
		defer func() {
			onDone(ok, err)
		}()

		hash, err := auth.HashPassword(password)
		if err != nil {
			return operations.NewPostUserInternalServerError()
		}

		_, err = db.AddUser(username, hash)
		if err != nil && !errors.Is(err, db.ErrNameTaken) {
			return operations.NewPostUserInternalServerError()
		}

		ok = err == nil
		payload := &operations.PostUserOKBody{
			Ok: &ok,
		}
		if !ok {
			payload.Error = err.Error()
		}
		return operations.NewPostUserOK().WithPayload(payload)
	})
}
