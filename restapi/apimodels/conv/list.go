package conv

import (
	"github.com/xopoww/wishes/internal/models"
	"github.com/xopoww/wishes/restapi/apimodels"
)

func Access(a *string) models.ListAccess {
	switch *a {
	case apimodels.ListAccessPrivate:
		return models.PrivateAccess
	case apimodels.ListAccessLink:
		return models.LinkAccess
	case apimodels.ListAccessPublic:
		return models.PublicAccess
	default:
		panic("wrong access")
	}
}

func SwagAccess(a models.ListAccess) *string {
	var s string
	switch a {
	case models.PrivateAccess:
		s = apimodels.ListAccessPrivate
	case models.LinkAccess:
		s = apimodels.ListAccessLink
	case models.PublicAccess:
		s = apimodels.ListAccessPublic
	default:
		panic("wrong access")
	}
	return &s
}

func List(l *apimodels.List) *models.List {
	return &models.List{
		Title: *l.Title,
		Access: Access(l.Access),
	}
}

func Item(i *apimodels.ListItem) models.ListItem {
	return models.ListItem{
		Title: *i.Title,
		Desc: i.Desc,
	}
}

func Items(is []*apimodels.ListItem) []models.ListItem {
	iis := make([]models.ListItem, 0, len(is))
	for _, i := range is {
		iis = append(iis, Item(i))
	}
	return iis
}

func SwagList(l *models.List) *apimodels.List {
	return &apimodels.List{
		Title: &l.Title,
		Access: SwagAccess(l.Access),
	}
}

func SwagItem(i models.ListItem) *apimodels.ListItem {
	return &apimodels.ListItem{
		Title: &i.Title,
		Desc: i.Desc,
	}
}

func SwagItems(is []models.ListItem) []*apimodels.ListItem {
	iis := make([]*apimodels.ListItem, 0, len(is))
	for _, i := range is {
		iis = append(iis, SwagItem(i))
	}
	return iis
}