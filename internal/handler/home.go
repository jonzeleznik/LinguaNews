package handler

import (
	"web-scrape/internal/scraper"
	"web-scrape/internal/view/components"
	"web-scrape/internal/view/pages"

	"github.com/labstack/echo/v4"
)

type HomeHandler struct{}

func (h HomeHandler) HandleHomeShow(c echo.Context) error {
	posts := scraper.HwrScrapeMoveiPosts()
	return render(c, pages.Home(posts))
}

func (h HomeHandler) HandleButtonClick(c echo.Context) error {
	text := c.FormValue("text")
	return render(c, components.TextArea(text))
}
