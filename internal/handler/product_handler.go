package handler

import (
	"boilerplate/internal/common"
	"boilerplate/internal/dto"
	"boilerplate/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type ExampleHanlder struct {
	exampleUsecase usecase.ExampleUseCase
}

func NewExampleHandler(uc usecase.ExampleUseCase) *ExampleHanlder {
	return &ExampleHanlder{
		exampleUsecase: uc,
	}
}

func (h *ExampleHanlder) Route(app fiber.Router) {
	var ae AbstractController

	api := app.Group("/example")

	api.Get("/id", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", h.getDetailProduct)
	})
}

func (h *ExampleHanlder) getDetailProduct(c *fiber.Ctx, ctxModel *common.ContextModel) (dto.Payload, common.ErrorModel) {
	exampleCode := c.Query("exampleCode")
	return h.exampleUsecase.GetExampleByID("NS6173010003515", exampleCode)
}
