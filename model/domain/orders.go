package domain

import "time"

type Orders struct {
	Id          int
	OrderDate   time.Time
	CustomerId  int
	TotalAmount int
	createdAt   time.Time
	UpdatedAt   time.Time
}
