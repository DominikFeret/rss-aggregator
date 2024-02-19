package main

import (
	"time"

	"github.com/DominikFeret/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"apiKey"`
}

func databaseUserToUser(u database.User) User {
	return User{
		ID:        u.ID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Name:      u.Name,
		ApiKey:    u.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseFeedToFeed(f database.Feed) Feed {
	return Feed{
		ID:        f.ID,
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
		Name:      f.Name,
		Url:       f.Url,
		UserID:    f.UserID,
	}
}

func databaseFeedsToFeeds(feeds []database.Feed) []Feed {
	feedsOut := make([]Feed, len(feeds))
	for i, f := range feeds {
		feedsOut[i] = databaseFeedToFeed(f)
	}
	return feedsOut
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func databaseFeedFollowToFeedFollow(ff database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        ff.ID,
		CreatedAt: ff.CreatedAt,
		UpdatedAt: ff.UpdatedAt,
		UserID:    ff.UserID,
		FeedID:    ff.FeedID,
	}
}

func databaseFeedFollowsToFeedFollows(ffs []database.FeedFollow) []FeedFollow {
	ffsOut := make([]FeedFollow, len(ffs))
	for i, ff := range ffs {
		ffsOut[i] = databaseFeedFollowToFeedFollow(ff)
	}
	return ffsOut
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description *string   `json:"description"` // pointer to handle nulls
	PublishedAt time.Time `json:"published_at"`
	Url         string    `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func databasePostToPost(p database.Post) Post {
	var description *string
	if p.Description.Valid {
		description = &p.Description.String
	}

	return Post{
		ID:          p.ID,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		Title:       p.Title,
		Description: description,
		PublishedAt: p.PublishedAt,
		Url:         p.Url,
		FeedID:      p.FeedID,
	}
}

func databasePostsToPosts(posts []database.Post) []Post {
	postsOut := make([]Post, len(posts))
	for i, p := range posts {
		postsOut[i] = databasePostToPost(p)
	}
	return postsOut
}
