package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/example/ts-background-jobs-1/internal/domain"
)

// ItemRepository handles all item persistence.
type ItemRepository struct { db *sql.DB }

func NewItemRepository(db *sql.DB) *ItemRepository { return &ItemRepository{db: db} }

func (r *ItemRepository) List(ctx context.Context, ownerID int64, offset, limit int, search string) ([]*domain.Item, int64, error) {
	args := []any{ownerID}
	where := "owner_id=$1"
	if search != "" {
		where += fmt.Sprintf(" AND (title ILIKE $%d OR description ILIKE $%d)", len(args)+1, len(args)+2)
		term := "%" + search + "%"
		args = append(args, term, term)
	}
	var total int64
	_ = r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM items WHERE "+where, args...).Scan(&total)
	args = append(args, limit, offset)
	rows, err := r.db.QueryContext(ctx,
		fmt.Sprintf("SELECT id,title,description,status,owner_id,created_at,updated_at FROM items WHERE %s ORDER BY created_at DESC LIMIT $%d OFFSET $%d", where, len(args)-1, len(args)),
		args...)
	if err != nil { return nil, 0, err }
	defer rows.Close()
	var items []*domain.Item
	for rows.Next() {
		var item domain.Item
		if err := rows.Scan(&item.ID, &item.Title, &item.Description, &item.Status, &item.OwnerID, &item.CreatedAt, &item.UpdatedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, &item)
	}
	return items, total, rows.Err()
}

func (r *ItemRepository) Create(ctx context.Context, ownerID int64, req domain.ItemCreate) (*domain.Item, error) {
	status := req.Status
	if status == "" { status = domain.ItemStatusActive }
	item := &domain.Item{}
	err := r.db.QueryRowContext(ctx,
		`INSERT INTO items (title,description,status,owner_id) VALUES ($1,$2,$3,$4)
		 RETURNING id,title,description,status,owner_id,created_at,updated_at`,
		req.Title, req.Description, status, ownerID,
	).Scan(&item.ID, &item.Title, &item.Description, &item.Status, &item.OwnerID, &item.CreatedAt, &item.UpdatedAt)
	return item, err
}

func (r *ItemRepository) FindByID(ctx context.Context, id, ownerID int64) (*domain.Item, error) {
	item := &domain.Item{}
	err := r.db.QueryRowContext(ctx,
		`SELECT id,title,description,status,owner_id,created_at,updated_at FROM items WHERE id=$1 AND owner_id=$2`,
		id, ownerID,
	).Scan(&item.ID, &item.Title, &item.Description, &item.Status, &item.OwnerID, &item.CreatedAt, &item.UpdatedAt)
	if err == sql.ErrNoRows { return nil, nil }
	return item, err
}

func (r *ItemRepository) Update(ctx context.Context, id, ownerID int64, req domain.ItemUpdate) (*domain.Item, error) {
	var sets []string
	var args []any
	if req.Title != nil {
		sets = append(sets, fmt.Sprintf("title=$%d", len(args)+1))
		args = append(args, *req.Title)
	}
	if req.Description != nil {
		sets = append(sets, fmt.Sprintf("description=$%d", len(args)+1))
		args = append(args, *req.Description)
	}
	if req.Status != nil {
		sets = append(sets, fmt.Sprintf("status=$%d", len(args)+1))
		args = append(args, *req.Status)
	}
	if len(sets) == 0 { return r.FindByID(ctx, id, ownerID) }
	sets = append(sets, "updated_at=NOW()")
	args = append(args, id, ownerID)
	item := &domain.Item{}
	err := r.db.QueryRowContext(ctx,
		fmt.Sprintf("UPDATE items SET %s WHERE id=$%d AND owner_id=$%d RETURNING id,title,description,status,owner_id,created_at,updated_at",
			strings.Join(sets, ","), len(args)-1, len(args)),
		args...,
	).Scan(&item.ID, &item.Title, &item.Description, &item.Status, &item.OwnerID, &item.CreatedAt, &item.UpdatedAt)
	if err == sql.ErrNoRows { return nil, nil }
	return item, err
}

func (r *ItemRepository) Delete(ctx context.Context, id, ownerID int64) (bool, error) {
	res, err := r.db.ExecContext(ctx, "DELETE FROM items WHERE id=$1 AND owner_id=$2", id, ownerID)
	if err != nil { return false, err }
	n, _ := res.RowsAffected()
	return n > 0, nil
}
