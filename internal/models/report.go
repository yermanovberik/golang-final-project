package models

import "time"

type Report struct {
	Id           int
	Date         time.Time
	TotalOrders  int
	TotalRevenue int
	CreatedAt    time.Time
}
