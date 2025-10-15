package middleware

import (
	"it-backend/internal/helper"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func MiddlewareValidateAuth(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
		log.Println("authorization empty")
		return ctx.Status(fiber.StatusUnauthorized).JSON(helper.Response{Message: "unauthorized"})
	}

	const bearerPrefix = "Bearer "
	if len(authHeader) < len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
		log.Println("invalid authorization format")
		return ctx.Status(fiber.StatusUnauthorized).JSON(helper.Response{Message: "unauthorized"})
	}

	accessToken := authHeader[len(bearerPrefix):]

	// user, err := d.UserRepository.GetUserSessionByAccessToken(ctx.Context(), accessToken)
	// if err != nil {
	// 	log.Println("failed to get user session on DB: ", err)
	// 	return ctx.Status(fiber.StatusUnauthorized).JSON(helper.Response{Message: "unauthorized"})
	// }

	claim, err := helper.ValidateToken(ctx.Context(), accessToken)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusUnauthorized).JSON(helper.Response{Message: "unauthorized"})
	}

	if time.Now().Unix() > claim.ExpiresAt.Unix() {
		log.Println("jwt token is expired: ", claim.ExpiresAt)
		return ctx.Status(fiber.StatusUnauthorized).JSON(helper.Response{Message: "unauthorized"})
	}

	ctx.Locals("accessToken", accessToken)
	ctx.Locals("userId", claim.UserID)

	return ctx.Next()
}
