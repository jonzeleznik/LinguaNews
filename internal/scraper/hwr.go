package scraper

import (
	"log"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type HWRPost struct {
	Url     string
	Title   string
	Content string
}

func HwrScrapeMoveiPosts() []HWRPost {
	var hwrMoveiPosts []HWRPost
	i := 0

	c := colly.NewCollector()

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong when parsing HWR posts: ", err)
	})

	c.OnHTML("div.story", func(e *colly.HTMLElement) {
		if i <= 1 {
			post := HWRPost{}

			post.Url = e.ChildAttr("a", "href")
			post.Title = e.ChildText("h3")

			// To bypass "Too Many Requests" ERROR
			time.Sleep(3 * time.Second)
			post.Content = HwrScrapePostContent(post.Url)

			hwrMoveiPosts = append(hwrMoveiPosts, post)
			i++
		}
	})

	c.Visit("https://www.hollywoodreporter.com/c/movies/movie-news/")

	return hwrMoveiPosts
}

func HwrScrapePostContent(target string) string {
	var content []string

	c := colly.NewCollector()

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong when parsing HWR post content: ", err)
	})

	c.OnHTML("p.paragraph", func(e *colly.HTMLElement) {

		paragraph := e.Text

		content = append(content, paragraph)
	})

	c.Visit(target)

	return strings.Replace(strings.Join(content[:], ""), "\n", "", -1)
}
