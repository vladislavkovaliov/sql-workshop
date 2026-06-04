package handlers

import (
	"fmt"
	"net/http"
	"strconv"

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

	users, total, err := h.service.List(c.Request.Context(), defaultLimit, defaultOffset)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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

	users, err := h.service.ListCursor(c.Request.Context(), defaultCursor, defaultLimit)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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

// SearchByEmail godoc
//
//	@Summary		Search user by email
//	@Description	Search users by email (partial match)
//	@Tags			users
//	@Produce		json
//	@Success		200	{object}	dto.ListUserResponse
//	@Router			/users/search [get]
//
//	@Param			email	query	string	true	"Email to search for"
func (h *UserHandler) SearchByEmail(c *gin.Context) {

	email := c.Query("email")

	fmt.Println(email)

	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email is required"})
		return
	}

	users, err := h.service.SearchByEmail(c.Request.Context(), email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
