package models

import (
	"time"
	"github.com/markorm/tests/dirtyroute_test/db"
	"strconv"
	"errors"
)

type PostModel struct {
	Id			int
	Title		string
	AuthorId	int
	Created		time.Time
	Modified	time.Time
	Tags		string
	Description string
	Body		string
}

// Single post by id
func GetPost(id string) (PostModel, error) {
	var err error
	post := PostModel{}
	conn := db.GetConnection()
	postId, err := strconv.Atoi(id)
	if err != nil { return post, err }
	rows, err := conn.Query("SELECT * FROM posts WHERE id = ? ", postId)
	if err != nil { return post, err }
	for rows.Next() {
		err = rows.Scan (
			&post.Id,
			&post.Title,
			&post.AuthorId,
			&post.Created,
			&post.Modified,
			&post.Tags,
			&post.Description,
			&post.Body,
		)
	}
	err = rows.Err()
	if post.Id == 0 {err = errors.New("Could not find a post with that Id")}
	return post, err
}

// Get all posts
func GetPosts() ([]PostModel, error) {
	var err error
	posts := []PostModel{}
	conn := db.GetConnection()
	rows, err := conn.Query("SELECT * FROM posts")
	if err != nil { return posts, err }
	for rows.Next() {
		post := PostModel{}
		err = rows.Scan (
			&post.Id,
			&post.Title,
			&post.AuthorId,
			&post.Created,
			&post.Modified,
			&post.Tags,
			&post.Description,
			&post.Body,
		)
		posts = append(posts, post)
	}
	err = rows.Err()
	return posts, err
}