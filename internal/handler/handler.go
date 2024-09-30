package handler

import (
	"encoding/json"
	"fmt"

	"boilerplate/internal/repository"
	"boilerplate/internal/usecase"
	"boilerplate/pkg/configs"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitHandler(db *gorm.DB) {
	app := fiber.New(fiber.Config{
		// DisableStartupMessage: true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       fmt.Sprintf("%s %s", configs.App.Name, configs.App.Version),
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

	api := app.Group("boilerplate/v1")

	exampleRepo := repository.NewBookRepository(db)
	exampleUseCase := usecase.NewBookUseCase(exampleRepo)
	exampleHandler := NewBookHandler(exampleUseCase)
	exampleHandler.Route(api)

	app.Use(NotFoundHandler)
	app.Listen(fmt.Sprintf(":%d", configs.App.Port))
}
