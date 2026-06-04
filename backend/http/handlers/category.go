package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"shop-api/http/dto"
	categoryservice "shop-api/service/category"
)

type CategoryHandler struct {
	service *categoryservice.Service
}

func NewCategoryHandler(service *categoryservice.Service) *CategoryHandler {
	return &CategoryHandler{service: service}
}

// ListCategories godoc
//
//	@Summary		List all categories
//	@Description	Returns all categories from the database
//	@Tags			category
//	@Produce		json
//	@Success		200	{object}	dto.ListCategoryResponse
//	@Router			/category [get]
//
//	@Param			limit	query	int	false	"Number of categories to return (default 10)"
//	@Param			offset	query	int	false	"Number of categories to skip (default 0)"
func (h *CategoryHandler) ListCategory(c *gin.Context) {
	defaultLimit := 10
	defaultOffset := 0

	if limit, err := strconv.Atoi(c.Query("limit")); err == nil && limit > 0 {
		defaultLimit = limit
	}

	if offset, err := strconv.Atoi(c.Query("offset")); err == nil && offset >= 0 {
		defaultOffset = offset
	}

	categories, total, err := h.service.List(c.Request.Context(), defaultLimit, defaultOffset)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := make([]dto.CategoryResponse, 0, len(categories))

	for _, p := range categories {
		res = append(res, dto.CategoryResponse{
			ID:    p.ID(),
			Title: p.Title(),
		})
	}

	c.JSON(http.StatusOK, dto.ListCategoryResponse{
		Data:  res,
		Total: total,
	})
}
