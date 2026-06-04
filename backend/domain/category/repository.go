package category

import "context"

type Repository interface {
	Count(ctx context.Context) (int, error)
	List(ctx context.Context, limit int, offset int) ([]*Category, error)
}
