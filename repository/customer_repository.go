package repository

import (
	"context"
	"database/sql"
)

type CustomerRepository interface {
	Save(ctx context.Context, tx *sql.Tx, customer domain.Customer) domain.Customer
	Update(ctx context.Context, tx *sql.Tx, customer domain.Customer) domain.Customer
	Delete(ctx context.Context, tx *sql.Tx, customer domain.Customer)
	FindById(ctx context.Context, tx *sql.Tx, customerId int) (domain.Customer, error)
	FindByAll(ctx context.Context, tx *sql.Tx) []domain.Customer
}
