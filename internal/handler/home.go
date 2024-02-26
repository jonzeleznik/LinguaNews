package handler

import (
	"fmt"
	"strconv"
	"web-scrape/internal/db"
	"web-scrape/internal/translate"
	"web-scrape/internal/view/components"
	"web-scrape/internal/view/pages"

	"github.com/labstack/echo/v4"
)

type HomeHandler struct {
	DB db.PostStorage
}

func (h HomeHandler) HandleHomeShow(c echo.Context) error {
	posts, err := h.DB.CustomSelect("SELECT * FROM posts ORDER BY id DESC LIMIT 10;")
	if err != nil {
		return render(c, components.ErrorCard("ERROR accured when loading post!", err.Error()))
	}

	return render(c, pages.Home(posts))
}

func (h HomeHandler) HandleButtonClick(c echo.Context) error {
	// TODO: add translated text to DB
	id, err := strconv.Atoi(c.QueryParam("id"))
	fmt.Println(id)
	if err != nil {
		return render(c, components.ErrorCard("ERROR accured when translating!", err.Error()))
	}

	post, err := h.DB.GetPostById(id)
	if err != nil {
		return render(c, components.ErrorCard("ERROR accured when translating!", err.Error()))
	}

	translated, err := translate.ChatGpt(post.Title, post.Description, post.Content)
	if err != nil {
		return render(c, components.ErrorCard("ERROR accured when translating!", err.Error()))
	}

	return render(c, components.TextArea(post.Title, translated.Choices[0].Message.Content))
}
