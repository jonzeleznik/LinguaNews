package scraper

import (
	"log"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func ShfScrapeMoveiPosts() []HWRPost {
	var hwrMoveiPosts []HWRPost
	i := 0

	c := colly.NewCollector()

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnHTML("div.article-description", func(e *colly.HTMLElement) {
		if i <= 3 {
			post := HWRPost{}

			post.Url = "https://www.slashfilm.com/" + e.ChildAttr("a", "href")
			post.Title = e.ChildText("h3")

			// To bypass "Too Many Requests" ERROR
			time.Sleep(1 * time.Second)
			post.Content = ShfScrapePostContent(post.Url)

			hwrMoveiPosts = append(hwrMoveiPosts, post)
			i++
		}
	})

	c.Visit("https://www.slashfilm.com/category/movies/")

	return hwrMoveiPosts
}

func ShfScrapePostContent(target string) string {
	var content []string

	c := colly.NewCollector()

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnHTML("div.columns-holder", func(e *colly.HTMLElement) {

		paragraph := e.ChildText("p")

		content = append(content, paragraph)
	})

	c.Visit(target)

	return strings.Replace(strings.Join(content[:], ""), "\n", "", -1)
}
