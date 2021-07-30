package repository

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/muhammadrijalkamal/backendtest/entity"
	"github.com/muhammadrijalkamal/backendtest/model"
)

type ArticleRepositoryImpl struct {
	DB *sql.DB
}

func NewArticleRepository(db *sql.DB) ArticleRepository {
	return &ArticleRepositoryImpl{
		DB: db,
	}
}

func (r *ArticleRepositoryImpl) Insert(request *entity.Article) error {
	query := "INSERT INTO articles (title, slug, category_id, content) VALUES (?, ?, ?, ?)"
	result, err1 := r.DB.ExecContext(context.Background(), query, request.Title, request.Slug, request.CategoryID, request.Content)
	if err1 != nil {
		return err1
	}

	affected, err2 := result.RowsAffected()
	if err2 != nil {
		return err2
	}

	if affected != 1 {
		return errors.New("no article saved")
	}

	return nil
}

func (r *ArticleRepositoryImpl) FindAll() (*[]model.ArticleResponse, error) {
	query := `SELECT a.id, a.title, a.slug, c.id AS category_id, c.category_name, c.category_slug, a.content, a.created_at, a.updated_at, a.deleted_at
				FROM articles AS a INNER JOIN categories AS c on a.category_id = c.id WHERE a.deleted_at IS NULL`
	rows, err1 := r.DB.QueryContext(context.Background(), query)
	if err1 != nil {
		return nil, err1
	}

	defer rows.Close()
	var articles []model.ArticleResponse
	for rows.Next() {
		var id, categoryID int64
		var title, slug, categoryName, categorySlug, content string
		var createdAt time.Time
		var updatedAt, deletedAt sql.NullTime
		err2 := rows.Scan(
			&id,
			&title,
			&slug,
			&categoryID,
			&categoryName,
			&categorySlug,
			&content,
			&createdAt,
			&updatedAt,
			&deletedAt,
		)

		if err2 != nil {
			return nil, err2
		}

		article := model.ArticleResponse{
			ID:           id,
			Title:        title,
			Slug:         slug,
			CategoryID:   categoryID,
			CategoryName: categoryName,
			CategorySlug: categorySlug,
			Content:      content,
			CreatedAt:    createdAt,
		}

		if updatedAt.Valid {
			article.UpdatedAt = updatedAt.Time
		}

		if deletedAt.Valid {
			article.DeletedAt = deletedAt.Time
		}

		articles = append(articles, article)
	}

	return &articles, nil
}

func (r *ArticleRepositoryImpl) FindAllByTitle(title string) (*[]model.ArticleResponse, error) {
	filter := strings.ToLower(title)
	query := `SELECT a.id, a.title, a.slug, c.id AS category_id, c.category_name, c.category_slug, a.content, a.created_at, a.updated_at, a.deleted_at
				FROM articles AS a INNER JOIN categories AS c on a.category_id = c.id WHERE a.title REGEXP ? AND a.deleted_at IS NULL`
	rows, err1 := r.DB.QueryContext(context.Background(), query, filter)
	if err1 != nil {
		return nil, err1
	}

	defer rows.Close()
	var articles []model.ArticleResponse
	for rows.Next() {
		var id, categoryID int64
		var title, slug, categoryName, categorySlug, content string
		var createdAt time.Time
		var updatedAt, deletedAt sql.NullTime
		err2 := rows.Scan(
			&id,
			&title,
			&slug,
			&categoryID,
			&categoryName,
			&categorySlug,
			&content,
			&createdAt,
			&updatedAt,
			&deletedAt,
		)

		if err2 != nil {
			return nil, err2
		}

		article := model.ArticleResponse{
			ID:           id,
			Title:        title,
			Slug:         slug,
			CategoryID:   categoryID,
			CategoryName: categoryName,
			CategorySlug: categorySlug,
			Content:      content,
			CreatedAt:    createdAt,
		}

		if updatedAt.Valid {
			article.UpdatedAt = updatedAt.Time
		}

		if deletedAt.Valid {
			article.DeletedAt = deletedAt.Time
		}

		articles = append(articles, article)
	}

	return &articles, nil
}

