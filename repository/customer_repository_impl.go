package repository

import (
	"context"
	"database/sql"
	"errors"
)

type CustomerRepositoryImpl struct {
}

func NewCustomerRepository() CustomerRepository {
	return &CustomerRepositoryImpl{}

}

func (Repository CustomerRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, customers domain.Customer) domain.Customer {
	SQL := "insert into customer(nama, address, email, phoneNumber) value (?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, customers.Name, customers.Address, customers.Email, customers.PhoneNumber)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	customers.Id = int(id)
	return customers
}

func (Repository CustomerRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, customer domain.Customer) domain.Customer {
	SQL := "update customer set nama = ?, address = ?, email = ?, phone_number = ?, where id = ?"
	_, err := tx.ExecContext(ctx, SQL, customer.Name, customer.Address, customer.Email, customer.PhoneNumber, customer.Id)
	helper.PanicIfError(err)

	return customer
}

func (Repository CustomerRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, customer domain.Customer) {
	SQL := "delete from customer where id = ?"
	_, err := tx.ExecContext(ctx, SQL, customer.Id)
	helper.PanicIfError(err)
}

func (Repository CustomerRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, customerId int) (domain.Customer, error) {
	logrus.Info("customer repository find by id started")
	SQL := "select id, nama, address, email, phone_number from customer where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, customerId)
	helper.PanicIfError(err)
	defer rows.Close()

	customer := domain.Customer{}
	if rows.Next() {
		err := rows.Scan(&customerId, &customer.Name, &customer.Address, &customer.Email, &customer.PhoneNumber)
		helper.PanicIfError(err)
		return customer, nil
	} else {
		logrus.Info("customer repository find by id ended")
		return customer, errors.New("customer is not found")
	}
}

func (Repository CustomerRepositoryImpl) FindByAll(ctx context.Context, tx *sql.Tx) []domain.Customer {
	logrus.Info("customer repository find by all started")
	SQL := "select id, nama, address, email, phone_number from customer"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var customers []domain.Customer
	for rows.Next() {
		customer := domain.Customer{}
		err := rows.Scan(&customer.Id, &customer.Name, &customer.Address, &customer.Email, &customer.PhoneNumber)
		helper.PanicIfError(err)
		customers = append(customers, customer)

	}
	logrus.Info("customer repository find by all ended")
	return customers
}
