package user

import "context"

type Repository interface {
	Count(ctx context.Context) (int, error)
	ListCursor(ctx context.Context, cursor int, limit int) ([]*User, error)
	List(ctx context.Context, limit int, offset int) ([]*User, error)
	SearchByEmail(ctx context.Context, email string) ([]*User, error)
}
