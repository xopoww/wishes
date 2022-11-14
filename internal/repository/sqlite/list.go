package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mattn/go-sqlite3"
	"github.com/xopoww/wishes/internal/models"
	"github.com/xopoww/wishes/internal/service"
)

func (r *repository) GetUserLists(ctx context.Context, id int64) (lids []int64, err error) {
	err = sqlx.SelectContext(ctx, r.tracer(r.db), &lids, `SELECT id FROM Lists WHERE owner_id = $1`, id)
	return
}

func (r *repository) GetList(ctx context.Context, id int64) (*models.List, error) {
	list := &models.List{ID: id}
	row := r.tracer(r.db).QueryRowxContext(ctx, `SELECT title, owner_id FROM Lists WHERE id = $1`, id)
	err := row.Scan(&list.Title, &list.OwnerID)
	if errors.Is(err, sql.ErrNoRows) {
		err = fmt.Errorf("list_id %d: %w", id, service.ErrNotFound)
	}
	if err != nil {
		return nil, fmt.Errorf("select list: %w", err)
	}

	return list, nil
}

func (r *repository) GetListItems(ctx context.Context, list *models.List) (*models.List, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, fmt.Errorf("begin: %w", err)
	}

	rows, err := r.tracer(tx).QueryxContext(ctx, `SELECT title, desc FROM Items WHERE list_id = $1`, list.ID)
	if err != nil {
		_ = tx.Rollback()
		return nil, fmt.Errorf("select items: %w", err)
	}
	for rows.Next() {
		item := models.ListItem{}
		err = rows.Scan(&item.Title, &item.Desc)
		if err != nil {
			_ = tx.Rollback()
			return nil, fmt.Errorf("scan item: %w", err)
		}
		list.Items = append(list.Items, item)
	}
	err = rows.Err()
	if err != nil {
		_ = tx.Rollback()
		return nil, fmt.Errorf("scan rows: %w", err)
	}

	_ = tx.Commit()
	return list, nil
}

func (r *repository) AddList(ctx context.Context, list *models.List) (*models.List, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, fmt.Errorf("begin: %w", err)
	}

	res, err := r.tracer(tx).ExecContext(ctx, `INSERT INTO Lists (title, owner_id) VALUES ($1, $2)`, list.Title, list.OwnerID)
	var serr sqlite3.Error
	if errors.As(err, &serr) && serr.ExtendedCode == sqlite3.ErrConstraintForeignKey {
		err = fmt.Errorf("user_id %d: %w", list.OwnerID, service.ErrNotFound)
	}
	if err != nil {
		_ = tx.Rollback()
		return nil, fmt.Errorf("insert list: %w", err)
	}
	list.ID, err = res.LastInsertId()
	if err != nil {
		_ = tx.Rollback()
		return nil, fmt.Errorf("last insert id: %w", err)
	}

	err = r.insertItems(ctx, list.Items, list.ID, tx)
	if err != nil {
		_ = tx.Rollback()
		return nil, fmt.Errorf("insert items: %w", err)
	}

	err = tx.Commit()
	return list, err
}

func (r *repository) EditList(ctx context.Context, list *models.List) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return fmt.Errorf("begin: %w", err)
	}

	res, err := r.tracer(tx).ExecContext(ctx, `UPDATE Lists SET title = $1 WHERE id = $2`, list.Title, list.ID)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("update list: %w", err)
	}
	ra, err := res.RowsAffected()
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("rows affected: %w", err)
	}
	if ra == 0 {
		_ = tx.Rollback()
		return fmt.Errorf("list_id %d: %w", list.ID, service.ErrNotFound)
	}

	_, err = r.tracer(tx).ExecContext(ctx, `DELETE FROM Items WHERE list_id = $1`, list.ID)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("delete items: %w", err)
	}

	err = r.insertItems(ctx, list.Items, list.ID, tx)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("insert items: %w", err)
	}

	err = tx.Commit()
	return err
}

func (r *repository) DeleteList(ctx context.Context, list *models.List) error {
	res, err := r.tracer(r.db).ExecContext(ctx, `DELETE FROM Lists WHERE id = $1`, list.ID)
	if err != nil {
		return fmt.Errorf("delete list: %w", err)
	}
	ra, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %w", err)
	}
	if ra == 0 {
		return fmt.Errorf("list_id %d: %w", list.ID, service.ErrNotFound)
	}
	return nil
}

func (r *repository) insertItems(ctx context.Context, items []models.ListItem, lid int64, ext sqlx.ExtContext) error {
	for i, item := range items {
		_, err := r.tracer(ext).ExecContext(ctx,
			`INSERT INTO Items (title, desc, list_id) VALUES ($1, $2, $3)`,
			item.Title, item.Desc, lid,
		)
		if err != nil {
			return fmt.Errorf("at items[%d]: %w", i, err)
		}
	}
	return nil
}
