package models

type ListItem struct {
	Title string
	Desc  string
}

type List struct {
	ID      int64
	OwnerID int64
	Title   string
	Items   []ListItem
}
