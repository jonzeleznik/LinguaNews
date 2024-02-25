package scraper

import (
	"log"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func HwrScrapeMoveiPosts() []Post {
	var hwrMoveiPosts []Post
	i := 0

	c := colly.NewCollector()

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong when parsing HWR posts: ", err)
	})

	c.OnHTML("div.story", func(e *colly.HTMLElement) {
		if i <= 4 {
			post := Post{}

			post.Url = e.ChildAttr("a", "href")
			post.Title = e.ChildText("h3")

			// To bypass "Too Many Requests" ERROR
			time.Sleep(2 * time.Second)
			post.Content,
				post.Description,
				post.Image_url = HwrScrapePostContent(post.Url)
			post.Source = "HWR"
			post.Date = time.Now()

			hwrMoveiPosts = append(hwrMoveiPosts, post)
			i++
		}
	})

	c.Visit("https://www.hollywoodreporter.com/c/movies/movie-news/")

	return hwrMoveiPosts
}

func HwrScrapePostContent(target string) (string, string, string) {
	var content []string
	var description string
	var image string

	c := colly.NewCollector()

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong when parsing HWR post content: ", err)
	})

	c.OnHTML("p.paragraph", func(e *colly.HTMLElement) {

		paragraph := e.Text

		content = append(content, paragraph)
	})

	c.OnHTML("p.article-excerpt", func(e *colly.HTMLElement) {
		description = e.Text
	})

	c.OnHTML("div.featured-image", func(e *colly.HTMLElement) {
		image = e.ChildAttr("img", "src")
	})

	c.Visit(target)

	return strings.Replace(strings.Join(content[:], ""), "\n", "", -1),
		description,
		image
}
