package order

import (
	"context"
	"time"

	orderdomain "shop-api/domain/order"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PgxRepository struct {
	pool *pgxpool.Pool
}

func NewPgxRepository(pool *pgxpool.Pool) *PgxRepository {
	return &PgxRepository{pool: pool}
}

func (r *PgxRepository) Count(ctx context.Context) (int, error) {
	var total int
	err := r.pool.QueryRow(ctx, "SELECT COUNT(*) FROM orders").Scan(&total)
	return total, err
}

func (r *PgxRepository) List(ctx context.Context, limit int, offset int) ([]*orderdomain.Order, error) {
	rows, err := r.pool.Query(ctx, "SELECT id, user_id, created_at FROM orders LIMIT $1 OFFSET $2", limit, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var orders []*orderdomain.Order

	for rows.Next() {
		var id int64
		var userID int64
		var createdAt time.Time

		if err := rows.Scan(&id, &userID, &createdAt); err != nil {
			return nil, err
		}

		orders = append(orders, orderdomain.NewOrder(id, userID, createdAt))
	}

	return orders, rows.Err()
}
