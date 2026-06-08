package order

import (
	"context"
	orderdomain "shop-api/domain/order"
	"shop-api/internal/events"
)

type Service struct {
	repo     orderdomain.Repository
	producer *events.Producer
}

func New(repo orderdomain.Repository, producer *events.Producer) *Service {
	return &Service{repo: repo, producer: producer}
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

func (s *Service) ListDailyPurchases(ctx context.Context) ([]*orderdomain.DailyPurchases, error) {
	dailyPurchases, err := s.repo.ListDailyPurchases(ctx)

	if err != nil {
		return nil, err
	}

	return dailyPurchases, err
}

func (s *Service) Create(ctx context.Context, userID int64) (*orderdomain.Order, error) {
	order, err := s.repo.Create(ctx, userID)

	if err != nil {
		return nil, err
	}

	if err := s.producer.PublishOrderCreated(ctx, events.OrderCreated{
		OrderID:   order.ID(),
		UserID:    order.UserID(),
		CreatedAt: order.CreatedAt(),
	}); err != nil {
		return nil, err
	}

	return order, nil
}
