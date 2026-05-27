package product

import "context"

type Repository interface {
	ListCursor(ctx context.Context, cursor int, limit int) ([]*Product, error)
	List(ctx context.Context, limit int, offset int) ([]*Product, error)
}
