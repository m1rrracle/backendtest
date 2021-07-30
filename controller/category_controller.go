package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhammadrijalkamal/backendtest/model"
	"github.com/muhammadrijalkamal/backendtest/service"
	"github.com/muhammadrijalkamal/backendtest/util"
)

type CategoryController struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService *service.CategoryService) CategoryController {
	return CategoryController{
		CategoryService: *categoryService,
	}
}

func (controller *CategoryController) SetupRoutes(app *fiber.App) {
	app.Get("/category/deleted", controller.ListSoftDeleted)
	app.Delete("/category/deleted/:id", controller.Delete)
	app.Post("/category", controller.Create)
	app.Get("/category", controller.List)
	app.Get("/category/:id", controller.FindOne)
	app.Put("/category/:id", controller.Update)
	app.Delete("/category/:id", controller.SoftDelete)
}

func (controller *CategoryController) Create(ctx *fiber.Ctx) error {
	var request *model.CategoryCreateRequest
	parserErr := ctx.BodyParser(&request)
	util.ReturnErrorIfNeeded(parserErr)

	controller.CategoryService.Create(request)

	return ctx.Status(fiber.StatusCreated).JSON(model.SuccessResponse{
		StatusCode: fiber.StatusCreated,
		Data:       "Category created",
	})
}

func (controller *CategoryController) List(ctx *fiber.Ctx) error {
	categories := controller.CategoryService.List()
	return ctx.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		StatusCode: fiber.StatusOK,
		Data:       categories,
	})
}

func (controller *CategoryController) FindOne(ctx *fiber.Ctx) error {
	categoryID := ctx.Params("id")

	category := controller.CategoryService.FindOne(categoryID)

	return ctx.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		StatusCode: fiber.StatusOK,
		Data:       category,
	})
}

func (controller *CategoryController) Update(ctx *fiber.Ctx) error {
	categoryID := ctx.Params("id")

	var request *model.CategoryUpdateRequest
	parserErr := ctx.BodyParser(&request)
	util.ReturnErrorIfNeeded(parserErr)

	controller.CategoryService.Update(categoryID, request)

	return ctx.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		StatusCode: fiber.StatusOK,
		Data:       "Category updated",
	})
}

func (controller *CategoryController) SoftDelete(ctx *fiber.Ctx) error {
	categoryID := ctx.Params("id")

	controller.CategoryService.SoftDelete(categoryID)

	return ctx.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		StatusCode: fiber.StatusOK,
		Data:       "Category deleted",
	})
}

func (controller *CategoryController) ListSoftDeleted(ctx *fiber.Ctx) error {
	categories := controller.CategoryService.ListSoftDeleted()
	return ctx.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		StatusCode: fiber.StatusOK,
		Data:       categories,
	})
}

func (controller *CategoryController) Delete(ctx *fiber.Ctx) error {
	categoryID := ctx.Params("id")

	controller.CategoryService.Delete(categoryID)

	return ctx.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		StatusCode: fiber.StatusOK,
		Data:       "Category deleted from database",
	})
}
