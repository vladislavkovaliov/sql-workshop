package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"shop-api/http/dto"
	userservice "shop-api/service/user"
)

type UserHandler struct {
	service *userservice.Service
}

func NewUserHandler(service *userservice.Service) *UserHandler {
	return &UserHandler{service: service}
}

// ListUsers godoc
//
//	@Summary		List all users
//	@Description	Returns all users from the database
//	@Tags			users
//	@Produce		json
//	@Success		200	{object}	dto.ListUserResponse
//	@Router			/users [get]
//
//	@Param			limit	query	int	false	"Number of users to return (default 10)"
//	@Param			offset	query	int	false	"Number of users to skip (default 0)"
func (h *UserHandler) ListUsers(c *gin.Context) {
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

	users, total, err := h.service.List(ctx, defaultLimit, defaultOffset)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	res := make([]dto.UserResponse, 0, len(users))

	for _, u := range users {
		res = append(res, dto.UserResponse{
			ID:    u.ID(),
			Name:  u.Name(),
			Email: u.Email(),
		})
	}

	c.JSON(http.StatusOK, dto.ListUserResponse{
		Data:  res,
		Total: total,
	})

}

// ListCursorUsers godoc
//
//	@Summary		List users (cursor-based)
//	@Description	Returns users with cursor-based pagination. Pass the last user ID from the previous response as cursor.
//	@Tags			users
//	@Produce		json
//	@Success		200	{object}	dto.CursorUserResponse
//	@Router			/users/cursor [get]
//
//	@Param			limit	query	int	false	"Number of users to return (default 10)"
//	@Param			cursor	query	int	false	"Last users ID from previous page"
func (h *UserHandler) ListCursorUsers(c *gin.Context) {
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

	users, err := h.service.ListCursor(ctx, defaultCursor, defaultLimit)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	res := make([]dto.UserResponse, 0, len(users))

	for _, u := range users {
		res = append(res, dto.UserResponse{
			ID:    u.ID(),
			Name:  u.Name(),
			Email: u.Email(),
		})
	}

	var nextCursor int64
	if len(res) > 0 {
		nextCursor = res[len(res)-1].ID
	}

	c.JSON(http.StatusOK, dto.CursorUserResponse{
		Users:      res,
		NextCursor: nextCursor,
	})
}

// Search godoc
//
//	@Summary		Search user by field
//	@Description	Search users by field (partial match)
//	@Tags			users
//	@Produce		json
//	@Success		200	{object}	dto.ListUserResponse
//	@Router			/users/search [get]
//
// @Param field query string true "Field to search by" Enums(email, name)
// @Param value	query string true "Value to search for"
func (h *UserHandler) Search(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")

	if field == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "field is required"})
		return
	}

	if value == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "value is required"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	users, err := h.service.Search(ctx, field, value)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	res := make([]dto.UserResponse, 0, len(users))

	for _, u := range users {
		res = append(res, dto.UserResponse{
			ID:    u.ID(),
			Name:  u.Name(),
			Email: u.Email(),
		})
	}

	c.JSON(http.StatusOK, dto.ListUserResponse{
		Data:  res,
		Total: len(res),
	})
}

// ListTop3Users godoc
//
//	@Summary		Get top 3 users by purchases
//	@Description	Returns top 3 users with the highest number of purchases
//	@Tags			users
//	@Produce		json
//	@Success		200	{object}	dto.ListUserWithPurchasesResponse
//	@Router			/users/top-3-users [get]
func (h *UserHandler) ListTop3Users(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)

	defer cancel()

	top3Users, err := h.service.ListTop3Users(ctx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	res := make([]dto.UserWithPurchases, 0, len(top3Users))

	for _, u := range top3Users {
		res = append(res, dto.UserWithPurchases{
			ID:        u.ID(),
			Name:      u.Name(),
			Email:     u.Email(),
			Purchases: u.Purchases(),
		})
	}

	c.JSON(http.StatusOK, dto.ListUserWithPurchasesResponse{
		Data:  res,
		Total: len(res),
	})
}

// ListUserByMostExpensiveProduct godoc
//
//	@Summary		Get users by the most expensive product
//	@Description	Returns users by the most expensive product
//	@Tags			users
//	@Produce		json
//	@Success		200	{object}	dto.ListUserByMostExpensiveProductResponse
//	@Router			/users/by-most-expensive-product [get]
func (h *UserHandler) ListUserByMostExpensiveProduct(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)

	defer cancel()

	users, err := h.service.ListUserByMostExpensiveProduct(ctx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	res := make([]dto.UserByMostExpensiveProduct, 0, len(users))

	for _, u := range users {
		res = append(res, dto.UserByMostExpensiveProduct{
			ID:    u.ID(),
			Name:  u.Name(),
			Email: u.Email(),
		})
	}

	c.JSON(http.StatusOK, dto.ListUserByMostExpensiveProductResponse{
		Data:  res,
		Total: len(res),
	})
}
