package models

import "time"

type Order struct {
	Id          int
	UserId      string
	OrderStatus string
	TotalPrice  int
	CreatedAt   time.Time
}
