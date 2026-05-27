package dto

type ProductResponse struct {
	ID    int64   `json:"id" example:"1"`
	Title string  `json:"title" example:"Keyboard"`
	Price float64 `json:"price" example:"150.00"`
}

type CursorProductsResponse struct {
	Products   []ProductResponse `json:"products"`
	NextCursor int64             `json:"next_cursor"`
}
