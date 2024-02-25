package handler

import (
	"strconv"
	"web-scrape/internal/scraper"
	"web-scrape/internal/translate"
	"web-scrape/internal/view/components"
	"web-scrape/internal/view/pages"

	"github.com/labstack/echo/v4"
)

type HomeHandler struct{}

var posts []scraper.Post

func (h HomeHandler) HandleHomeShow(c echo.Context) error {
	posts = scraper.HwrScrapeMoveiPosts()
	return render(c, pages.Home(posts))
}

func (h HomeHandler) HandleButtonClick(c echo.Context) error {
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		return render(c, components.ErrorCard("ERROR accured when translating!", err.Error()))
	}

	text := posts[id]
	translated, err := translate.ChatGpt(text.Title, text.Description, text.Content)
	if err != nil {
		return render(c, components.ErrorCard("ERROR accured when translating!", err.Error()))
	}
	return render(c, components.TextArea(text.Title, translated.Choices[0].Message.Content))
}
