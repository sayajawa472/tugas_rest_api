package repository

import (
	"context"
	"database/sql"
)

type ProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, Product domain.Product) domain.Product
	Update(ctx context.Context, tx *sql.Tx, Product domain.Product) domain.Product
	Delete(ctx context.Context, tx *sql.Tx, Product domain.Product)
	FindById(ctx context.Context, tx *sql.Tx, ProductId int) (web.ProductResponse, error)
	FindByAll(ctx context.Context, tx *sql.Tx) []web.ProductResponse
}
