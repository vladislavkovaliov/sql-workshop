package product

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	proddomain "shop-api/domain/product"
)

type PgxRepository struct {
	pool *pgxpool.Pool
}

func NewPgxRepository(pool *pgxpool.Pool) *PgxRepository {
	return &PgxRepository{pool: pool}
}

func (r *PgxRepository) ListCursor(ctx context.Context, cursor int, limit int) ([]*proddomain.Product, error) {
	var query string
	var args []any

	if cursor > 0 {
		query = "SELECT id, title, price FROM products WHERE id > $1 ORDER BY id ASC LIMIT $2"
		args = []any{cursor, limit}
	} else {
		query = "SELECT id, title, price FROM products ORDER BY id ASC LIMIT $1"
		args = []any{limit}
	}

	rows, err := r.pool.Query(ctx, query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*proddomain.Product

	for rows.Next() {
		var id int64
		var title string
		var price float64

		if err := rows.Scan(&id, &title, &price); err != nil {
			return nil, err
		}

		products = append(products, proddomain.NewProduct(id, title, price))
	}

	return products, rows.Err()
}

func (r *PgxRepository) List(ctx context.Context, limit int, offset int) ([]*proddomain.Product, error) {
	rows, err := r.pool.Query(ctx, "SELECT id, title, price FROM products LIMIT $1 OFFSET $2", limit, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*proddomain.Product

	for rows.Next() {
		var id int64
		var title string
		var price float64

		if err := rows.Scan(&id, &title, &price); err != nil {
			return nil, err
		}

		products = append(products, proddomain.NewProduct(id, title, price))
	}

	return products, rows.Err()
}
