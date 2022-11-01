package handlers

import (
	"github.com/go-openapi/errors"
	"github.com/xopoww/wishes/internal/auth"
	"github.com/xopoww/wishes/internal/models"
)

type (
	OnKeySecurityAuthStartInfo struct{}
	OnKeySecurityAuthDoneInfo  struct {
		Principal *models.Principal
		Err       error
	}
)

func KeySecurityAuth(t Trace) func(token string) (*models.Principal, error) {
	return func(token string) (*models.Principal, error) {
		onDone := traceOnKeySecurityAuth(t)
		principal, err := auth.ValidateToken(token)
		if err != nil {
			onDone(nil, err)
			return nil, errors.New(401, "incorrect api key auth")
		}
		onDone(principal, nil)
		return principal, nil
	}
}
