package service

import (
	"context"

	"github.com/example/ts-background-jobs-1/internal/domain"
	"github.com/example/ts-background-jobs-1/internal/repository"
)

// ItemService encapsulates item business logic.
type ItemService struct { repo *repository.ItemRepository }

func NewItemService(repo *repository.ItemRepository) *ItemService { return &ItemService{repo: repo} }

func (s *ItemService) List(ctx context.Context, ownerID int64, page, limit int, search string) ([]*domain.Item, int64, error) {
	offset := (page - 1) * limit
	return s.repo.List(ctx, ownerID, offset, limit, search)
}

func (s *ItemService) Create(ctx context.Context, ownerID int64, req domain.ItemCreate) (*domain.Item, error) {
	return s.repo.Create(ctx, ownerID, req)
}

func (s *ItemService) Get(ctx context.Context, id, ownerID int64) (*domain.Item, error) {
	return s.repo.FindByID(ctx, id, ownerID)
}

func (s *ItemService) Update(ctx context.Context, id, ownerID int64, req domain.ItemUpdate) (*domain.Item, error) {
	return s.repo.Update(ctx, id, ownerID, req)
}

func (s *ItemService) Delete(ctx context.Context, id, ownerID int64) error {
	_, err := s.repo.Delete(ctx, id, ownerID)
	return err
}
