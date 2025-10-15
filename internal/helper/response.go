package helper

import "github.com/gofiber/fiber/v2"

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  interface{} `json:"errors"`
}

func SendResponse(ctx *fiber.Ctx, statusCode int, success bool, message string, data interface{}, errors interface{}) error {
	return ctx.Status(statusCode).JSON(Response{
		Success: success,
		Message: message,
		Data:    data,
		Errors:  errors,
	})
}

type PaginatedResult[T any] struct {
	Data       []T `json:"data"`
	Total      int `json:"total"`
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	TotalPages int `json:"total_pages"`
}
