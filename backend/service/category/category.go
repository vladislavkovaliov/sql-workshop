package product

import (
	"context"

	categorydomain "shop-api/domain/category"
)

type Service struct {
	repo categorydomain.Repository
}

func New(repo categorydomain.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) List(ctx context.Context, limit int, offset int) ([]*categorydomain.Category, int, error) {
	categories, err := s.repo.List(ctx, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	return categories, total, nil
}
