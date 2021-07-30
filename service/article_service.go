package service

import (
	"github.com/muhammadrijalkamal/backendtest/model"
)

type ArticleService interface {
	Create(request *model.ArticleCreateRequest)

	List() *[]model.ArticleResponse

	ListByTitle(title string) *[]model.ArticleResponse

	ListSoftDeleted() *[]model.ArticleResponse

	FindOne(articleID string) *model.ArticleResponse

	Update(articleID string, request *model.ArticleUpdateRequest)

	SoftDelete(articleID string)

	Delete(articleID string)
}
