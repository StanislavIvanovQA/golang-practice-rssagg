package models

import (
	"StanislavIvanovQA/golang-practice-rssagg/internal/database"
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	FeedId      uuid.UUID `json:"feed_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func DatabasePostToPost(dbPost database.Post) Post {
	var description *string
	if dbPost.Description.Valid {
		description = &dbPost.Description.String
	}

	return Post{
		ID:          dbPost.ID,
		Title:       dbPost.Title,
		Description: description,
		PublishedAt: dbPost.PublishedAt,
		FeedId:      dbPost.FeedID,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   dbPost.UpdatedAt,
	}
}
func DatabasePostsToPosts(dbPosts []database.Post) []Post {
	var posts []Post
	for _, dbPost := range dbPosts {
		posts = append(posts, DatabasePostToPost(dbPost))
	}
	return posts
}
