package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/the-att-21/rssagg/internal/database"
)

func (apiCnf *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed, err := apiCnf.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		Name:      params.Name,
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Url:       params.URL,
		UsersID:   user.ID,
	})
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error creating feed: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseFeedtoFeed(feed))
}

func (apiCnf *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCnf.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Couldn't get feeds: %v", err))
		return
	}
	respondWithJSON(w, 200, databaseFeedstoFeeds(feeds))
}
