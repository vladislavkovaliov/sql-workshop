package user

import "context"

type Repository interface {
	Count(ctx context.Context) (int, error)
	ListCursor(ctx context.Context, cursor int, limit int) ([]*User, error)
	List(ctx context.Context, limit int, offset int) ([]*User, error)
	Search(ctx context.Context, field string, value string) ([]*User, error)
	ListTop3Users(ctx context.Context) ([]*UserWithPurchases, error)
	ListUserByMostExpensiveProduct(ctx context.Context) ([]*UserByMostExpensiveProduct, error)
}
