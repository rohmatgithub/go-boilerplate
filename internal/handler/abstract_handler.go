package handler

import (
	"boilerplate/internal/common"
	"boilerplate/internal/constanta"
	"boilerplate/internal/dto"
	"boilerplate/pkg/util"
	"context"
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"golang.org/x/text/language"
)

type AbstractController struct{}

func (ae AbstractController) ServeJwtToken(action string, c *fiber.Ctx, menuConst string, runFunc func(*fiber.Ctx, *common.ContextModel) (dto.Payload, common.ErrorModel)) error {
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
	ctxModel.Locale = language.Indonesian.String()
	defer func() {
		if r := recover(); r != nil {
			// applog.Error().Msg(string(debug.Stack()))
			ctxModel.LoggerModel.Message = string(debug.Stack())
			generateResponseError(c, &ctxModel, &payload, common.GenerateInternalServerError(nil))
		}

		adaptor.CopyContextToFiberContext(context.WithValue(c.Context(), constanta.ApplicationContextConstanta, &ctxModel.LoggerModel), c.Context())
		err = c.JSON(payload)
	}()
	payload, errModel := runFunc(c, &ctxModel)
	if errModel.ErrorCode != "" {
		generateResponseError(c, &ctxModel, &payload, errModel)
		return err
	}

	generateResponseSuccess(action, c, &ctxModel, &payload)
	return nil
}

func generateResponseError(c *fiber.Ctx, ctxModel *common.ContextModel, payload *dto.Payload, errMdl common.ErrorModel) {
	ctxModel.LoggerModel.Code = errMdl.ErrorCode
	if errMdl.CausedBy != nil {
		ctxModel.LoggerModel.Message = errMdl.CausedBy.Error()
	}
	// write failed
	c.Status(errMdl.Code)
	payload.Status.Success = false
	payload.Status.Code = errMdl.ErrorCode
	payload.Status.Message = util.GetI18nErrorMessage(ctxModel.Locale, errMdl.ErrorCode, errMdl.ErrorParameter)
}

func generateResponseSuccess(action string, c *fiber.Ctx, ctxModel *common.ContextModel, payload *dto.Payload) {
	ctxModel.LoggerModel.Code = "OK"

	// write failed
	c.Status(200)
	payload.Status.Success = true
	payload.Status.Code = "OK"
	payload.Status.Message = util.GetI18nErrorMessage(ctxModel.Locale, action, nil)
}
