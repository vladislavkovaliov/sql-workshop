package product

import (
	"context"

	proddomain "shop-api/domain/product"
)

type Service struct {
	repo proddomain.Repository
}

func New(repo proddomain.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) List(ctx context.Context, limit int, offset int) ([]*proddomain.Product, error) {
	return s.repo.List(ctx, limit, offset)
}

func (s *Service) ListCursor(ctx context.Context, cursor int, limit int) ([]*proddomain.Product, error) {
	return s.repo.ListCursor(ctx, cursor, limit)
}
