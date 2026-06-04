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

func (r *PgxRepository) Count(ctx context.Context) (int, error) {
	var total int
	err := r.pool.QueryRow(ctx, "SELECT COUNT(*) FROM products").Scan(&total)
	return total, err
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

func (r *PgxRepository) Create(ctx context.Context, title string, price float64) (*proddomain.Product, error) {
	var id int64

	err := r.pool.QueryRow(ctx, "INSERT INTO products (title, price) VALUES ($1, $2) RETURNING id", title, price).Scan(&id)

	if err != nil {
		return nil, err
	}

	return proddomain.NewProduct(id, title, price), nil
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

func (r *PgxRepository) TotalRevenue(ctx context.Context) ([]*proddomain.TotalRevenue, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT p.title, SUM(p.price * oi.quantity) AS revenue
		FROM order_items oi
		JOIN products p ON p.id = oi.product_id
		GROUP BY p.id, p.title
		ORDER BY revenue DESC;
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var totalRevenues []*proddomain.TotalRevenue

	for rows.Next() {
		var title string
		var revenue float64

		if err := rows.Scan(&title, &revenue); err != nil {
			return nil, err
		}

		totalRevenues = append(totalRevenues, proddomain.NewTotalRevenue(title, revenue))
	}

	return totalRevenues, rows.Err()
}
