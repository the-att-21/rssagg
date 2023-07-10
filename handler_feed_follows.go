package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/the-att-21/rssagg/internal/database"
)

func (apiCnf *apiConfig) handlerCreateFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feedFollow, err := apiCnf.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UsersID:   user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error creating feed_follow: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseFeedFollowtoFeedFollow(feedFollow))
}

func (apiCnf *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollow, err := apiCnf.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error Getting feed_follow: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseFeedsFollowtoFeedsFollow(feedFollow))
}

func (apiCnf *apiConfig) handlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDstr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDstr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't parse feed follow id: %v", err))
		return
	}

	err = apiCnf.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:      feedFollowID,
		UsersID: user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't delete feed follow: %v", err))
	}
	respondWithJSON(w, 200, struct{}{})
}
