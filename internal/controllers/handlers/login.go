package handlers

import (
	"context"

	"github.com/go-openapi/runtime/middleware"
	"github.com/xopoww/wishes/restapi/operations"
)

type (
	OnLoginStartInfo struct {
		Username string
	}
	OnLoginDoneInfo struct {
		Ok    bool
		Error error
	}
)

func (ac *ApiController) Login() operations.LoginHandler {
	return operations.LoginHandlerFunc(func(lp operations.LoginParams) middleware.Responder {
		username := string(*lp.Credentials.Username)
		password := string(*lp.Credentials.Password)

		onDone := traceOnLogin(ac.t, username)
		var (
			ok  bool
			err error
		)
		defer func() {
			onDone(ok, err)
		}()

		var token string
		token, err = ac.s.Login(context.TODO(), username, password)
		// TODO: handle internal server error
		ok = err == nil

		payload := &operations.LoginOKBody{}
		payload.Ok = &ok
		if ok {
			payload.Token = token
		}
		return operations.NewLoginOK().WithPayload(payload)
	})
}
