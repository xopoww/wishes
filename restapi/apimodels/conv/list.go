package conv

import (
	"github.com/xopoww/wishes/internal/models"
	"github.com/xopoww/wishes/restapi/apimodels"
)

func List(l *apimodels.List) *models.List {
	return &models.List{
		Title: *l.Title,
		Items: Items(l.Items),
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
		Items: SwagItems(l.Items),
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