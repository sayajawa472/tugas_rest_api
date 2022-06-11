package repository

import (
	"context"
	"database/sql"
	"errors"
)

type OrdersRepositoryImpl struct {
}

func NewOrdersRepository() OrdersRepository {
	return &OrdersRepositoryImpl{}
}

func (o OrdersRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, orders domain.Orders) domain.Orders {
	SQL := "Insert into orders(customer_id, total_amount) value (?,?)"
	result, err := tx.ExecContext(ctx, SQL, orders.CustomerId, orders.TotalAmount)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	orders.Id = int(id)
	return orders
}

func (o OrdersRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, orders domain.Orders) domain.Orders {
	SQL := "update orders set Customer_id = ?, Total_amount = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, orders.CustomerId, orders.TotalAmount, orders.Id)
	helper.PanicIfError(err)

	return orders
}

func (o OrdersRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, orders domain.Orders) {
	SQL := "delete from Orders where id = ?"
	_, err := tx.ExecContext(ctx, SQL, orders.Id)
	helper.PanicIfError(err)
}

func (o OrdersRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, ordersId int) (domain.Orders, error) {
	logrus.Info("orders repository find by id started")
	//harus sama yg tabel php my admin
	SQL := "select id, Customer_id,Total_amount from orders where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, ordersId)
	helper.PanicIfError(err)
	defer rows.Close()

	orders := domain.Orders{}
	if rows.Next() {
		err := rows.Scan(&ordersId, &orders.CustomerId, &orders.TotalAmount)
		helper.PanicIfError(err)
		return orders, nil
	} else {
		logrus.Info("orders repository find by id ended")
		return orders, errors.New("Orders is not found")
	}
}

func (o OrdersRepositoryImpl) FindByAll(ctx context.Context, tx *sql.Tx) []domain.Orders {
	logrus.Info("orders repository find by all started")
	SQL := "select id, Customer_id, Total_amount from orders"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var orders []domain.Orders
	for rows.Next() {
		order := domain.Orders{}
		err := rows.Scan(&order.Id, &order.CustomerId, &order.TotalAmount)
		helper.PanicIfError(err)
		orders = append(orders, order)

	}
	logrus.Info("orders repository find by all ended")
	return orders
}
