package main

import (
	"fmt"
	"log"
	"web-scrape/internal/db"
)

func main() {
	storage, err := db.NewPostStorage()
	if err != nil {
		log.Fatal(err)
	}

	// posts := scraper.HwrScrapeMoveiPosts()

	// for _, p := range posts {
	// 	storage.InsertPost(p)
	// }

	post, err := storage.GetPost(3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(post)
}
