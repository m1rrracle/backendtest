package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/muhammadrijalkamal/backendtest/controller"
	"github.com/muhammadrijalkamal/backendtest/model"
	"github.com/muhammadrijalkamal/backendtest/repository"
	"github.com/muhammadrijalkamal/backendtest/service"
)

var (
	dbHost     = os.Getenv("DB_HOST")
	dbPort     = os.Getenv("DB_PORT")
	dbUser     = os.Getenv("DB_USER")
	dbPass     = os.Getenv("DB_PASS")
	dbName     = os.Getenv("DB_NAME")
	Connection *sql.DB
)

func init() {
	var err error
	dbUri := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)

	Connection, err = sql.Open("mysql", dbUri)
	if err != nil {
		panic(err)
	}

	err = Connection.Ping()
	if err != nil {
		panic(err)
	}

	Connection.SetMaxIdleConns(500)
	Connection.SetMaxOpenConns(500)
	Connection.SetConnMaxIdleTime(5 * time.Minute)
	Connection.SetConnMaxLifetime(60 * time.Minute)
}

func main() {
	articleRepository := repository.NewArticleRepository(Connection)
	articleService := service.NewArticleService(&articleRepository)
	articleController := controller.NewArticleController(&articleService)

	categoryRepository := repository.NewCategoryRepository(Connection)
	categoryService := service.NewCategoryService(&categoryRepository)
	categoryController := controller.NewCategoryController(&categoryService)

	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if err != nil {
				return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
					StatusCode: fiber.StatusBadRequest,
					Error:      err.Error(),
				})
			}
			return nil
		},
	})

	app.Use(cors.New())
	app.Use(recover.New())

	articleController.SetupRoutes(app)
	categoryController.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
