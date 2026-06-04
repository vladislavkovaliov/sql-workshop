package dto

type CategoryResponse struct {
	ID    int64  `json:"id" example:"1"`
	Title string `json:"title" example:"Keyboard"`
}

type ListCategoryResponse struct {
	Data  []CategoryResponse `json:"data"`
	Total int                `json:"total"`
}
