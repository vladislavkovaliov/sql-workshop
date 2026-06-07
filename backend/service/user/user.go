package user

import (
	"context"
	userdomain "shop-api/domain/user"
)

type Service struct {
	repo userdomain.Repository
}

func New(repo userdomain.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) List(ctx context.Context, limit int, offset int) ([]*userdomain.User, int, error) {
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

func (s *Service) ListCursor(ctx context.Context, cursor int, limit int) ([]*userdomain.User, error) {
	return s.repo.ListCursor(ctx, cursor, limit)
}

func (s *Service) Search(ctx context.Context, field string, value string) ([]*userdomain.User, error) {
	return s.repo.Search(ctx, field, value)
}

func (s *Service) ListTop3Users(ctx context.Context) ([]*userdomain.UserWithPurchases, error) {
	return s.repo.ListTop3Users(ctx)
}

func (s *Service) ListUserByMostExpensiveProduct(ctx context.Context) ([]*userdomain.UserByMostExpensiveProduct, error) {
	return s.repo.ListUserByMostExpensiveProduct(ctx)
}
