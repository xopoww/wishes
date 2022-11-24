package handlers

import (
	"context"
	"errors"

	"github.com/go-openapi/runtime/middleware"
	"github.com/xopoww/wishes/internal/service"
	"github.com/xopoww/wishes/restapi/apimodels/conv"
	"github.com/xopoww/wishes/restapi/operations"
)

type (
	OnOAuthRegisterStartInfo struct {
		Username string
		Provider string
	}
	OnOAuthRegisterDoneInfo struct {
		Error error
	}
)

func (ac *ApiController) OAuthRegister() operations.OAuthRegisterHandler {
	return operations.OAuthRegisterHandlerFunc(func(params operations.OAuthRegisterParams) middleware.Responder {
		username := string(*params.Body.Username)
		provider := *params.Body.ProviderID
		token := *params.Body.Token

		onDone := traceOnOAuthRegister(ac.t, username, provider)
		var err error
		defer func() { onDone(err) }()

		uid, err := ac.s.OAuthRegister(context.TODO(), username, provider, token)
		if err == nil || errors.Is(err, service.ErrOAuth) || errors.Is(err, service.ErrConflict) {
			ok := err == nil
			payload := &operations.OAuthRegisterOKBody{Ok: &ok}
			if ok {
				payload.User = conv.SwagID(uid)
			} else {
				payload.Error = err.Error()
			}
			return operations.NewOAuthRegisterOK().WithPayload(payload)
		} else {
			return operations.NewOAuthRegisterInternalServerError()
		}
	})
}

type (
	OnOAuthLoginStartInfo struct {
		Provider string
	}
	OnOAuthLoginDoneInfo struct {
		Error error
	}
)

func (ac *ApiController) OAuthLogin() operations.OAuthLoginHandler {
	return operations.OAuthLoginHandlerFunc(func(params operations.OAuthLoginParams) middleware.Responder {
		provider := *params.Body.ProviderID
		token := *params.Body.Token

		onDone := traceOnOAuthLogin(ac.t, provider)
		var err error
		defer func() { onDone(err) }()

		tok, err := ac.s.OAuthLogin(context.TODO(), provider, token)
		if errors.Is(err, service.ErrNotFound) {
			return operations.NewOAuthLoginNotFound()
		}
		if errors.Is(err, service.ErrOAuth) {
			return operations.NewOAuthLoginBadRequest()
		}
		if err != nil {
			return operations.NewOAuthLoginInternalServerError()
		}
		return operations.NewOAuthLoginOK().WithPayload(&operations.OAuthLoginOKBody{
			Token: &tok,
		})
	})
}
