package service

import (
	"strconv"

	"github.com/gosimple/slug"
	"github.com/muhammadrijalkamal/backendtest/entity"
	"github.com/muhammadrijalkamal/backendtest/model"
	"github.com/muhammadrijalkamal/backendtest/repository"
	"github.com/muhammadrijalkamal/backendtest/util"
)

type CategoryServiceImpl struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{
		categoryRepository: *repo,
	}
}

func (service *CategoryServiceImpl) Create(request *model.CategoryCreateRequest) {
	categorySlug := slug.Make(request.CategoryName)
	article := entity.Category{
		CategoryName: request.CategoryName,
		CategorySlug: categorySlug,
	}
	txErr := service.categoryRepository.Insert(&article)
	util.ReturnErrorIfNeeded(txErr)
}

func (service *CategoryServiceImpl) List() *[]model.CategoryResponse {
	categories, txErr := service.categoryRepository.FindAll()
	util.ReturnErrorIfNeeded(txErr)
	return categories
}

func (service *CategoryServiceImpl) ListSoftDeleted() *[]model.CategoryResponse {
	categories, txErr := service.categoryRepository.FindAllSoftDeleted()
	util.ReturnErrorIfNeeded(txErr)
	return categories
}

func (service *CategoryServiceImpl) FindOne(categoryID string) *model.CategoryResponse {
	id, err := strconv.Atoi(categoryID)
	util.ReturnErrorIfNeeded(err)

	category, txErr := service.categoryRepository.FindByID(int64(id))
	util.ReturnErrorIfNeeded(txErr)

	return category
}

func (service *CategoryServiceImpl) Update(categoryID string, request *model.CategoryUpdateRequest) {
	id, err := strconv.Atoi(categoryID)
	util.ReturnErrorIfNeeded(err)

	categorySlug := slug.Make(request.CategoryName)

	category := entity.Category{
		CategoryName: request.CategoryName,
		CategorySlug: categorySlug,
	}

	txErr := service.categoryRepository.Update(int64(id), &category)
	util.ReturnErrorIfNeeded(txErr)
}

func (service *CategoryServiceImpl) SoftDelete(categoryID string) {
	id, err := strconv.Atoi(categoryID)
	util.ReturnErrorIfNeeded(err)

	txErr := service.categoryRepository.SoftDelete(int64(id))
	util.ReturnErrorIfNeeded(txErr)
}

func (service *CategoryServiceImpl) Delete(categoryID string) {
	id, err := strconv.Atoi(categoryID)
	util.ReturnErrorIfNeeded(err)

	txErr := service.categoryRepository.Delete(int64(id))
	util.ReturnErrorIfNeeded(txErr)
}
