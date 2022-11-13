package conv

import (
	"github.com/xopoww/wishes/restapi/apimodels"
)

func ID(id *apimodels.ID) int64 {
	return *id.ID
}

func SwagID(id int64) *apimodels.ID {
	return &apimodels.ID{ID: &id}
}
