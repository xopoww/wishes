package handlers

import (
	"errors"

	"github.com/go-openapi/runtime/middleware"
	"github.com/xopoww/wishes/internal/db"
	"github.com/xopoww/wishes/restapi/operations"
	"github.com/xopoww/wishes/internal/auth"
)

func Login() operations.LoginHandler {
	return operations.LoginHandlerFunc(func(lp operations.LoginParams) middleware.Responder {
		username := string(*lp.Login.Username)
		password := string(*lp.Login.Password)

		user, hash, err := db.GetFullUser(username)
		if errors.Is(err, db.ErrNotFound) {
			ok := false
			return operations.NewLoginOK().WithPayload(&operations.LoginOKBody{Ok: &ok})
		}
		if err != nil {
			return operations.NewLoginInternalServerError().WithPayload(err.Error())
		}

		ok := auth.ComparePassword(password, hash)

		payload := &operations.LoginOKBody{}
		payload.Ok = &ok
		if ok {
			token, err := auth.GenerateToken(user)
			if err != nil {
				return operations.NewLoginInternalServerError().WithPayload(err.Error())
			}
			payload.Token = token
		}
		return operations.NewLoginOK().WithPayload(payload)
	})
}