package routes

import (
	"pitascho-api/src/handlers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, userHandler *handlers.UserHandler) {
	// e.GET("/healthcheck", healthCheckHandler.HealthCheck)
	// Auth
	// [TODO]/api/userは「/api/signup」として切り分けたい
	// e.POST("/api/login", authHandler.Login)
	e.POST("/api/user", userHandler.CreateUser)
	// e.POST("/api/confirm-account/:token", userHandler.ConfirmAccount)

	// var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	// // Configure JWT middleware
	// jwtMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningKey: []byte(jwtKey), // replace with your own secret
	// 	ErrorHandler: func(err error) error {
	// 		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	// 	},
	// })

	// // Group for routes that require authentication
	// authenticated := e.Group("")
	// authenticated.Use(jwtMiddleware)

	// Authenticated /me route
	// authenticated.GET("/api/me", authHandler.GetMe)

}
