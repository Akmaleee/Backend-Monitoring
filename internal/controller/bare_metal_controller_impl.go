package controller

import (
	"it-backend/internal/helper"
	"it-backend/internal/model/dto"
	"it-backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type BareMetalControllerImpl struct {
	BareMetalService service.BareMetalService
	AuthService      service.AuthService
}

func NewBareMetalController(bareMetalService service.BareMetalService, authService service.AuthService) *BareMetalControllerImpl {
	return &BareMetalControllerImpl{
		BareMetalService: bareMetalService,
		AuthService:      authService,
	}
}

func (c *BareMetalControllerImpl) GetAll(ctx *fiber.Ctx) error {
	log := helper.Logger

	bareMetals, err := c.BareMetalService.GetAll(ctx.Context())
	if err != nil {
		log.Error("Failed to get all bare metal: ", err)
		return helper.SendResponse(ctx, http.StatusInternalServerError, false, "Internal Server Error", nil, err.Error())
	}

	return helper.SendResponse(ctx, http.StatusOK, true, "Get All Bare Metal", bareMetals, nil)
}

func (c *BareMetalControllerImpl) GetOne(ctx *fiber.Ctx) error {
	log := helper.Logger

	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return err
	}

	bareMetal, err := c.BareMetalService.GetOne(ctx.Context(), uint64(id))
	if err != nil {
		log.Error("Failed to get all bare metal: ", err)
		return helper.SendResponse(ctx, http.StatusInternalServerError, false, "Internal Server Error", nil, err.Error())
	}

	return helper.SendResponse(ctx, http.StatusOK, true, "Get Bare Metal", bareMetal, nil)
}

func (c *BareMetalControllerImpl) GetNodeStatusHistory(ctx *fiber.Ctx) error {
	log := helper.Logger

	id, err := strconv.Atoi(ctx.Params("node_id"))

	if err != nil {
		return err
	}

	BareMetalNodeStatusHistory, err := c.BareMetalService.GetNodeStatusHistory(ctx.Context(), uint64(id))
	if err != nil {
		log.Error("Failed to get bare metal node status history: ", err)
		return helper.SendResponse(ctx, http.StatusInternalServerError, false, "Internal Server Error", nil, err.Error())
	}

	return helper.SendResponse(ctx, http.StatusOK, true, "Get Bare Metal Node Status History", BareMetalNodeStatusHistory, nil)
}

func (c *BareMetalControllerImpl) Create(ctx *fiber.Ctx) error {
	log := helper.Logger

	var req dto.BareMetalRequest
	if err := ctx.BodyParser(&req); err != nil {
		log.Error("Failed to parse request body: ", err)
		return helper.SendResponse(ctx, http.StatusBadRequest, false, "Bad Request", nil, err.Error())
	}

	bareMetal, err := c.BareMetalService.Create(ctx.Context(), req)
	if err != nil {
		log.Error("Failed to create bare metal: ", err)
		return helper.SendResponse(ctx, http.StatusInternalServerError, false, "Internal Server Error", nil, err.Error())
	}

	return helper.SendResponse(ctx, http.StatusOK, true, "Create Bare Metal", bareMetal, nil)
}

func (c *BareMetalControllerImpl) Update(ctx *fiber.Ctx) error {
	log := helper.Logger

	var req dto.BareMetalRequest
	if err := ctx.BodyParser(&req); err != nil {
		log.Error("Failed to parse request body: ", err)
		return helper.SendResponse(ctx, http.StatusBadRequest, false, "Bad Request", nil, err.Error())
	}

	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return err
	}

	bareMetal, err := c.BareMetalService.Update(ctx.Context(), req, uint64(id))
	if err != nil {
		log.Error("Failed to update bare metal: ", err)
		return helper.SendResponse(ctx, http.StatusInternalServerError, false, "Internal Server Error", nil, err.Error())
	}

	return helper.SendResponse(ctx, http.StatusOK, true, "Update Bare Metal", bareMetal, nil)
}

func (c *BareMetalControllerImpl) Delete(ctx *fiber.Ctx) error {
	log := helper.Logger

	var req dto.LoginRequest

	username, err := helper.GetUsernameFromPayload(ctx)
	if err != nil {
		return helper.SendResponse(ctx, http.StatusUnauthorized, false, "Unauthorized", nil, err.Error())
	}

	req.Username = username

	if err := ctx.BodyParser(&req); err != nil {
		log.Error("Failed to get password: ", err)
		return helper.SendResponse(ctx, http.StatusBadRequest, false, "Invalid Request Body", nil, err.Error())
	}

	err = c.AuthService.CheckLDAP(ctx.Context(), req)
	if err != nil {
		log.Error("Failed to check LDAP: ", err)
		return helper.SendResponse(ctx, http.StatusBadRequest, false, "Invalid password", nil, err.Error())
	}

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return helper.SendResponse(ctx, http.StatusBadRequest, false, "Invalid Request Id", nil, err.Error())
	}

	err = c.BareMetalService.Delete(ctx.Context(), uint64(id))
	if err != nil {
		log.Error("Failed to delete bare metal: ", err)
		return helper.SendResponse(ctx, http.StatusInternalServerError, false, "Internal Server Error", nil, err.Error())
	}

	return helper.SendResponse(ctx, http.StatusOK, true, "Delete Bare Metal", nil, nil)
}
