package domain

import "time"

type Customer struct {
	Id          int
	Name        string
	Address     string
	Email       string
	PhoneNumber string
	createdAt   time.Time
	UpdatedAt   time.Time
}
