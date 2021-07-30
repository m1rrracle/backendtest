package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhammadrijalkamal/backendtest/model"
	"github.com/muhammadrijalkamal/backendtest/service"
	"github.com/muhammadrijalkamal/backendtest/util"
)

type ArticleController struct {
	ArticleService service.ArticleService
}

func NewArticleController(articleService *service.ArticleService) ArticleController {
	return ArticleController{
		ArticleService: *articleService,
	}
}

func (controller *ArticleController) SetupRoutes(app *fiber.App) {
	app.Get("/article/deleted", controller.ListSoftDeleted)
	app.Delete("/article/deleted/:id", controller.Delete)
	app.Post("/article", controller.Create)
	app.Get("/article", controller.List)
	app.Get("/article/:id", controller.FindOne)
	app.Put("/article/:id", controller.Update)
	app.Delete("/article/:id", controller.SoftDelete)
}

func (controller *ArticleController) Create(ctx *fiber.Ctx) error {
	var request *model.ArticleCreateRequest
	parserErr := ctx.BodyParser(&request)
	util.ReturnErrorIfNeeded(parserErr)

	controller.ArticleService.Create(request)

	return ctx.Status(fiber.StatusCreated).JSON(model.SuccessResponse{
		StatusCode: fiber.StatusCreated,
		Data:       "Article created",
	})
}

func (controller *ArticleController) List(ctx *fiber.Ctx) error {
	title := ctx.Query("title")

	var articles *[]model.ArticleResponse

	if title != "" {
		articles = controller.ArticleService.ListByTitle(title)
	} else {
		articles = controller.ArticleService.List()
	}

	return ctx.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		StatusCode: fiber.StatusOK,
		Data:       articles,
	})
}

func (controller *ArticleController) FindOne(ctx *fiber.Ctx) error {
	articleID := ctx.Params("id")

	article := controller.ArticleService.FindOne(articleID)

	return ctx.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		StatusCode: fiber.StatusOK,
		Data:       article,
	})
}

func (controller *ArticleController) Update(ctx *fiber.Ctx) error {
	articleID := ctx.Params("id")

	var request *model.ArticleUpdateRequest
	parserErr := ctx.BodyParser(&request)
	util.ReturnErrorIfNeeded(parserErr)

	controller.ArticleService.Update(articleID, request)

	return ctx.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		StatusCode: fiber.StatusOK,
		Data:       "Article updated",
	})
}

func (controller *ArticleController) SoftDelete(ctx *fiber.Ctx) error {
	articleID := ctx.Params("id")

	controller.ArticleService.SoftDelete(articleID)

	return ctx.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		StatusCode: fiber.StatusOK,
		Data:       "Article deleted",
	})
}

func (controller *ArticleController) ListSoftDeleted(ctx *fiber.Ctx) error {
	articles := controller.ArticleService.ListSoftDeleted()
	return ctx.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		StatusCode: fiber.StatusOK,
		Data:       articles,
	})
}

func (controller *ArticleController) Delete(ctx *fiber.Ctx) error {
	articleID := ctx.Params("id")

	controller.ArticleService.Delete(articleID)

	return ctx.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		StatusCode: fiber.StatusOK,
		Data:       "Article deleted from database",
	})
}
