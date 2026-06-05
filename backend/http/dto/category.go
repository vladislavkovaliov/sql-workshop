package dto

type CategoryResponse struct {
	ID    int64  `json:"id" example:"1"`
	Title string `json:"title" example:"Keyboard"`
}

type CategoryAveragePrice struct {
	Category string  `json:"category"`
	AvgPrice float64 `json:"avg_price"`
}

type ListCategoryResponse struct {
	Data  []CategoryResponse `json:"data"`
	Total int                `json:"total"`
}

type ListCategoryAvaragePriceResponse struct {
	Data  []CategoryAveragePrice `json:"data"`
	Total int                    `json:"total"`
}
