package cmd

import (
	"it-backend/internal/helper"
	"it-backend/internal/middleware"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func ServeHTTP() {
	config, err := helper.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	dependencies := InitDependency(config)

	app := fiber.New()

	app.Use(cors.New())

	app.Get("/health", dependencies.HealthCheck)
	app.Post("/auth/login", dependencies.AuthController.Login)

	bareMetal := app.Group("/bare-metal")
	bareMetal.Post("/", middleware.MiddlewareValidateAuth, dependencies.BareMetalController.Create)
	bareMetal.Put("/:id", middleware.MiddlewareValidateAuth, dependencies.BareMetalController.Update)
	bareMetal.Get("/", middleware.MiddlewareValidateAuth, dependencies.BareMetalController.GetAll)
	bareMetal.Get("/:id", middleware.MiddlewareValidateAuth, dependencies.BareMetalController.GetOne)
	bareMetal.Get("/status-history/:node_id", middleware.MiddlewareValidateAuth, dependencies.BareMetalController.GetNodeStatusHistory)
	bareMetal.Delete("/:id", middleware.MiddlewareValidateAuth, dependencies.BareMetalController.Delete)

	virtualMachine := app.Group("/virtual-machine")
	virtualMachine.Post("/", middleware.MiddlewareValidateAuth, dependencies.VirtualMachineController.Create)
	virtualMachine.Put("/:id", middleware.MiddlewareValidateAuth, dependencies.VirtualMachineController.Update)
	virtualMachine.Put("/config/:vm_id", middleware.MiddlewareValidateAuth, dependencies.VirtualMachineController.UpdateConfig)
	virtualMachine.Get("/", middleware.MiddlewareValidateAuth, dependencies.VirtualMachineController.GetAll)
	virtualMachine.Get("/:id", middleware.MiddlewareValidateAuth, dependencies.VirtualMachineController.GetOne)
	virtualMachine.Get("/status-history/:vm_id", middleware.MiddlewareValidateAuth, dependencies.VirtualMachineController.GetNodeStatusHistory)
	virtualMachine.Delete("/:id", middleware.MiddlewareValidateAuth, dependencies.VirtualMachineController.Delete)

	port := config.APP_URL
	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}
