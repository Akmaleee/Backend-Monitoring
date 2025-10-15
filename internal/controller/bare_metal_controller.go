package controller

import "github.com/gofiber/fiber/v2"

type BareMetalController interface {
	GetAll(ctx *fiber.Ctx) error
	GetOne(ctx *fiber.Ctx) error
	GetNodeStatusHistory(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
