package models

import (
	"StanislavIvanovQA/golang-practice-rssagg/internal/database"
	"github.com/google/uuid"
	"time"
)

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func DatabaseFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeedFollow.ID,
		UserID:    dbFeedFollow.UserID,
		FeedID:    dbFeedFollow.FeedID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
	}
}

func DatabaseFeedFollowsToFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollow {
	var feedFollows []FeedFollow
	for _, dbFeedFollow := range dbFeedFollows {
		feedFollows = append(feedFollows, DatabaseFeedFollowToFeedFollow(dbFeedFollow))
	}
	return feedFollows
}
