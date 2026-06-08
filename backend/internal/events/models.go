package events

import "time"

type OrderCreated struct {
	OrderID   int64     `json:"order_id"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
