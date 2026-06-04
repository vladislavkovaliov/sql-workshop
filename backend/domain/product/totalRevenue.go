package product

type TotalRevenue struct {
	title   string
	revenue float64
}

func (t *TotalRevenue) Title() string {
	return t.title
}

func (t *TotalRevenue) Revenue() float64 {
	return t.revenue
}

func NewTotalRevenue(title string, revenue float64) *TotalRevenue {
	return &TotalRevenue{title: title, revenue: revenue}
}
