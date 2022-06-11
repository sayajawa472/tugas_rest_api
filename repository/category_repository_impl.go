package repository

import (
	"context"
	"database/sql"
	"errors"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}

}

func (c CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "insert into category(nama) values (?)"
	result, err := tx.ExecContext(ctx, SQL, category.Nama)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	return category

}

func (c CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "update category set nama = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Nama, category.Id)
	helper.PanicIfError(err)

	return category
}

func (c CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "delete from category where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfError(err)
}

func (c CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	logrus.Info("category repository find by id start")
	SQL := "select id, nama from category where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Nama)
		helper.PanicIfError(err)
		return category, nil
	} else {
		logrus.Info("category repository find by ended")
		return category, errors.New("category is not found")
	}
}

func (c CategoryRepositoryImpl) FindByAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	logrus.Info("category repository find by all start")
	SQL := "select id, nama from category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Nama)
		helper.PanicIfError(err)
		categories = append(categories, category)

	}
	logrus.Info("category repository find by all ended")
	return categories
}
