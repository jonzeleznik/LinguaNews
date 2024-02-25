package scraper

import "time"

type Post struct {
	Id          int
	Url         string
	Image_url   string
	Title       string
	Description string
	Content     string
	Source      string
	Date        time.Time
}
