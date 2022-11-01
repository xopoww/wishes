package db

// import (
// 	"database/sql"
// 	"errors"
// )

// type ListItem struct {
// 	ID	  int
// 	Title string
// 	Desc  string
// }

// type List struct {
// 	ID 	  int
// 	Owner int //TODO: switch user logic to id
// 	Items []ListItem
// }

// func GetList(id int) (*List, error) {
// 	if db == nil {
// 		return nil, ErrNotConnected
// 	}

// 	list := &List{}
// 	row := db.QueryRow(`SELECT id FROM Lists WHERE id = $1`, id)
// 	err := row.Scan(&list.ID)
// 	if errors.Is(err, sql.ErrNoRows) {
// 		return nil, ErrNotFound
// 	}
// 	if err != nil {
// 		return nil, err
// 	}

// 	rows := db.Query(`SELECT id, `)
// }
