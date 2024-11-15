package main

import (
	"energia/config"
	"energia/middleware"
	"energia/routes"
	"log"

	authController "energia/controller/auth"
	authRepo "energia/repository/auth"
	authService "energia/service/auth"

	deviceController "energia/controller/device"
	deviceRepo "energia/repository/device"
	deviceService "energia/service/device"

	deviceUsageController "energia/controller/device-usage"
	deviceUsageRepo "energia/repository/device-usage"
	deviceUsageService "energia/service/device-usage"

	userUsageController "energia/controller/user-usage"
	userUsageRepo "energia/repository/user-usage"
	userUsageService "energia/service/user-usage"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	loadEnv()

	db, _ := config.ConnectDatabase()
	config.MigrateDB(db)

	e := echo.New()

	authJwt := middleware.JwtLink{}

	authRepo := authRepo.NewAuthRepo(db)
	authService := authService.NewAuthService(authRepo, authJwt)
	authController := authController.NewAuthController(authService)

	deviceRepo := deviceRepo.NewDeviceRepo(db)
	deviceService := deviceService.NewDeviceService(deviceRepo)
	deviceController := deviceController.NewDeviceController(deviceService)

	deviceUsageRepo := deviceUsageRepo.NewDeviceUsageRepo(db)
	deviceUsageService := deviceUsageService.NewDeviceUsageService(deviceRepo, deviceUsageRepo)
	deviceUsageController := deviceUsageController.NewDeviceUsageController(deviceUsageService)

	userUsageRepo := userUsageRepo.NewUserUsageRepo(db)
	userUsageService := userUsageService.NewUserUsageService(userUsageRepo, deviceUsageRepo)
	userUsageController := userUsageController.NewUserUsageController(userUsageService)

	routeController := routes.RouteController{
		AuthRoutes:        &routes.AuthRoutes{AuthController: authController},
		DeviceRoutes:      &routes.DeviceRoutes{DeviceController: deviceController},
		DeviceUsageRoutes: &routes.DeviceUsageRoutes{DeviceUsageController: deviceUsageController},
		UserUsageRoutes:   &routes.UserUsageRoutes{UserUsageController: userUsageController},
	}
	routeController.InitRoute(e)

	e.Start(":8000")
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		panic("Failed loading .env file")
	}
}
