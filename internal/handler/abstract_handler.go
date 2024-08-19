package handler

import (
	"boilerplate/internal/common"
	"boilerplate/internal/constanta"
	"boilerplate/internal/dto"
	"boilerplate/pkg/applog"

	"github.com/gofiber/fiber/v2"
)

type AbstractController struct{}

func (ae AbstractController) ServeJwtToken(c *fiber.Ctx, menuConst string, runFunc func(*fiber.Ctx, *common.ContextModel) (dto.Payload, common.ErrorModel)) error {
	// validate client_id
	tokenStr := c.Get(constanta.TokenHeaderNameConstanta)

	// validasi token jwt
	applog.Info().Msg(tokenStr)

	// defer func() {
	// 	if r := recover(); r != nil {
	// 		contextModel.LoggerModel.Message = string(debug.Stack())
	// 		generateEResponseError(c, &contextModel, &payload, model.GenerateUnknownError(nil))
	// 	}
	// 	response.Payload = payload

	// 	adaptor.CopyContextToFiberContext(context.WithValue(c.Context(), constanta.ApplicationContextConstanta, &contextModel.LoggerModel), c.Context())
	// 	err = c.JSON(response)
	// }()
	return nil
}

func generateEResponseError(c *fiber.Ctx, ctxModel *common.ContextModel, payload *dto.Payload, errMdl common.ErrorModel) {
	// ctxModel.LoggerModel.Code = errMdl.Error.Error()
	// ctxModel.LoggerModel.Class = errMdl.Line
	// if errMdl.CausedBy != nil {
	// 	ctxModel.LoggerModel.Message = errMdl.CausedBy.Error()
	// }
	// // write failed
	// c.Status(errMdl.Code)
	// payload.Status.Success = false
	// payload.Status.Code = errMdl.Error.Error()
	// payload.Status.Message = common.GenerateI18NErrorMessage(errMdl, ctxModel.AuthAccessTokenModel.Locale)
}
