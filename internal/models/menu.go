package models

import "time"

type Menu struct {
	Id          int
	Name        string
	Description string
	Price       int
	Available   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
