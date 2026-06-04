package product

import (
	"context"
	"errors"

	proddomain "shop-api/domain/product"
)

type Service struct {
	repo proddomain.Repository
}

func New(repo proddomain.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) List(ctx context.Context, limit int, offset int) ([]*proddomain.Product, int, error) {
	products, err := s.repo.List(ctx, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.Count(ctx)

	if err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

func (s *Service) ListCursor(ctx context.Context, cursor int, limit int) ([]*proddomain.Product, error) {
	return s.repo.ListCursor(ctx, cursor, limit)
}

func (s *Service) Create(ctx context.Context, title string, price float64) (*proddomain.Product, error) {
	if title == "" {
		return nil, errors.New("title is required")
	}

	if price <= 0 {
		return nil, errors.New("price must be positive")
	}

	return s.repo.Create(ctx, title, price)
}

func (s *Service) TotalRevenue(ctx context.Context) ([]*proddomain.TotalRevenue, error) {
	return s.repo.TotalRevenue(ctx)
}
