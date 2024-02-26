package scraper

import (
	"fmt"
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

			// To bypass "Too Many Requests" ERROR
			time.Sleep(2 * time.Second)
			p := HwrScrapePostContent(post.Url)

			post.Title = p.Title
			post.Image_url = p.Image_url
			post.Description = p.Description
			post.Content = p.Content
			post.Source = p.Source
			post.Date = p.Date

			fmt.Println(post)
			hwrMoveiPosts = append(hwrMoveiPosts, post)
			i++
		}
	})

	c.Visit("https://www.hollywoodreporter.com/c/movies/movie-news/")

	return hwrMoveiPosts
}

func HwrScrapePostContent(target string) Post {
	var content []string
	var post Post

	c := colly.NewCollector()

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong when parsing HWR post content: ", err)
	})

	c.OnHTML("p.paragraph", func(e *colly.HTMLElement) {

		paragraph := e.Text

		content = append(content, paragraph)
	})

	c.OnHTML("p.article-excerpt", func(e *colly.HTMLElement) {
		post.Description = e.Text
	})

	c.OnHTML("div.featured-image", func(e *colly.HTMLElement) {
		post.Image_url = e.ChildAttr("img", "src")
	})

	c.OnHTML("h1.article-title", func(e *colly.HTMLElement) {
		post.Title = e.Text
	})

	c.Visit(target)

	post.Source = "HWR"
	post.Date = time.Now()
	post.Content = strings.Replace(strings.Join(content[:], ""), "\n", "", -1)
	return post
}
