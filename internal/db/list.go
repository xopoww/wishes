package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mattn/go-sqlite3"
)

var ErrAccessDenied = errors.New("access denied")

type ListItem struct {
	Title string `db:"title"`
	Desc  string `db:"desc"`
}

type List struct {
	ID      int64  `db:"id"`
	OwnerID int64  `db:"owner_id"`
	Title   string `db:"title"`
	Items   []ListItem
}

// GetList retrieves List from database. Access rights are determined
// based on clientId which is an ID of the user requesting the list.
func GetList(id, clientId int64) (*List, error) {
	if db == nil {
		return nil, ErrNotConnected
	}

	tx, err := db.Beginx()
	if err != nil {
		return nil, err
	}

	list := &List{ID: id}
	err = sqlx.Get(tracer(tx), list, `SELECT title, owner_id FROM Lists WHERE id = $1`, list.ID)
	if errors.Is(err, sql.ErrNoRows) {
		err = ErrNotFound
	}
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	//TODO: add access rights for lists

	err = sqlx.Select(tracer(tx), &list.Items, `SELECT title, desc FROM Items WHERE list_id = $1`, list.ID)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	_ = tx.Commit()
	return list, nil
}

func AddList(title string, items []ListItem, owner int64) (id int64, err error) {
	if db == nil {
		return 0, ErrNotConnected
	}

	tx, err := db.Beginx()
	if err != nil {
		return 0, err
	}

	r, err := sqlx.NamedExec(tracer(tx),
		`INSERT INTO Lists (title, owner_id) VALUES (:title, :owner_id)`,
		map[string]interface{}{
			"title":    title,
			"owner_id": owner,
		},
	)
	var serr sqlite3.Error
	if errors.As(err, &serr) && serr.ExtendedCode == sqlite3.ErrConstraintForeignKey {
		err = ErrNotFound
	}
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	id, err = r.LastInsertId()
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	err = insertItems(items, id, tx)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	err = tx.Commit()
	return id, err
}

// EditList saves new list info to database. Target list is determined
// by list.ID. Value of list.OwnerID is not used in any way (including
// access rights check)
func EditList(list *List, clientId int64) error {
	if db == nil {
		return ErrNotConnected
	}

	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	err = sqlx.Get(tracer(db), list, `SELECT owner_id FROM Lists WHERE id = $1`, list.ID)
	if errors.Is(err, sql.ErrNoRows) {
		err = ErrNotFound
	}
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	if list.OwnerID != clientId {
		_ = tx.Rollback()
		return ErrAccessDenied
	}

	_, err = sqlx.NamedExec(tracer(db), `UPDATE Lists SET title = :title WHERE id = :id`, list)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	_, err = sqlx.NamedExec(tracer(db), `DELETE FROM Items WHERE list_id = :id`, list)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	err = insertItems(list.Items, list.ID, tx)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	err = tx.Commit()
	return err
}

func DeleteList(id int64, clientId int64) error {
	if db == nil {
		return ErrNotConnected
	}

	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	var owner int64
	err = sqlx.Get(tracer(tx), &owner, `SELECT owner_id FROM Lists WHERE id = $1`, id)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	if owner != clientId {
		_ = tx.Rollback()
		return ErrAccessDenied
	}

	_, err = sqlx.NamedExec(tracer(tx), `DELETE FROM Lists WHERE id = :id`, map[string]interface{}{"id": id})
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	err = tx.Commit()
	return err
}

// GetUsersLists returns ids of all lists owned by userId and visible by
// clientId.
func GetUserLists(userId, clientId int64) (lids []int64, err error) {
	if db == nil {
		return nil, ErrNotConnected
	}
	//TODO: access rights
	err = sqlx.Select(tracer(db), &lids, `SELECT id FROM Lists WHERE owner_id = $1`, userId)
	return
}

func insertItems(items []ListItem, lid int64, ext sqlx.Ext) error {
	for i, item := range items {
		_, err := sqlx.NamedExec(tracer(ext),
			`INSERT INTO Items (title, desc, list_id) VALUES (:title, :desc, :list_id)`,
			map[string]interface{}{
				"title":   item.Title,
				"desc":    item.Desc,
				"list_id": lid,
			},
		)
		if err != nil {
			return fmt.Errorf("at items[%d]: %w", i, err)
		}
	}
	return nil
}
