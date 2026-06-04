package handlers

import (
	"net/http"
	"strconv"

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

	orders, total, err := h.service.List(c.Request.Context(), defaultLimit, defaultOffset)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := make([]dto.OrderResponse, 0, len(orders))

	for _, p := range orders {
		res = append(res, dto.OrderResponse{
			ID:        p.ID(),
			UserId:    p.UserID(),
			CreatedAt: p.CreatedAt(),
		})
	}

	c.JSON(http.StatusOK, dto.ListOrderResponse{
		Data:  res,
		Total: total,
	})
}
