package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/example/ts-background-jobs-1/internal/domain"
)

// UserRepository handles all user persistence.
type UserRepository struct { db *sql.DB }

func NewUserRepository(db *sql.DB) *UserRepository { return &UserRepository{db: db} }

func (r *UserRepository) Create(ctx context.Context, email, fullName, hashedPassword string) (*domain.User, error) {
	user := &domain.User{}
	err := r.db.QueryRowContext(ctx,
		`INSERT INTO users (email, full_name, hashed_password) VALUES ($1,$2,$3)
		 RETURNING id, email, full_name, is_active, created_at, updated_at`,
		email, fullName, hashedPassword,
	).Scan(&user.ID, &user.Email, &user.FullName, &user.IsActive, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	return user, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, string, error) {
	var user domain.User
	var hash string
	err := r.db.QueryRowContext(ctx,
		`SELECT id, email, full_name, is_active, hashed_password, created_at, updated_at FROM users WHERE email=$1`,
		email,
	).Scan(&user.ID, &user.Email, &user.FullName, &user.IsActive, &hash, &user.CreatedAt, &user.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, "", nil
	}
	return &user, hash, err
}

func (r *UserRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var exists bool
	err := r.db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)", email).Scan(&exists)
	return exists, err
}

var _ = time.Now // avoid unused import
