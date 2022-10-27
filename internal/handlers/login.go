package handlers

import (
	"errors"

	"github.com/go-openapi/runtime/middleware"
	"github.com/xopoww/wishes/internal/db"
	"github.com/xopoww/wishes/restapi/operations"
	"github.com/xopoww/wishes/internal/auth"
)

type (
	OnLoginStartInfo struct {
		Username string
	}
	OnLoginDoneInfo struct {
		Ok bool
		Error error
	}
)

func Login(t Trace) operations.LoginHandler {
	return operations.LoginHandlerFunc(func(lp operations.LoginParams) middleware.Responder {
		username := string(*lp.Credentials.Username)
		password := string(*lp.Credentials.Password)

		onDone := traceOnLogin(t, username)
		var (
			ok bool
			err error
		)
		defer func(){
			onDone(ok, err)
		}()

		user, hash, err := db.GetFullUser(username)
		if errors.Is(err, db.ErrNotFound) {
			ok = false
			err = nil
			return operations.NewLoginOK().WithPayload(&operations.LoginOKBody{Ok: &ok})
		}
		if err != nil {
			return operations.NewLoginInternalServerError()
		}

		ok = auth.ComparePassword(password, hash)

		payload := &operations.LoginOKBody{}
		payload.Ok = &ok
		if ok {
			token, err := auth.GenerateToken(user)
			if err != nil {
				return operations.NewLoginInternalServerError()
			}
			payload.Token = token
		}
		return operations.NewLoginOK().WithPayload(payload)
	})
}