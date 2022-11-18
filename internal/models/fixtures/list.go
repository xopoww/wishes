package fixtures

import (
	"fmt"

	"github.com/xopoww/wishes/internal/models"
)

func Item() models.ListItem {
	return models.ListItem{
		Title: "item",
	}
}

func Items(n int) []models.ListItem {
	items := make([]models.ListItem, n)
	for i := range items {
		items[i].Title = fmt.Sprintf("item #%d", i)
	}
	return items
}

func ItemDesc() models.ListItem {
	return models.ListItem{
		Title: "item with desc",
		Desc:  "some description",
	}
}

func ItemsDesc(n int) []models.ListItem {
	items := make([]models.ListItem, n)
	for i := range items {
		items[i].Title = fmt.Sprintf("item with desc #%d", i)
		items[i].Desc = "some description"
	}
	return items
}

func List(items ...models.ListItem) *models.List {
	return &models.List{
		Title: "list",
		Items: items,
	}
}
