package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/the-att-21/rssagg/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func databaseUsertoUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		ApiKey:    dbUser.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	User_ID   uuid.UUID `json:"user_id"`
}

func databaseFeedtoFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		Name:      dbFeed.Name,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Url:       dbFeed.Url,
		User_ID:   dbFeed.UsersID,
	}
}

func databaseFeedstoFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}

	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, databaseFeedtoFeed(dbFeed))
	}
	return feeds
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User_ID   uuid.UUID `json:"user_id"`
	Feed_ID   uuid.UUID `json:"feed_id"`
}

func databaseFeedFollowtoFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeedFollow.FeedID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		User_ID:   dbFeedFollow.UsersID,
		Feed_ID:   dbFeedFollow.FeedID,
	}
}

func databaseFeedsFollowtoFeedsFollow(dbFeedsFollow []database.FeedFollow) []FeedFollow {
	feedsfollow := []FeedFollow{}

	for _, dbFeedFollow := range dbFeedsFollow {
		feedsfollow = append(feedsfollow, databaseFeedFollowtoFeedFollow(dbFeedFollow))
	}
	return feedsfollow
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Url         string    `json:"url"`
	Description *string   `json:"description"`
	Feed_ID     uuid.UUID `json:"feed_id"`
	PublishedAt time.Time `json:"published_at"`
}

func databasePosttoPost(dbPost database.Post) Post {
	var desc *string
	if dbPost.Description.Valid {
		desc = &dbPost.Description.String
	}

	return Post{
		ID:          dbPost.ID,
		Title:       dbPost.Title,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   dbPost.UpdatedAt,
		Url:         dbPost.Url,
		Description: desc,
		Feed_ID:     dbPost.FeedID,
		PublishedAt: dbPost.PublishedAt,
	}
}

func databasePoststoPosts(dbPosts []database.Post) []Post {
	posts := []Post{}

	for _, dbPost := range dbPosts {
		posts = append(posts, databasePosttoPost(dbPost))
	}
	return posts
}