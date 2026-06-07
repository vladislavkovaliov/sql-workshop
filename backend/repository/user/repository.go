package user

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	userdomain "shop-api/domain/user"
)

var allowedFields = map[string]string{
	"email": "email",
	"name":  "name",
}

type PgxRepository struct {
	pool *pgxpool.Pool
}

func NewPgxRepository(pool *pgxpool.Pool) *PgxRepository {
	return &PgxRepository{pool: pool}
}

func (r *PgxRepository) Count(ctx context.Context) (int, error) {
	var total int
	err := r.pool.QueryRow(ctx, "SELECT COUNT(*) FROM users").Scan(&total)
	return total, err
}

func (r *PgxRepository) ListCursor(ctx context.Context, cursor int, limit int) ([]*userdomain.User, error) {
	var query string
	var args []any

	if cursor > 0 {
		query = "SELECT id, name, email FROM users WHERE id > $1 ORDER BY id ASC LIMIT $2"
		args = []any{cursor, limit}
	} else {
		query = "SELECT id, name, email FROM users ORDER BY id ASC LIMIT $1"
		args = []any{limit}
	}

	rows, err := r.pool.Query(ctx, query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*userdomain.User

	for rows.Next() {
		var id int64
		var name string
		var email string

		if err := rows.Scan(&id, &name, &email); err != nil {
			return nil, err
		}

		users = append(users, userdomain.NewUser(id, name, email))
	}

	return users, rows.Err()
}

func (r *PgxRepository) List(ctx context.Context, limit int, offset int) ([]*userdomain.User, error) {
	rows, err := r.pool.Query(ctx, "SELECT id, name, email FROM users LIMIT $1 OFFSET $2", limit, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*userdomain.User

	for rows.Next() {
		var id int64
		var name string
		var email string

		if err := rows.Scan(&id, &name, &email); err != nil {
			return nil, err
		}

		users = append(users, userdomain.NewUser(id, name, email))
	}

	return users, rows.Err()
}

func (r *PgxRepository) Search(ctx context.Context, field string, value string) ([]*userdomain.User, error) {
	col, ok := allowedFields[field]

	if !ok {
		return nil, fmt.Errorf("invalid search field: %s", field)
	}

	query := fmt.Sprintf("SELECT id, name, email FROM users WHERE %s ILIKE '%%' || $1 || '%%'", col)

	rows, err := r.pool.Query(ctx, query, value)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*userdomain.User

	for rows.Next() {
		var id int64
		var name string
		var email string

		if err := rows.Scan(&id, &name, &email); err != nil {
			return nil, err
		}

		users = append(users, userdomain.NewUser(id, name, email))
	}

	return users, rows.Err()
}

func (r *PgxRepository) ListTop3Users(ctx context.Context) ([]*userdomain.UserWithPurchases, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT u.id, u.name, u.email, COUNT(*) as "purchases" FROM users u
		JOIN orders o ON o.user_id = u.id
		GROUP BY u.id, u.email, u.name
		ORDER BY purchases DESC
		LIMIT 3; 
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var top3Users []*userdomain.UserWithPurchases

	for rows.Next() {
		var id int64
		var name string
		var email string
		var purchases int

		if err := rows.Scan(&id, &name, &email, &purchases); err != nil {
			return nil, err
		}

		top3Users = append(top3Users, userdomain.NewUserWithPurchases(id, name, email, purchases))
	}

	return top3Users, rows.Err()
}

func (r *PgxRepository) ListUserByMostExpensiveProduct(ctx context.Context) ([]*userdomain.UserByMostExpensiveProduct, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT DISTINCT u.id, u.email, u.name FROM users u
		JOIN orders o ON o.user_id = u.id
		JOIN order_items oi ON oi.order_id = o.id
		WHERE oi.product_id = (
			SELECT id FROM products
			ORDER BY price DESC
			LIMIT 10
		);
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var userByMostExpensiveProducts []*userdomain.UserByMostExpensiveProduct

	for rows.Next() {
		var id int64
		var name string
		var email string

		if err := rows.Scan(&id, &name, &email); err != nil {
			return nil, err
		}

		userByMostExpensiveProducts = append(
			userByMostExpensiveProducts,
			userdomain.NewUserByMostExpensiveProduct(id, name, email),
		)
	}

	return userByMostExpensiveProducts, rows.Err()

}
