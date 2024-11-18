package handler

import (
	"boilerplate/internal/common"
	"boilerplate/internal/constanta"
	"boilerplate/internal/dto"
	"boilerplate/internal/usecase"
	"boilerplate/pkg/util"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type BookHanlder struct {
	bookUsecase usecase.BookUseCase
	validate    *util.AppValidator
}

func NewBookHandler(uc usecase.BookUseCase, v *util.AppValidator) *BookHanlder {
	return &BookHanlder{
		bookUsecase: uc,
		validate:    v,
	}
}

func (h *BookHanlder) Route(app fiber.Router) {
	var ae AbstractController

	api := app.Group("/book")

	api.Get("/:id", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(constanta.ActionGetDetail, c, "", h.getDetailBook)
	})

	api.Post("", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(constanta.ActionAdd, c, "", h.createBook)
	})
}

func (h *BookHanlder) getDetailBook(c *fiber.Ctx, ctxModel *common.ContextModel) (dto.Payload, common.ErrorModel) {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	return h.bookUsecase.GetBookByID(id)
}

func (h *BookHanlder) createBook(c *fiber.Ctx, ctxModel *common.ContextModel) (result dto.Payload, errMdl common.ErrorModel) {
	var req dto.BookRequest
	if err := c.BodyParser(&req); err != nil {
		errMdl = common.GenerateInvalidJSONFormatError()
		return
	}

	if detailError, err := h.validate.ValidateRequest(ctxModel, req); err != nil {
		result.Status.Detail = detailError
		errMdl = common.GenerateValidationFailedError()
		return
	}
	return h.bookUsecase.InsertBook(&req, ctxModel)
}
