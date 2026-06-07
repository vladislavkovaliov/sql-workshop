package order

import "time"

type DailyPurchases struct {
	orderDate time.Time
	purchases int
}

func (d *DailyPurchases) OrderDate() time.Time {
	return d.orderDate
}

func (d *DailyPurchases) Purchases() int {
	return d.purchases
}

func NewDailyPurchases(orderDate time.Time, purchases int) *DailyPurchases {
	return &DailyPurchases{orderDate: orderDate, purchases: purchases}
}
