package repository

import (
	"context"
	"database/sql"
	"errors"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (p ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, Product domain.Product) domain.Product {
	SQL := "Insert into Product(nama, price, category_id) value (?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, Product.Name, Product.Price, Product.CategoryId)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	Product.Id = int(id)
	return Product
}

func (p ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, Product domain.Product) domain.Product {
	SQL := "update product set  nama = ?, price = ?, category_id = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, Product.Id, Product.Name, Product.Price, Product.CategoryId)
	helper.PanicIfError(err)

	return Product
}

func (p ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, product domain.Product) {
	SQL := "delete from product where id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Id)
	helper.PanicIfError(err)
}

func (p ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, ProductId int) (web.ProductResponse, error) {
	//SELECT a.*, b.name as 'category_name' FROM `product` a
	//INNER JOIN category b on a.category_id=b.id;
	logrus.Info("product repository start")
	SQL := "select p.id, p.nama, p.price, p.category_id, c.nama from product p inner join category c  on p.category_id=c.id where p.id = ?"
	rows, err := tx.QueryContext(ctx, SQL, ProductId)
	helper.PanicIfError(err)
	defer rows.Close()

	Product := web.ProductResponse{}
	if rows.Next() {
		err := rows.Scan(&Product.Id, &Product.Nama, &Product.Price, &Product.CategoryId, &Product.CategoryName)
		helper.PanicIfError(err)
		return Product, nil
	} else {
		return Product, errors.New("Product is not found")
	}
}

func (p ProductRepositoryImpl) FindByAll(ctx context.Context, tx *sql.Tx) []web.ProductResponse {
	//SELECT a.*, b.name as 'category_name' FROM `product` a
	//INNER JOIN category b on a.category_id=b.id;
	logrus.Info("product repository find by all start")
	SQL := "select p.id, p.nama, p.price, p.category_id, c.nama from product p inner join category c on p.category_id= c.id"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var Products []web.ProductResponse
	for rows.Next() {
		Product := web.ProductResponse{}
		err := rows.Scan(&Product.Id, &Product.Nama, &Product.Price, &Product.CategoryId, &Product.CategoryName)
		helper.PanicIfError(err)
		Products = append(Products, Product)

	}
	logrus.Info("product repository ended")
	return Products
}
