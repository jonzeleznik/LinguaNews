package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"web-scrape/internal/scraper"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

var fileDB = "posts.db"

const create string = `
	CREATE TABLE IF NOT EXISTS posts (
		id INTEGER NOT NULL PRIMARY KEY, 
		title TEXT, 
		image_url TEXT,
		url TEXT,
		description TEXT,
		content TEXT,
		source TEXT,
		date DATETIME
  );`

type PostStorage struct {
	db *sql.DB
}

func NewPostStorage() (*PostStorage, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file could't be loaded! " + err.Error())
	}
	fileDB = os.Getenv("DB_PATH")

	db, err := sql.Open("sqlite3", fileDB)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(create); err != nil {
		return nil, err
	}

	return &PostStorage{
		db: db,
	}, nil
}

func (s *PostStorage) InsertPost(post scraper.Post) (int, error) {
	res, err := s.db.Exec(
		"INSERT INTO posts VALUES(NULL,?,?,?,?,?,?,?);",
		post.Title,
		post.Image_url,
		post.Url,
		post.Description,
		post.Content,
		post.Source,
		post.Date,
	)
	if err != nil {
		return 0, err
	}

	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return 0, err
	}
	return int(id), nil
}

func (s *PostStorage) GetPostById(id int) (scraper.Post, error) {
	row := s.db.QueryRow("SELECT * FROM posts WHERE id=?", id)

	post := scraper.Post{}
	err := row.Scan(
		&post.Id,
		&post.Title,
		&post.Image_url,
		&post.Url,
		&post.Description,
		&post.Content,
		&post.Source,
		&post.Date,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return scraper.Post{}, fmt.Errorf("post with ID %d not found", id)
		}
		return scraper.Post{}, err
	}
	return post, nil
}

func (s *PostStorage) GetPostByTitle(title string) ([]scraper.Post, error) {
	rows, err := s.db.Query("SELECT * FROM posts WHERE title = ?", title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := []scraper.Post{}
	for rows.Next() {
		i := scraper.Post{}
		err = rows.Scan(
			&i.Id,
			&i.Title,
			&i.Image_url,
			&i.Url,
			&i.Description,
			&i.Content,
			&i.Source,
			&i.Date,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, i)
	}
	return data, nil
}
func (s *PostStorage) CustomSelect(custom string) ([]scraper.Post, error) {
	// MUST BE 'SELECT * ...'
	rows, err := s.db.Query(custom)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := []scraper.Post{}
	for rows.Next() {
		i := scraper.Post{}
		err = rows.Scan(
			&i.Id,
			&i.Title,
			&i.Image_url,
			&i.Url,
			&i.Description,
			&i.Content,
			&i.Source,
			&i.Date,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, i)
	}
	return data, nil
}

func (s *PostStorage) DelPost(id int64) (int, error) {
	res, err := s.db.Exec("delete from sessions where id = ?;", id)

	if err != nil {
		return 0, err
	}

	if id, err = res.LastInsertId(); err != nil {
		return 0, err
	}
	return int(id), nil
}
