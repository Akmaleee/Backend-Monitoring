package controller

import (
	"it-backend/internal/helper"
	"it-backend/internal/model/dto"
	"it-backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type VirtualMachineControllerImpl struct {
	VirtualMachineService service.VirtualMachineService
	AuthService           service.AuthService
}

func NewVirtualMachineController(virtualMachineService service.VirtualMachineService, authService service.AuthService) *VirtualMachineControllerImpl {
	return &VirtualMachineControllerImpl{
		VirtualMachineService: virtualMachineService,
		AuthService:           authService,
	}
}

func (c *VirtualMachineControllerImpl) GetAll(ctx *fiber.Ctx) error {
	log := helper.Logger

	virtualMachines, err := c.VirtualMachineService.GetAll(ctx.Context())
	if err != nil {
		log.Error("Failed to get all virtual machines: ", err)
		return helper.SendResponse(ctx, http.StatusInternalServerError, false, "Internal Server Error", nil, err.Error())
	}

	return helper.SendResponse(ctx, http.StatusOK, true, "Get All Virtual Machines", virtualMachines, nil)
}

func (c *VirtualMachineControllerImpl) GetOne(ctx *fiber.Ctx) error {
	log := helper.Logger

	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return err
	}

	virtualMachine, err := c.VirtualMachineService.GetOne(ctx.Context(), uint64(id))
	if err != nil {
		log.Error("Failed to get virtual machine: ", err)
		return helper.SendResponse(ctx, http.StatusInternalServerError, false, "Internal Server Error", nil, err.Error())
	}

	return helper.SendResponse(ctx, http.StatusOK, true, "Get Virtual Machine's Detail", virtualMachine, nil)
}

func (c *VirtualMachineControllerImpl) GetNodeStatusHistory(ctx *fiber.Ctx) error {
	log := helper.Logger

	id, err := strconv.Atoi(ctx.Params("vm_id"))

	if err != nil {
		return err
	}

	virtualMachineStatusHistory, err := c.VirtualMachineService.GetStatusHistory(ctx.Context(), uint64(id))
	if err != nil {
		log.Error("Failed to get virtual machine status history: ", err)
		return helper.SendResponse(ctx, http.StatusInternalServerError, false, "Internal Server Error", nil, err.Error())
	}

	return helper.SendResponse(ctx, http.StatusOK, true, "Get Virtual Machine Node Status History", virtualMachineStatusHistory, nil)
}

func (c *VirtualMachineControllerImpl) Create(ctx *fiber.Ctx) error {
	log := helper.Logger

	var req dto.VirtualMachineRequest
	if err := ctx.BodyParser(&req); err != nil {
		log.Error("Failed to parse request body: ", err)
		return helper.SendResponse(ctx, http.StatusBadRequest, false, "Bad Request", nil, err.Error())
	}

	virtualMachine, err := c.VirtualMachineService.Create(ctx.Context(), req)
	if err != nil {
		log.Error("Failed to create virtual machine: ", err)
		return helper.SendResponse(ctx, http.StatusInternalServerError, false, "Internal Server Error", nil, err.Error())
	}

	return helper.SendResponse(ctx, http.StatusOK, true, "Create Virtual Machine", virtualMachine, nil)
}

func (c *VirtualMachineControllerImpl) Update(ctx *fiber.Ctx) error {
	log := helper.Logger

	var req dto.VirtualMachineRequest
	if err := ctx.BodyParser(&req); err != nil {
		log.Error("Failed to parse request body: ", err)
		return helper.SendResponse(ctx, http.StatusBadRequest, false, "Bad Request", nil, err.Error())
	}

	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return err
	}

	virtualMachine, err := c.VirtualMachineService.Update(ctx.Context(), req, uint64(id))
	if err != nil {
		log.Error("Failed to update virtual machine: ", err)
		return helper.SendResponse(ctx, http.StatusInternalServerError, false, "Internal Server Error", nil, err.Error())
	}

	return helper.SendResponse(ctx, http.StatusOK, true, "Update Virtual Machine", virtualMachine, nil)
}

func (c *VirtualMachineControllerImpl) UpdateConfig(ctx *fiber.Ctx) error {
	log := helper.Logger

	var req dto.VirtualMachineConfigRequest
	if err := ctx.BodyParser(&req); err != nil {
		log.Error("Failed to parse request body: ", err)
		return helper.SendResponse(ctx, http.StatusBadRequest, false, "Bad Request", nil, err.Error())
	}

	id, err := strconv.Atoi(ctx.Params("vm_id"))

	if err != nil {
		return err
	}

	virtualMachineConfig, err := c.VirtualMachineService.UpdateConfig(ctx.Context(), req, uint64(id))
	if err != nil {
		log.Error("Failed to update virtual machine config: ", err)
		return helper.SendResponse(ctx, http.StatusInternalServerError, false, "Internal Server Error", nil, err.Error())
	}

	return helper.SendResponse(ctx, http.StatusOK, true, "Update Virtual Machine Config", virtualMachineConfig, nil)
}

func (c *VirtualMachineControllerImpl) Delete(ctx *fiber.Ctx) error {
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

	err = c.VirtualMachineService.Delete(ctx.Context(), uint64(id))
	if err != nil {
		log.Error("Failed to delete virtual machine: ", err)
		return helper.SendResponse(ctx, http.StatusInternalServerError, false, "Internal Server Error", nil, err.Error())
	}

	return helper.SendResponse(ctx, http.StatusOK, true, "Delete Virtual Machine", nil, nil)
}
