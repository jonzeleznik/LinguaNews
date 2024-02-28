package handler

import (
	"strconv"

	"github.com/jonzeleznik/LinguaNews/internal/db"
	"github.com/jonzeleznik/LinguaNews/internal/translate"
	"github.com/jonzeleznik/LinguaNews/internal/view/components"
	"github.com/jonzeleznik/LinguaNews/internal/view/pages"

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
	if err != nil {
		return render(c, components.ErrorCard("ERROR accured when translating!", err.Error()))
	}

	post, err := h.DB.GetPostById(id)
	if err != nil {
		return render(c, components.ErrorCard("ERROR accured when translating!", err.Error()))
	}

	var translated translate.Respone
	if post.Translated == "" {
		translated, err = translate.ChatGpt(post.Title, post.Description, post.Content)
		if err != nil {
			return render(c, components.ErrorCard("ERROR accured when translating!", err.Error()))
		}

		post.Translated = translated.Choices[0].Message.Content
		_, err = h.DB.UpdatePost(post)
		if err != nil {
			return render(c, components.ErrorCard("ERROR accured when translating!", err.Error()))
		}
	}

	return render(c, components.TextArea(post.Title, post.Translated))
}
