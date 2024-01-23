package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/notprQn/rssagg/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"userId"`
}

func dataBaseUserToUser(DbUser database.User) User {
	return User{
		ID:        DbUser.ID,
		CreatedAt: DbUser.CreatedAt,
		UpdatedAt: DbUser.UpdatedAt,
		Name:      DbUser.Name,
		APIKey:    DbUser.ApiKey,
	}
}

func dataBaseFeedtoFeed(DbFeed database.Feed) Feed {
	return Feed{
		ID:        DbFeed.ID,
		CreatedAt: DbFeed.CreatedAt,
		UpdatedAt: DbFeed.UpdatedAt,
		Name:      DbFeed.Name,
		Url:       DbFeed.Url,
		UserID:    DbFeed.UserID,
	}
}

func dataBaseFeedstoFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, dataBaseFeedtoFeed(dbFeed))
	}
	return feeds
}
