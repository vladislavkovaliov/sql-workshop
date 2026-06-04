package order

import (
	"context"
	orderdomain "shop-api/domain/order"
)

type Service struct {
	repo orderdomain.Repository
}

func New(repo orderdomain.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) List(ctx context.Context, limit int, offset int) ([]*orderdomain.Order, int, error) {
	orders, err := s.repo.List(ctx, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.Count(ctx)

	if err != nil {
		return nil, 0, err
	}

	return orders, total, nil
}
