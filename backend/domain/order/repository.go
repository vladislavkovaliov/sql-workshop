package order

import (
	"context"
	"time"
)

type Repository interface {
	Count(ctx context.Context) (int, error)
	List(ctx context.Context, limit int, offset int) ([]*Order, error)
	ListDailyPurchases(ctx context.Context) ([]*DailyPurchases, error)
	Create(ctx context.Context, userID int64) (*Order, error)
	UpsertDailyPurchase(ctx context.Context, orderDate time.Time) error
}
