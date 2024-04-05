package models

import "time"

type OrderDetail struct {
	Id        int
	OrderId   int
	MenuId    int
	Quantity  int
	Price     int
	CreatedAt time.Time
}
