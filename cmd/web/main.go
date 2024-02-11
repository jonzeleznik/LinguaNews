package main

import (
	"web-scrape/internal/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	app.Static("/dist", "internal/assets/dist")

	userHandler := handler.UserHandler{}
	app.GET("/home", userHandler.HandleUserShow)

	app.Start(":3000")
}
