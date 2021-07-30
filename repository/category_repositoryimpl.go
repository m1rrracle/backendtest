package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/muhammadrijalkamal/backendtest/entity"
	"github.com/muhammadrijalkamal/backendtest/model"
)

type CategoryRepositoryImpl struct {
	DB *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &CategoryRepositoryImpl{
		DB: db,
	}
}

func (r *CategoryRepositoryImpl) Insert(request *entity.Category) error {
	query := "INSERT INTO categories (category_name, category_slug) VALUES (?, ?)"
	result, err1 := r.DB.ExecContext(context.Background(), query, request.CategoryName, request.CategorySlug)
	if err1 != nil {
		return err1
	}

	affected, err2 := result.RowsAffected()
	if err2 != nil {
		return err2
	}

	if affected != 1 {
		return errors.New("no category saved")
	}

	return nil
}

func (r *CategoryRepositoryImpl) FindAll() (*[]model.CategoryResponse, error) {
	query := "SELECT * FROM categories WHERE deleted_at IS NULL"
	rows, err1 := r.DB.QueryContext(context.Background(), query)
	if err1 != nil {
		return nil, err1
	}

	defer rows.Close()
	var categories []model.CategoryResponse
	for rows.Next() {
		var id int64
		var categoryName, categorySlug string
		var createdAt time.Time
		var updatedAt, deletedAt sql.NullTime
		err2 := rows.Scan(
			&id,
			&categoryName,
			&categorySlug,
			&createdAt,
			&updatedAt,
			&deletedAt,
		)

		if err2 != nil {
			return nil, err2
		}

		category := model.CategoryResponse{
			ID:           id,
			CategoryName: categoryName,
			CategorySlug: categorySlug,
			CreatedAt:    createdAt,
		}

		if updatedAt.Valid {
			category.UpdatedAt = updatedAt.Time
		}

		if deletedAt.Valid {
			category.DeletedAt = deletedAt.Time
		}

		categories = append(categories, category)
	}

	return &categories, nil
}

func (r *CategoryRepositoryImpl) FindAllSoftDeleted() (*[]model.CategoryResponse, error) {
	query := "SELECT * FROM categories WHERE deleted_at IS NOT NULL"
	rows, err1 := r.DB.QueryContext(context.Background(), query)
	if err1 != nil {
		return nil, err1
	}

	defer rows.Close()
	var categories []model.CategoryResponse
	for rows.Next() {
		var id int64
		var categoryName, categorySlug string
		var createdAt time.Time
		var updatedAt, deletedAt sql.NullTime
		err2 := rows.Scan(
			&id,
			&categoryName,
			&categorySlug,
			&createdAt,
			&updatedAt,
			&deletedAt,
		)

		if err2 != nil {
			return nil, err2
		}

		category := model.CategoryResponse{
			ID:           id,
			CategoryName: categoryName,
			CategorySlug: categorySlug,
			CreatedAt:    createdAt,
		}

		if updatedAt.Valid {
			category.UpdatedAt = updatedAt.Time
		}

		if deletedAt.Valid {
			category.DeletedAt = deletedAt.Time
		}

		categories = append(categories, category)
	}

	return &categories, nil
}

func (r *CategoryRepositoryImpl) FindByID(categoryID int64) (*model.CategoryResponse, error) {
	query := "SELECT * FROM categories WHERE id = ?"
	rows, err1 := r.DB.QueryContext(context.Background(), query, categoryID)
	if err1 != nil {
		return nil, err1
	}

	defer rows.Close()
	if rows.Next() {
		var id int64
		var categoryName, categorySlug string
		var createdAt time.Time
		var updatedAt, deletedAt sql.NullTime
		err2 := rows.Scan(
			&id,
			&categoryName,
			&categorySlug,
			&createdAt,
			&updatedAt,
			&deletedAt,
		)

		if err2 != nil {
			return nil, err2
		}

		category := model.CategoryResponse{
			ID:           id,
			CategoryName: categoryName,
			CategorySlug: categorySlug,
			CreatedAt:    createdAt,
		}

		if updatedAt.Valid {
			category.UpdatedAt = updatedAt.Time
		}

		if deletedAt.Valid {
			category.DeletedAt = deletedAt.Time
		}

		return &category, nil
	}

	return nil, nil
}

func (r *CategoryRepositoryImpl) Update(categoryID int64, request *entity.Category) error {
	query := "UPDATE categories SET category_name = ?, category_slug = ? WHERE id = ?"
	result, err1 := r.DB.ExecContext(context.Background(), query, request.CategoryName, request.CategorySlug, categoryID)
	if err1 != nil {
		return err1
	}

	affected, err2 := result.RowsAffected()
	if err2 != nil {
		return err2
	}

	if affected != 1 {
		return errors.New("no category updated")
	}

	return nil
}

func (r *CategoryRepositoryImpl) SoftDelete(categoryID int64) error {
	query := "UPDATE categories SET deleted_at = NOW() WHERE id = ?"
	result, e1 := r.DB.ExecContext(context.Background(), query, categoryID)
	if e1 != nil {
		return e1
	}

	affected, e2 := result.RowsAffected()
	if e2 != nil {
		return e2
	}

	if affected != 1 {
		return errors.New("no category deleted")
	}

	return nil
}

func (r *CategoryRepositoryImpl) Delete(categoryID int64) error {
	query := "DELETE FROM categories WHERE id = ?"
	result, err1 := r.DB.ExecContext(context.Background(), query, categoryID)
	if err1 != nil {
		return err1
	}

	affected, err2 := result.RowsAffected()
	if err2 != nil {
		return err2
	}

	if affected != 1 {
		return errors.New("no category deleted")
	}

	return nil
}
