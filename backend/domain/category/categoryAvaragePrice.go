package category

type CategoryAvaragePrice struct {
	category     string
	avaragePrice float64
}

func (c *CategoryAvaragePrice) Category() string {
	return c.category
}

func (c *CategoryAvaragePrice) AveragePrice() float64 {
	return c.avaragePrice
}

func NewCategoryAvaragePrice(category string, avaragePrice float64) *CategoryAvaragePrice {
	return &CategoryAvaragePrice{category: category, avaragePrice: avaragePrice}
}
