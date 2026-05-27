package product

type Product struct {
	id    int64
	title string
	price float64
}

func (p *Product) ID() int64 {
	return p.id
}

func (p *Product) Title() string {
	return p.title
}

func (p *Product) Price() float64 {
	return p.price
}

func NewProduct(id int64, title string, price float64) *Product {
	return &Product{id: id, title: title, price: price}
}
