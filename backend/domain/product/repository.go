package product

import "context"

type Repository interface {
	Count(ctx context.Context) (int, error)
	ListCursor(ctx context.Context, cursor int, limit int) ([]*Product, error)
	List(ctx context.Context, limit int, offset int) ([]*Product, error)
	Create(ctx context.Context, title string, price float64) (*Product, error)
	TotalRevenue(ctx context.Context) ([]*TotalRevenue, error)
}