func (r *ArticleRepositoryImpl) FindAllSoftDeleted() (*[]model.ArticleResponse, error) {
	query := `SELECT a.id, a.title, a.slug, c.id AS category_id, c.category_name, c.category_slug, a.content, a.created_at, a.updated_at, a.deleted_at
				FROM articles AS a INNER JOIN categories AS c on a.category_id = c.id WHERE a.deleted_at IS NOT NULL`
	rows, err1 := r.DB.QueryContext(context.Background(), query)
	if err1 != nil {
		return nil, err1
	}

	defer rows.Close()
	var articles []model.ArticleResponse
	for rows.Next() {
		var id, categoryID int64
		var title, slug, categoryName, categorySlug, content string
		var createdAt time.Time
		var updatedAt, deletedAt sql.NullTime
		err2 := rows.Scan(
			&id,
			&title,
			&slug,
			&categoryID,
			&categoryName,
			&categorySlug,
			&content,
			&createdAt,
			&updatedAt,
			&deletedAt,
		)

		if err2 != nil {
			return nil, err2
		}

		article := model.ArticleResponse{
			ID:           id,
			Title:        title,
			Slug:         slug,
			CategoryID:   categoryID,
			CategoryName: categoryName,
			CategorySlug: categorySlug,
			Content:      content,
			CreatedAt:    createdAt,
		}

		if updatedAt.Valid {
			article.UpdatedAt = updatedAt.Time
		}

		if deletedAt.Valid {
			article.DeletedAt = deletedAt.Time
		}

		articles = append(articles, article)
	}

	return &articles, nil
}

func (r *ArticleRepositoryImpl) FindByID(articleID int64) (*model.ArticleResponse, error) {
	query := `SELECT a.id, a.title, a.slug, c.id AS category_id, c.category_name, c.category_slug, a.content, a.created_at, a.updated_at, a.deleted_at
				FROM articles AS a INNER JOIN categories AS c on a.category_id = c.id WHERE a.id = ?`
	rows, err1 := r.DB.QueryContext(context.Background(), query, articleID)
	if err1 != nil {
		return nil, err1
	}

	defer rows.Close()
	if rows.Next() {
		var id, categoryID int64
		var title, slug, categoryName, categorySlug, content string
		var createdAt time.Time
		var updatedAt, deletedAt sql.NullTime
		err2 := rows.Scan(
			&id,
			&title,
			&slug,
			&categoryID,
			&categoryName,
			&categorySlug,
			&content,
			&createdAt,
			&updatedAt,
			&deletedAt,
		)

		if err2 != nil {
			return nil, err2
		}

		article := model.ArticleResponse{
			ID:           id,
			Title:        title,
			Slug:         slug,
			CategoryID:   categoryID,
			CategoryName: categoryName,
			CategorySlug: categorySlug,
			Content:      content,
			CreatedAt:    createdAt,
		}

		if updatedAt.Valid {
			article.UpdatedAt = updatedAt.Time
		}

		if deletedAt.Valid {
			article.DeletedAt = deletedAt.Time
		}

		return &article, nil
	}

	return nil, nil
}

func (r *ArticleRepositoryImpl) Update(articleID int64, request *entity.Article) error {
	query := "UPDATE articles SET title = ?, slug = ? , category_id = ?, content = ? WHERE id = ?"
	result, err1 := r.DB.ExecContext(context.Background(), query, request.Title, request.Slug, request.CategoryID, request.Content, articleID)
	if err1 != nil {
		return err1
	}

	affected, err2 := result.RowsAffected()
	if err2 != nil {
		return err2
	}

	if affected != 1 {
		return errors.New("no article updated")
	}

	return nil
}

func (r *ArticleRepositoryImpl) SoftDelete(articleID int64) error {
	query := "UPDATE articles SET deleted_at = NOW() WHERE id = ?"
	result, err1 := r.DB.ExecContext(context.Background(), query, articleID)
	if err1 != nil {
		return err1
	}

	affected, err2 := result.RowsAffected()
	if err2 != nil {
		return err2
	}

	if affected != 1 {
		return errors.New("no article deleted")
	}

	return nil
}

func (r *ArticleRepositoryImpl) Delete(articleID int64) error {
	query := "DELETE FROM articles WHERE id = ?"
	result, err1 := r.DB.ExecContext(context.Background(), query, articleID)
	if err1 != nil {
		return err1
	}

	affected, err2 := result.RowsAffected()
	if err2 != nil {
		return err2
	}

	if affected != 1 {
		return errors.New("no article deleted")
	}

	return nil
}
