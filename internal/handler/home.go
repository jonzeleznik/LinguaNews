package handler

import (
	"web-scrape/internal/view/pages"

	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	return render(c, pages.Home())
}
