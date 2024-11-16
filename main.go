package main

import (
	"energia/config"
	"energia/helper/openaiadapter"
	"energia/middleware"
	"energia/routes"
	"log"
	"os"

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

	weatherController "energia/controller/weather"
	weatherRepo "energia/repository/weather"
	weatherService "energia/service/weather"

	suggestionController "energia/controller/suggestion"
	suggestionService "energia/service/suggestion"

	emailController "energia/controller/email"
	emailService "energia/service/email"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	openai "github.com/sashabaranov/go-openai"
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
	deviceUsageController := deviceUsageController.NewDeviceUsageController(deviceUsageService, deviceService)

	userUsageRepo := userUsageRepo.NewUserUsageRepo(db)
	userUsageService := userUsageService.NewUserUsageService(userUsageRepo, deviceUsageRepo)
	userUsageController := userUsageController.NewUserUsageController(userUsageService)

	weatherRepo := weatherRepo.NewWeatherRepo(db)
	weatherService := weatherService.NewWeatherService(weatherRepo)
	weatherController := weatherController.NewWeatherController(weatherService)

	openAIKey := os.Getenv("OPENAI_API_KEY")
	openaiClient := openai.NewClient(openAIKey)

	openaiAdapter := openaiadapter.NewOpenAIClientAdapter(openaiClient)

	suggestionService := suggestionService.NewSuggestionService(deviceRepo, weatherRepo, openaiAdapter)
	suggestionController := suggestionController.NewSuggestionController(suggestionService)

	emailService := emailService.NewEmailService(deviceUsageRepo, deviceService)
	emailController := emailController.NewEmailController(emailService)

	routeController := routes.RouteController{
		AuthRoutes:        &routes.AuthRoutes{AuthController: authController},
		DeviceRoutes:      &routes.DeviceRoutes{DeviceController: deviceController},
		DeviceUsageRoutes: &routes.DeviceUsageRoutes{DeviceUsageController: deviceUsageController},
		UserUsageRoutes:   &routes.UserUsageRoutes{UserUsageController: userUsageController},
		WeatherRoutes:     &routes.WeatherRoutes{WeatherController: weatherController},
		SuggestionRoutes:  &routes.SuggestionRoutes{SuggestionController: suggestionController},
		EmailRoutes:       &routes.EmailRoutes{EmailController: emailController},
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
