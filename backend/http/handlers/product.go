package handlers

import (
	"net/http"
	"strconv"

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

	products, total, err := h.service.List(c.Request.Context(), defaultLimit, defaultOffset)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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

	products, err := h.service.ListCursor(c.Request.Context(), defaultCursor, defaultLimit)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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

	product, err := h.service.Create(c.Request.Context(), req.Title, req.Price)
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
