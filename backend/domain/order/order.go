package order

import "time"

type Order struct {
	id        int64
	userID    int64
	createdAt time.Time
}

func (o *Order) ID() int64 {
	return o.id
}

func (o *Order) UserID() int64 {
	return o.userID
}

func (o *Order) CreatedAt() time.Time {
	return o.createdAt
}

func NewOrder(id int64, userID int64, createdAt time.Time) *Order {
	return &Order{id: id, userID: userID, createdAt: createdAt}
}
