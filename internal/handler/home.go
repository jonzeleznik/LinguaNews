package handler

import (
	"strconv"
	"web-scrape/internal/db"
	"web-scrape/internal/scraper"
	"web-scrape/internal/translate"
	"web-scrape/internal/view/components"
	"web-scrape/internal/view/pages"

	"github.com/labstack/echo/v4"
)

type HomeHandler struct {
	DB    db.PostStorage
	Posts []scraper.Post
}

func (h HomeHandler) HandleHomeShow(c echo.Context) error {
	var err error
	h.Posts, err = h.DB.CustomSelect("SELECT * FROM posts ORDER BY id DESC LIMIT 10;")
	if err != nil {
		return render(c, components.ErrorCard("ERROR accured when loading post!", err.Error()))
	}

	return render(c, pages.Home(h.Posts))
}

func (h HomeHandler) HandleButtonClick(c echo.Context) error {
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		return render(c, components.ErrorCard("ERROR accured when translating!", err.Error()))
	}

	translated, err := translate.ChatGpt(h.Posts[id].Title, h.Posts[id].Description, h.Posts[id].Content)
	if err != nil {
		return render(c, components.ErrorCard("ERROR accured when translating!", err.Error()))
	}
	return render(c, components.TextArea(h.Posts[id].Title, translated.Choices[0].Message.Content))
}
