package category

type Category struct {
	id    int64
	title string
}

func (c *Category) ID() int64 {
	return c.id
}

func (c *Category) Title() string {
	return c.title
}

func NewCategory(id int64, title string) *Category {
	return &Category{id: id, title: title}
}
