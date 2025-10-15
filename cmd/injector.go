//go:build wireinject
// +build wireinject

package cmd

import (
	"it-backend/database"
	"it-backend/internal/controller"
	"it-backend/internal/helper"
	"it-backend/internal/repository"
	"it-backend/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

type Dependency struct {
	Database database.DatabaseMySQL

	AuthController           controller.AuthController
	BareMetalController      controller.BareMetalController
	VirtualMachineController controller.VirtualMachineController
}

func (d *Dependency) HealthCheck(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"message": "OK",
	})
}

func InitDependency(config *helper.Config) *Dependency {
	wire.Build(
		database.NewDatabaseMySQL,

		repository.NewAuthRepository,
		wire.Bind(new(repository.AuthRepository), new(*repository.AuthRepositoryImpl)),
		repository.NewBareMetalRepository,
		wire.Bind(new(repository.BareMetalRepository), new(*repository.BareMetalRepositoryImpl)),
		repository.NewVirtualMachineRepository,
		wire.Bind(new(repository.VirtualMachineRepository), new(*repository.VirtualMachineRepositoryImpl)),

		service.NewAuthService,
		wire.Bind(new(service.AuthService), new(*service.AuthServiceImpl)),
		service.NewBareMetalService,
		wire.Bind(new(service.BareMetalService), new(*service.BareMetalServiceImpl)),
		service.NewVirtualMachineService,
		wire.Bind(new(service.VirtualMachineService), new(*service.VirtualMachineServiceImpl)),

		controller.NewAuthController,
		wire.Bind(new(controller.AuthController), new(*controller.AuthControllerImpl)),
		controller.NewBareMetalController,
		wire.Bind(new(controller.BareMetal), new(*controller.BareMetalControllerImpl)),
		controller.NewVirtualMachineController,
		wire.Bind(new(controller.VirtualMachineController), new(*controller.VirtualMachineControllerImpl)),

		wire.Struct(new(Dependency), "*"),
	)
	return &Dependency{}
}
