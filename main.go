package main

import (
	"energia/config"
	"energia/middleware"
	"energia/routes"
	"log"

	authController "energia/controller/auth"
	authRepo "energia/repository/auth"
	authService "energia/service/auth"

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

	routeController := routes.RouteController{
		AuthRoutes: &routes.AuthRoutes{AuthController: authController},
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
