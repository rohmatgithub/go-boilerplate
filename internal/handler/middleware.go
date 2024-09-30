package handler

import (
	"boilerplate/internal/constanta"
	"boilerplate/internal/dto"
	"context"
	"os"
	"strconv"

	"boilerplate/internal/common"
	"boilerplate/pkg/applog"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func middleware(c *fiber.Ctx) error {
	logModel := &common.LoggerModel{
		Pid: strconv.Itoa(os.Getpid()),
		// RequestID: c.Locals("requestid").(string),
		// Version:     config.ApplicationConfiguration.GetServerConfig().Version,
		// ByteIn: len(c.Body()),
		Path: c.BaseURL(),
	}

	logger := context.WithValue(c.Context(), constanta.ApplicationContextConstanta, logModel)
	adaptor.CopyContextToFiberContext(logger, c.Context())

	err := c.Next()
	if err != nil {
		return err
	}
	logModel = c.Context().Value(constanta.ApplicationContextConstanta).(*common.LoggerModel)
	logModel.Status = c.Response().StatusCode()
	logModel.Path = c.OriginalURL()

	l := applog.GetLogger()
	l.Info().Interface("pid", logModel.Pid).
		Interface("request_id", c.Locals("requestid")).
		Interface("version", "").
		Interface("status", logModel.Status).
		Interface("path", logModel.Path).
		Interface("method", c.Method()).
		Interface("byte_in", len(c.Body())).
		Interface("byte_out", len(c.Response().Body())).
		Msg(logModel.Message)
	return err
}

func NotFoundHandler(c *fiber.Ctx) error {
	// Customize the response for the 404 error
	return c.Status(fiber.StatusNotFound).JSON(
		dto.Payload{
			Status: dto.StatusPayload{
				Success: false,
				Code:    strconv.Itoa(fiber.StatusNotFound),
				Message: "Not Found",
			},
		},
	)
}

func customErrorHandler(c *fiber.Ctx, err error) {
	// Handle the error here
	if err != nil {
		applog.Error().Msg(err.Error())
	}

	// Return a custom error response
	c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": "Something went wrong",
	})
}
