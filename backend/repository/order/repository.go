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

func (r *PgxRepository) ListDailyPurchases(ctx context.Context) ([]*orderdomain.DailyPurchases, error) {
	// rows, err := r.pool.Query(ctx, `
	// SELECT
	// 	created_at::date AS order_date,
	// 	COUNT(*)         AS purchases
	// FROM orders
	// GROUP BY order_date
	// ORDER BY order_date
	// `)

	rows, err := r.pool.Query(ctx, `
		SELECT order_date, purchases FROM daily_purchases ORDER BY order_date
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var dailyPurchases []*orderdomain.DailyPurchases

	for rows.Next() {
		var order_date time.Time
		var purchases int

		if err := rows.Scan(&order_date, &purchases); err != nil {
			return nil, err
		}

		dailyPurchases = append(dailyPurchases, orderdomain.NewDailyPurchases(order_date, purchases))
	}

	return dailyPurchases, rows.Err()
}

func (r *PgxRepository) Create(ctx context.Context, userID int64) (*orderdomain.Order, error) {
	var id int64
	var createdAt time.Time

	err := r.pool.QueryRow(ctx, `
		INSERT INTO orders (user_id) VALUES ($1) RETURNING id, created_at
	`, userID).Scan(&id, &createdAt)

	if err != nil {
		return nil, err
	}

	return orderdomain.NewOrder(id, userID, createdAt), nil
}

func (r *PgxRepository) UpsertDailyPurchase(ctx context.Context, orderDate time.Time) error {
	_, err := r.pool.Exec(ctx, `
		INSERT INTO daily_purchases (order_date, purchases)
		VALUES ($1, 1)
		ON CONFLICT (order_date) DO UPDATE
		SET purchases = daily_purchases.purchases + 1
	`, orderDate)

	return err
}
