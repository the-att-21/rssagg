package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/the-att-21/rssagg/internal/database"
)

func (apiCnf *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCnf.DB.CreateUser(r.Context(), database.CreateUserParams{
		Name:      params.Name,
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error creating user: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseUsertoUser(user))
}

func (apiCnf *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, 200, databaseUsertoUser(user))
}

func (apiCnf *apiConfig) handlerGetPostForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCnf.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UsersID: user.ID,
		Limit:  10,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error getting posts: %v", err))
		return
	}

	respondWithJSON(w, 200, databasePoststoPosts(posts))
}
