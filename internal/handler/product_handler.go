package handler

import (
	"boilerplate/internal/common"
	"boilerplate/internal/dto"
	"boilerplate/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	productUseCase usecase.ProductUseCase
}

func NewProductHandler(uc usecase.ProductUseCase) *ProductHandler {
	return &ProductHandler{
		productUseCase: uc,
	}
}

func (h *ProductHandler) Route(app fiber.Router) {
	var ae AbstractController

	api := app.Group("/product")

	api.Get("/id", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", h.getDetailProduct)
	})
}

func (h *ProductHandler) getDetailProduct(c *fiber.Ctx, ctxModel *common.ContextModel) (dto.Payload, common.ErrorModel) {
	productCode := c.Query("productCode")
	return h.productUseCase.GetProductByID("NS6173010003515", productCode)
}
