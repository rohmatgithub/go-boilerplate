package handler

import (
	"boilerplate/internal/common"
	"boilerplate/internal/dto"
	"boilerplate/internal/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type BookHanlder struct {
	bookUsecase usecase.BookUseCase
}

func NewBookHandler(uc usecase.BookUseCase) *BookHanlder {
	return &BookHanlder{
		bookUsecase: uc,
	}
}

func (h *BookHanlder) Route(app fiber.Router) {
	var ae AbstractController

	api := app.Group("/book")

	api.Get("/:id", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", h.getDetailBook)
	})
}

func (h *BookHanlder) getDetailBook(c *fiber.Ctx, ctxModel *common.ContextModel) (dto.Payload, common.ErrorModel) {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	return h.bookUsecase.GetBookByID(id)
}
