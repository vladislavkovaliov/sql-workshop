package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"shop-api/http/dto"
	productservice "shop-api/service/product"
)

type ProductHandler struct {
	service *productservice.Service
}

func NewProductHandler(service *productservice.Service) *ProductHandler {
	return &ProductHandler{service: service}
}

// ListProducts godoc
//
//	@Summary		List all products
//	@Description	Returns all products from the database
//	@Tags			products
//	@Produce		json
//	@Success		200	{object}	dto.ListProductResponse
//	@Router			/products [get]
//
//	@Param			limit	query	int	false	"Number of products to return (default 10)"
//	@Param			offset	query	int	false	"Number of products to skip (default 0)"
func (h *ProductHandler) ListProducts(c *gin.Context) {
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

	products, total, err := h.service.List(ctx, defaultLimit, defaultOffset)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	res := make([]dto.ProductResponse, 0, len(products))

	for _, p := range products {
		res = append(res, dto.ProductResponse{
			ID:    p.ID(),
			Title: p.Title(),
			Price: p.Price(),
		})
	}

	c.JSON(http.StatusOK, dto.ListProductResponse{
		Data:  res,
		Total: total,
	})

}

// ListCursorProducts godoc
//
//	@Summary		List products (cursor-based)
//	@Description	Returns products with cursor-based pagination. Pass the last product ID from the previous response as cursor.
//	@Tags			products
//	@Produce		json
//	@Success		200	{object}	dto.CursorProductsResponse
//	@Router			/products/cursor [get]
//
//	@Param			limit	query	int	false	"Number of products to return (default 10)"
//	@Param			cursor	query	int	false	"Last product ID from previous page"
func (h *ProductHandler) ListCursorProducts(c *gin.Context) {
	defaultLimit := 10
	defaultCursor := 0

	if limit, err := strconv.Atoi(c.Query("limit")); err == nil && limit > 0 {
		defaultLimit = limit
	}

	if cursor, err := strconv.Atoi(c.Query("cursor")); err == nil && cursor >= 0 {
		defaultCursor = cursor
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	products, err := h.service.ListCursor(ctx, defaultCursor, defaultLimit)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	res := make([]dto.ProductResponse, 0, len(products))

	for _, p := range products {
		res = append(res, dto.ProductResponse{
			ID:    p.ID(),
			Title: p.Title(),
			Price: p.Price(),
		})
	}

	var nextCursor int64
	if len(res) > 0 {
		nextCursor = res[len(res)-1].ID
	}

	c.JSON(http.StatusOK, dto.CursorProductsResponse{
		Products:   res,
		NextCursor: nextCursor,
	})
}

// CreateProduct godoc
//
//	@Summary		Create a product
//	@Description	Add a new product to the database
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			body	body	dto.CreateProductRequest	true	"Product data"
//	@Success		201		{object}	dto.ProductResponse
//	@Failure		400		{object}	map[string]string
//	@Router			/products [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req dto.CreateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	product, err := h.service.Create(ctx, req.Title, req.Price)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.ProductResponse{
		ID:    product.ID(),
		Title: product.Title(),
		Price: product.Price(),
	})
}

// TotalRevenue godoc
//
//	@Summary		Get total revenue per product
//	@Description	Returns products sorted by total revenue (SUM of price * quantity per product)
//	@Tags			products
//	@Produce		json
//	@Success		200	{object}	dto.ListTotalRevenueResponse
//	@Router			/products/revenue [get]
func (h *ProductHandler) TotalRevenue(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	totalRevenues, err := h.service.TotalRevenue(ctx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	res := make([]dto.TotalRevenueResponse, 0, len(totalRevenues))

	for _, p := range totalRevenues {
		res = append(res, dto.TotalRevenueResponse{
			Title:   p.Title(),
			Revenue: p.Revenue(),
		})
	}

	c.JSON(http.StatusOK, dto.ListTotalRevenueResponse{
		Data:  res,
		Total: len(totalRevenues),
	})
}
