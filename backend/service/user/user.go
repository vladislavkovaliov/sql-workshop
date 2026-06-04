package user

import (
	"context"
	"fmt"
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

func (s *Service) SearchByEmail(ctx context.Context, email string) ([]*userdomain.User, error) {
	fmt.Println(email)

	return s.repo.SearchByEmail(ctx, email)
}
