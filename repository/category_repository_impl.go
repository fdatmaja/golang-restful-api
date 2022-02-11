package repository

import (
	"belajar_golang_restful_api/helper"
	"belajar_golang_restful_api/model/domain"
	"context"
	"database/sql"
	"errors"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepositroy() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (mp *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	query := "insert into category(name) values(?)"
	result, err := tx.ExecContext(ctx, query, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	return category

}

func (mp *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	query := "update category set name=? where id=?"
	_, err := tx.ExecContext(ctx, query, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}

func (mp *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	query := "delete from category where id=?"
	_, err := tx.ExecContext(ctx, query, category.Id)
	helper.PanicIfError(err)
}

func (mp *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	query := "select id, name from category where id=?"
	rows, err := tx.QueryContext(ctx, query, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category not found")
	}
}

func (mp *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	query := "select id, name from category"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	defer rows.Close()

	categories := []domain.Category{}
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}

	return categories
}
