package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"shop-api/http/dto"
	orderservice "shop-api/service/order"
)

type OrderHandler struct {
	service *orderservice.Service
}

func NewOrderHandler(service *orderservice.Service) *OrderHandler {
	return &OrderHandler{service: service}
}

// ListOrders godoc
//
//	@Summary		List all orders
//	@Description	Returns all orders from the database
//	@Tags			orders
//	@Produce		json
//	@Success		200	{object}	dto.ListOrderResponse
//	@Router			/orders [get]
//
//	@Param			limit	query	int	false	"Number of products to return (default 10)"
//	@Param			offset	query	int	false	"Number of products to skip (default 0)"
func (h *OrderHandler) ListOrder(c *gin.Context) {
	defaultLimit := 10
	defaultOffset := 0

	if limit, err := strconv.Atoi(c.Query("limit")); err == nil && limit > 0 {
		defaultLimit = limit
	}

	if offset, err := strconv.Atoi(c.Query("offset")); err == nil && offset >= 0 {
		defaultOffset = offset
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	orders, total, err := h.service.List(ctx, defaultLimit, defaultOffset)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	res := make([]dto.OrderResponse, 0, len(orders))

	for _, p := range orders {
		res = append(res, dto.OrderResponse{
			ID:        p.ID(),
			UserID:    p.UserID(),
			CreatedAt: p.CreatedAt(),
		})
	}

	c.JSON(http.StatusOK, dto.ListOrderResponse{
		Data:  res,
		Total: total,
	})
}

// ListDailyPurchases godoc
//
//	@Summary		List all daily purchases
//	@Description	Returns all orders from the database
//	@Tags			orders
//	@Produce		json
//	@Success		200	{object}	dto.ListDailyPurchasesResponse
//	@Router			/orders/daily-purchases [get]
func (h *OrderHandler) ListDailyPurchases(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)

	defer cancel()

	dailyPurchases, err := h.service.ListDailyPurchases(ctx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	res := make([]dto.DailyPurchases, 0, len(dailyPurchases))

	for _, p := range dailyPurchases {
		res = append(res, dto.DailyPurchases{
			OrderDate: p.OrderDate(),
			Purchases: p.Purchases(),
		})
	}

	c.JSON(http.StatusOK, dto.ListDailyPurchasesResponse{
		Data:  res,
		Total: len(dailyPurchases),
	})
}

// CreateOrder godoc
//
//	@Summary		Create a new order
//	@Description	Creates an order and publishes order.created event to Kafka
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			body	body	dto.CreateOrderRequest	true	"Order data"
//	@Success		201		{object}	dto.OrderResponse
//	@Failure		400		{object}	map[string]string
//	@Router			/orders [post]
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req dto.CreateOrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)

	defer cancel()

	order, err := h.service.Create(ctx, req.UserID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})

		return
	}

	c.JSON(http.StatusCreated, dto.OrderResponse{
		ID:        order.ID(),
		UserID:    order.UserID(),
		CreatedAt: order.CreatedAt(),
	})
}
