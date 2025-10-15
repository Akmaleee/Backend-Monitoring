// package controller

// import (
// 	"it-backend/internal/helper"
// 	"it-backend/internal/model/dto"
// 	"it-backend/internal/service"
// 	"net/http"

// 	"github.com/gofiber/fiber/v2"
// )

// type AuthControllerImpl struct {
// 	AuthService service.AuthService
// }

// func NewAuthController(authService service.AuthService) *AuthControllerImpl {
// 	return &AuthControllerImpl{
// 		AuthService: authService,
// 	}
// }

// func (c *AuthControllerImpl) Login(ctx *fiber.Ctx) error {
// 	log := helper.Logger
// 	var req dto.LoginRequest

// 	if err := ctx.BodyParser(&req); err != nil {
// 		log.Error("Failed to parse request body: ", err)
// 		return helper.SendResponse(ctx, http.StatusBadRequest, false, "Bad Request", nil, err.Error())
// 	}

// 	if err := req.Validate(); err != nil {
// 		log.Error("Failed to validate request body: ", err)
// 		return helper.SendResponse(ctx, http.StatusBadRequest, false, "Username and password are required", nil, err.Error())
// 	}

// 	token, err := c.AuthService.LoginLDAP(ctx.Context(), req)
// 	if err != nil {
// 		log.Error("Failed to login LDAP: ", err)
// 		return helper.SendResponse(ctx, http.StatusUnauthorized, false, "Invalid username or password", nil, err.Error())
// 	}

// 	return helper.SendResponse(ctx, http.StatusOK, true, "Login Success (LDAP)", dto.LoginResponse{Token: token}, nil)
// }



package controller

import (
	"it-backend/internal/helper"
	"it-backend/internal/model/dto"
	"it-backend/internal/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthControllerImpl {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}

func (c *AuthControllerImpl) Login(ctx *fiber.Ctx) error {
	log := helper.Logger
	var req dto.LoginRequest

	if err := ctx.BodyParser(&req); err != nil {
		log.Error("Failed to parse request body: ", err)
		return helper.SendResponse(ctx, http.StatusBadRequest, false, "Bad Request", nil, err.Error())
	}

	if err := req.Validate(); err != nil {
		log.Error("Failed to validate request body: ", err)
		return helper.SendResponse(ctx, http.StatusBadRequest, false, "Username and password are required", nil, err.Error())
	}

	token, err := c.AuthService.Login(ctx.Context(), req)
	// token, err := c.AuthService.LoginLDAP(ctx.Context(), req)
	if err != nil {
		log.Error("Failed to login LDAP: ", err)
		return helper.SendResponse(ctx, http.StatusUnauthorized, false, "Invalid username or password", nil, err.Error())
	}

	return helper.SendResponse(ctx, http.StatusOK, true, "Login Success (LDAP)", dto.LoginResponse{Token: token}, nil)
}
