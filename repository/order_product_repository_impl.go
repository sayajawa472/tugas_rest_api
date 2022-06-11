package repository

import (
	"context"
	"database/sql"
	"errors"
)

type OrderProductRepositoryImpl struct {
}

func NewOrderProductRepository() OrderProductRepository {
	return &OrderProductRepositoryImpl{}

}

func (Repository OrderProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, OrderProduct domain.OrderProduct) domain.OrderProduct {
	SQL := "insert into order_product(order_id, product_id, qty, price, amount) values (?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, OrderProduct.OrderId, OrderProduct.ProductId, OrderProduct.Qty, OrderProduct.Price, OrderProduct.Amount)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	OrderProduct.Id = int(id)
	return OrderProduct

}

func (Repository OrderProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, OrderProduct domain.OrderProduct) domain.OrderProduct {
	SQL := "update order_product set order_id=?, product_id=?, qty=?, price=?, amount = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, OrderProduct.OrderId, OrderProduct.ProductId, OrderProduct.Qty, OrderProduct.Price, OrderProduct.Amount, OrderProduct.Id)
	helper.PanicIfError(err)

	return OrderProduct
}

func (Repository OrderProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, OrderProduct domain.OrderProduct) {
	SQL := "delete from order_product where id = ?"
	_, err := tx.ExecContext(ctx, SQL, OrderProduct.Id)
	helper.PanicIfError(err)
}

func (Repository OrderProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, OrderProductId int) (domain.OrderProduct, error) {
	SQL := "select id, order_id, product_id, qty, price, amount from order_product where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, OrderProductId)
	helper.PanicIfError(err)
	defer rows.Close()

	OrderProduct := domain.OrderProduct{}
	if rows.Next() {
		err := rows.Scan(&OrderProduct.Id, &OrderProduct.OrderId, &OrderProduct.ProductId, &OrderProduct.Qty, &OrderProduct.Price, &OrderProduct.Amount)
		helper.PanicIfError(err)
		return OrderProduct, nil
	} else {
		return OrderProduct, errors.New("OrderProduct is not found")
	}
}

func (Repository OrderProductRepositoryImpl) FindByAll(ctx context.Context, tx *sql.Tx) []domain.OrderProduct {
	SQL := "select id, order_id, product_id, qty, price, amount from order_product"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var OrderProducts []domain.OrderProduct
	for rows.Next() {
		OrderProduct := domain.OrderProduct{}
		err := rows.Scan(&OrderProduct.Id, &OrderProduct.OrderId, &OrderProduct.ProductId, &OrderProduct.Qty, &OrderProduct.Price, &OrderProduct.Amount)
		helper.PanicIfError(err)
		OrderProducts = append(OrderProducts, OrderProduct)

	}
	return OrderProducts
}
