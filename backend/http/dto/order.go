package dto

import "time"

type OrderResponse struct {
	ID        int64     `json:"id" example:"1"`
	UserID    int64     `json:"user_id" example:"1"`
	CreatedAt time.Time `json:"created_at" example:"2026-05-11 01:45:24.864701"`
}

type ListOrderResponse struct {
	Data  []OrderResponse `json:"data"`
	Total int             `json:"total"`
}

type DailyPurchases struct {
	OrderDate time.Time `json:"order_date"`
	Purchases int       `json:"purchases"`
}

type ListDailyPurchasesResponse struct {
	Data  []DailyPurchases `json:"data"`
	Total int              `json:"total"`
}

type CreateOrderRequest struct {
	UserID int64 `json:"user_id" example:"1"`
}
