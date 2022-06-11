package repository

import (
	"context"
	"database/sql"
)

type OrderProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, OrderProduct domain.OrderProduct) domain.OrderProduct
	Update(ctx context.Context, tx *sql.Tx, OrderProduct domain.OrderProduct) domain.OrderProduct
	Delete(ctx context.Context, tx *sql.Tx, OrderProduct domain.OrderProduct)
	FindById(ctx context.Context, tx *sql.Tx, OrderProduct int) (domain.OrderProduct, error)
	FindByAll(ctx context.Context, tx *sql.Tx) []domain.OrderProduct
}
