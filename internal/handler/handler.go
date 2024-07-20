package handler

import (
	"encoding/json"
	"fmt"

	"boilerplate/internal/repository"
	"boilerplate/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitHandler(db *gorm.DB) {
	app := fiber.New(fiber.Config{
		// DisableStartupMessage: true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "POS App v2.0.0",
		ColorScheme:   fiber.Colors{Green: ""},
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
	})

	// app.Use(logger.New())
	app.Use(func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
				customErrorHandler(c, fmt.Errorf("%v", r))
			}
		}()
		return c.Next()
	})

	app.Use(middleware)
	// Mendefinisikan route untuk endpoint '/'
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api := app.Group("nexdist/pos/api")

	productRepo := repository.NewProductRepository(db)
	productUseCase := usecase.NewProductUseCase(productRepo)
	productHandler := NewProductHandler(productUseCase)
	productHandler.Route(api)

	app.Listen(":8080")
}
