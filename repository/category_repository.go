package repository

import (
	"github.com/muhammadrijalkamal/backendtest/entity"
	"github.com/muhammadrijalkamal/backendtest/model"
)

type CategoryRepository interface {
	Insert(request *entity.Category) error

	FindAll() (*[]model.CategoryResponse, error)

	FindAllSoftDeleted() (*[]model.CategoryResponse, error)

	FindByID(categoryID int64) (*model.CategoryResponse, error)

	Update(categoryID int64, request *entity.Category) error

	SoftDelete(categoryID int64) error

	Delete(categoryID int64) error
}
