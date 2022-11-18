package conv

import (
	"github.com/go-openapi/strfmt"
)

func ListToken(t *strfmt.Base64) *string {
	if t == nil {
		return nil
	}
	s := string(*t)
	return &s
}
