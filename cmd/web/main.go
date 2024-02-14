package main

import (
	"web-scrape/internal/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	app.Static("/dist", "internal/assets/dist")

	homeHandler := handler.HomeHandler{}
	app.GET("/home", homeHandler.HandleHomeShow)
	app.POST("/get-info", homeHandler.HandleButtonClick)

	app.Start(":3000")
}
