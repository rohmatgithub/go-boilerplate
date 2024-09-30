package handler

import (
	"boilerplate/internal/common"
	"boilerplate/internal/constanta"
	"boilerplate/internal/dto"
	"context"
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

type AbstractController struct{}

func (ae AbstractController) ServeJwtToken(c *fiber.Ctx, menuConst string, runFunc func(*fiber.Ctx, *common.ContextModel) (dto.Payload, common.ErrorModel)) error {
	// validate client_id
	// tokenStr := c.Get(constanta.TokenHeaderNameConstanta)

	// TODO validasi token jwt
	// applog.Info().Msg("tes")

	var (
		ctxModel common.ContextModel
		payload  dto.Payload
		err      error
	)

	ctxModel.LoggerModel = *c.Context().Value(constanta.ApplicationContextConstanta).(*common.LoggerModel)
	defer func() {
		if r := recover(); r != nil {
			ctxModel.LoggerModel.Message = string(debug.Stack())
			generateEResponseError(c, &ctxModel, &payload, common.GenerateUnknownError(nil))
		}

		adaptor.CopyContextToFiberContext(context.WithValue(c.Context(), constanta.ApplicationContextConstanta, &ctxModel.LoggerModel), c.Context())
		err = c.JSON(payload)
	}()
	payload, errModel := runFunc(c, &ctxModel)
	if errModel.Error != nil {
		generateEResponseError(c, &ctxModel, &payload, errModel)
		return err
	}
	return err
	// return nil
}

func generateEResponseError(c *fiber.Ctx, ctxModel *common.ContextModel, payload *dto.Payload, errMdl common.ErrorModel) {
	ctxModel.LoggerModel.Code = errMdl.Error.Error()
	if errMdl.CausedBy != nil {
		ctxModel.LoggerModel.Message = errMdl.CausedBy.Error()
	}
	// write failed
	c.Status(errMdl.Code)
	payload.Status.Success = false
	payload.Status.Code = errMdl.Error.Error()
	payload.Status.Message = errMdl.Error.Error()
	//  common.GenerateI18NErrorMessage(errMdl, ctxModel.AuthAccessTokenModel.Locale)
}
