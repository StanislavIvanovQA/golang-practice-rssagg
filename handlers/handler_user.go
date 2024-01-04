package handlers

import (
	"StanislavIvanovQA/golang-practice-rssagg/internal/database"
	"StanislavIvanovQA/golang-practice-rssagg/models"
	"StanislavIvanovQA/golang-practice-rssagg/responses"

	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (apiConfig *ApiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responses.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}

	user, err := apiConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		responses.RespondWithError(w, 400, fmt.Sprintf("Cant create user: %s", err))
		return
	}

	responses.RespondWithJSON(w, 201, models.DatabaseUserToUser(user))
}

func (apiConfig *ApiConfig) HandlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	responses.RespondWithJSON(w, 200, models.DatabaseUserToUser(user))
}

func (apiConfig *ApiConfig) HandlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiConfig.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		responses.RespondWithError(w, 400, fmt.Sprintf("Couldn't get posts: %v", err))
		return
	}

	responses.RespondWithJSON(w, 200, models.DatabasePostsToPosts(posts))
}
