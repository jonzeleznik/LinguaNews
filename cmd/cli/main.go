package main

import (
	"fmt"
	"log"
	"os"
	"web-scrape/internal/db"
	"web-scrape/internal/scraper"
)

func main() {
	args := os.Args[1]

	switch args {
	case "checkPosts":
		fmt.Println("Checking posts")
		CheckNewPosts()
	case "delOldPosts":
		fmt.Println("Deleting old posts")
		DelOldPosts()
	case "checkBrokenPosts":
		fmt.Println("Checking broken posts")
		CheckBrokenPosts()
	case "test":
		fmt.Println("Testing")
		Test()
	case "help":
		fmt.Println("| checkPosts | delOldPosts |")
	}
}

func CheckNewPosts() {
	storage, err := db.NewPostStorage()
	if err != nil {
		log.Fatal(err)
	}

	posts := scraper.HwrScrapeMoveiPosts()

	for _, p := range posts {
		post, err := storage.GetPostByTitle(p.Title)
		if err != nil {
			log.Fatal(err)
		}

		if len(post) == 0 {
			var id int
			id, err = storage.InsertPost(p)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%d was inserted", id)
		}

	}
}

func DelOldPosts() {
	storage, err := db.NewPostStorage()
	if err != nil {
		log.Fatal(err)
	}

	posts, err := storage.CustomSelect("SELECT * FROM posts WHERE datetime(date) < datetime('now', '-5 days');")
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range posts {
		storage.DelPost(int64(p.Id))
		fmt.Printf("Deleted id %d", p.Id)
	}
}

func CheckBrokenPosts() {
	storage, err := db.NewPostStorage()
	if err != nil {
		log.Fatal(err)
	}

	posts, err := storage.CustomSelect("SELECT * FROM posts WHERE title = '';")
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range posts {
		post := scraper.HwrScrapePostContent(p.Url)

		fmt.Println(post)

		post.Id = p.Id
		row, err := storage.UpdatePost(post)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Updated id %d, rows affected %d", p.Id, row)
	}
}

func Test() {
	storage, err := db.NewPostStorage()
	if err != nil {
		log.Fatal(err)
	}

	p, err := storage.GetPostById(6)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(p.Translated)
}
