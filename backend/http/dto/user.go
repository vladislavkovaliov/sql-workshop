package dto

type UserResponse struct {
	ID    int64  `json:"id" example:"1"`
	Name  string `json:"name" example:"username"`
	Email string `json:"email" example:"text@gmail.com"`
}

type ListUserResponse struct {
	Data  []UserResponse `json:"data"`
	Total int            `json:"total"`
}

type CursorUserResponse struct {
	Users      []UserResponse `json:"users"`
	NextCursor int64          `json:"next_cursor"`
}
