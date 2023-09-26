package main

import (
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Echo instance
	e := echo.New()

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, Pitascho API!")
	})

	// GORM database setup
	dsn := "username:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Start the Echo server
	e.Logger.Fatal(e.Start(":1323"))
}
