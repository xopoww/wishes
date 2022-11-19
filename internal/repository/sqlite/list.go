package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/mattn/go-sqlite3"
	"github.com/xopoww/wishes/internal/models"
	"github.com/xopoww/wishes/internal/service"
)

func (r *handle) GetUserLists(ctx context.Context, id int64, publicOnly bool) (lids []int64, err error) {
	query := `SELECT Lists.id FROM Lists JOIN ListAccessEnum ON Lists.access = ListAccessEnum.N WHERE Lists.owner_id = $1`
	if publicOnly {
		query += ` AND ListAccessEnum.S = 'public'`
	}
	err = sqlx.SelectContext(ctx, r.tracer(), &lids, query, id)
	return
}

func (r *handle) GetList(ctx context.Context, id int64) (*models.List, error) {
	list := &models.List{ID: id}
	row := r.tracer().QueryRowxContext(ctx, `SELECT title, owner_id, access, revision FROM Lists WHERE id = $1`, id)
	err := row.Scan(&list.Title, &list.OwnerID, &list.Access, &list.RevisionID)
	if errors.Is(err, sql.ErrNoRows) {
		err = fmt.Errorf("list_id %d: %w", id, service.ErrNotFound)
	}
	if err != nil {
		return nil, fmt.Errorf("select list: %w", err)
	}

	return list, nil
}

func (r *handle) GetListItems(ctx context.Context, list *models.List) ([]models.ListItem, error) {
	rows, err := r.tracer().QueryxContext(ctx, `SELECT id, title, desc FROM Items WHERE list_id = $1`, list.ID)
	if err != nil {
		return nil, fmt.Errorf("select items: %w", err)
	}
	items := make([]models.ListItem, 0)
	for rows.Next() {
		item := models.ListItem{}
		err = rows.Scan(&item.ID, &item.Title, &item.Desc)
		if err != nil {
			return nil, fmt.Errorf("scan item: %w", err)
		}
		items = append(items, item)
	}
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("scan rows: %w", err)
	}
	return items, nil
}

func (r *handle) AddList(ctx context.Context, list *models.List) (*models.List, error) {
	res, err := r.tracer().ExecContext(ctx,
		`INSERT INTO Lists (title, owner_id, access, revision) VALUES ($1, $2, $3, $4)`,
		list.Title, list.OwnerID, list.Access, list.RevisionID,
	)
	var serr sqlite3.Error
	if errors.As(err, &serr) && serr.ExtendedCode == sqlite3.ErrConstraintForeignKey {
		err = fmt.Errorf("user_id %d: %w", list.OwnerID, service.ErrNotFound)
	}
	if err != nil {
		return nil, fmt.Errorf("insert list: %w", err)
	}
	list.ID, err = res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("last insert id: %w", err)
	}

	list.Items, err = r.insertItems(ctx, list.Items, list.ID)
	if err != nil {
		return nil, fmt.Errorf("insert items: %w", err)
	}

	return list, nil
}

func (r *handle) EditList(ctx context.Context, list *models.List) (*models.List, error) {
	res, err := r.tracer().ExecContext(ctx,
		`UPDATE Lists SET title = $1, access = $2, revision = $3 WHERE id = $4`,
		list.Title, list.Access, list.RevisionID, list.ID,
	)
	if err != nil {
		return nil, fmt.Errorf("update list: %w", err)
	}
	ra, err := res.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("rows affected: %w", err)
	}
	if ra == 0 {
		return nil, fmt.Errorf("list_id %d: %w", list.ID, service.ErrNotFound)
	}
	return list, nil
}

func (r *handle) AddListItems(ctx context.Context, list *models.List, items []models.ListItem) ([]models.ListItem, error) {
	var err error
	items, err = r.insertItems(ctx, items, list.ID)
	var serr sqlite3.Error
	if errors.As(err, &serr) && serr.ExtendedCode == sqlite3.ErrConstraintForeignKey {
		err = fmt.Errorf("list_id %d: %w", list.ID, service.ErrNotFound)
	}
	return items, err
}

func (r *handle) DeleteListItems(ctx context.Context, list *models.List, ids []int64) error {
	queryBuilder := &strings.Builder{}
	queryBuilder.WriteString(`DELETE FROM Items WHERE list_id == $1 AND id IN (`)
	args := make([]interface{}, 1, len(ids) + 1)
	args[0] = list.ID
	for i := range ids {
		fmt.Fprintf(queryBuilder, "$%d", i+2)
		if i != len(ids)-1 {
			queryBuilder.WriteRune(',')
		}
		args = append(args, ids[i])
	}
	queryBuilder.WriteRune(')')
	res, err := r.tracer().ExecContext(ctx, queryBuilder.String(), args...)
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}
	ra, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %w", err)
	}
	if int(ra) != len(ids) {
		return fmt.Errorf("%w (deleted only %d rows)", service.ErrNotFound, ra)
	}
	return nil
}

func (r *handle) DeleteList(ctx context.Context, list *models.List) error {
	res, err := r.tracer().ExecContext(ctx, `DELETE FROM Lists WHERE id = $1`, list.ID)
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

func (r *handle) insertItems(ctx context.Context, items []models.ListItem, lid int64) ([]models.ListItem, error) {
	for i, item := range items {
		res, err := r.tracer().ExecContext(ctx,
			`INSERT INTO Items (title, desc, list_id) VALUES ($1, $2, $3)`,
			item.Title, item.Desc, lid,
		)
		if err != nil {
			return nil, fmt.Errorf("at items[%d]: %w", i, err)
		}
		liid, err := res.LastInsertId()
		if err != nil {
			return nil, fmt.Errorf("at items[%d] (get liid): %w", i, err)
		}
		items[i].ID = liid
	}
	return items, nil
}
