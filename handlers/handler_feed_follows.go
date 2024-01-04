package handlers

import (
	"StanislavIvanovQA/golang-practice-rssagg/internal/database"
	"StanislavIvanovQA/golang-practice-rssagg/models"
	"StanislavIvanovQA/golang-practice-rssagg/responses"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (apiConfig *ApiConfig) HandlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responses.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}

	feedFollow, err := apiConfig.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		responses.RespondWithError(w, 400, fmt.Sprintf("Cant create feed follow: %s", err))
		return
	}

	responses.RespondWithJSON(w, 201, models.DatabaseFeedFollowToFeedFollow(feedFollow))
}

func (apiConfig *ApiConfig) HandlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apiConfig.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		responses.RespondWithError(w, 400, fmt.Sprintf("Cant get feed follows: %s", err))
	}
	responses.RespondWithJSON(w, 200, models.DatabaseFeedFollowsToFeedFollows(feedFollows))
}

func (apiConfig *ApiConfig) HandlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowId, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		responses.RespondWithError(w, 400, fmt.Sprintf("Cant delete feed follow: %s", err))
	}

	err = apiConfig.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowId,
		UserID: user.ID,
	})
	if err != nil {
		responses.RespondWithError(w, 400, fmt.Sprintf("Cant delete feed follow: %s", err))
	}
	responses.RespondWithJSON(w, 200, struct{}{})
}
