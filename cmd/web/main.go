package main

import (
	"log"

	"github.com/jonzeleznik/LinguaNews/internal/db"
	"github.com/jonzeleznik/LinguaNews/internal/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	app.Static("/dist", "internal/assets/dist")

	storage, err := db.NewPostStorage()
	if err != nil {
		log.Fatal(err)
	}

	homeHandler := handler.HomeHandler{
		DB: *storage,
	}
	app.GET("/home", homeHandler.HandleHomeShow)
	app.GET("/get-info", homeHandler.HandleButtonClick)

	app.Start(":3000")
}
