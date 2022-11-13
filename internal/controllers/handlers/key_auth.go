package handlers

import (
	"context"

	"github.com/go-openapi/errors"

	"github.com/xopoww/wishes/internal/models"
	"github.com/xopoww/wishes/restapi/apimodels"
	"github.com/xopoww/wishes/restapi/apimodels/conv"
)

type (
	OnKeySecurityAuthStartInfo struct{}
	OnKeySecurityAuthDoneInfo  struct {
		Client *models.User
		Err    error
	}
)

func (ac *ApiController) KeySecurityAuth() func(token string) (*apimodels.Principal, error) {
	return func(token string) (*apimodels.Principal, error) {
		onDone := traceOnKeySecurityAuth(ac.t)
		client, err := ac.s.Auth(context.TODO(), token)
		if err != nil {
			onDone(nil, err)
			return nil, errors.New(401, "incorrect api key auth")
		}
		principal := conv.Principal(client)
		onDone(client, nil)
		return principal, nil
	}
}
