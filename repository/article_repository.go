package repository

import (
	"github.com/muhammadrijalkamal/backendtest/entity"
	"github.com/muhammadrijalkamal/backendtest/model"
)

type ArticleRepository interface {
	Insert(request *entity.Article) error

	FindAll() (*[]model.ArticleResponse, error)

	FindAllByTitle(title string) (*[]model.ArticleResponse, error)

	FindAllSoftDeleted() (*[]model.ArticleResponse, error)

	FindByID(articleID int64) (*model.ArticleResponse, error)

	Update(articleID int64, request *entity.Article) error

	SoftDelete(articleID int64) error

	Delete(articleID int64) error
}
