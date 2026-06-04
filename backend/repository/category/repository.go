package product

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	categorydomain "shop-api/domain/category"
)

type PgxRepository struct {
	pool *pgxpool.Pool
}

func NewPgxRepository(pool *pgxpool.Pool) *PgxRepository {
	return &PgxRepository{pool: pool}
}

func (r *PgxRepository) Count(ctx context.Context) (int, error) {
	var total int
	err := r.pool.QueryRow(ctx, "SELECT COUNT(*) FROM categories").Scan(&total)
	return total, err
}

func (r *PgxRepository) List(ctx context.Context, limit int, offset int) ([]*categorydomain.Category, error) {
	rows, err := r.pool.Query(ctx, "SELECT id, title FROM categories LIMIT $1 OFFSET $2", limit, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*categorydomain.Category

	for rows.Next() {
		var id int64
		var title string

		if err := rows.Scan(&id, &title); err != nil {
			return nil, err
		}

		products = append(products, categorydomain.NewCategory(id, title))
	}

	return products, rows.Err()
}
