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

func (apiConfig *ApiConfig) HandlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responses.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}

	feed, err := apiConfig.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		responses.RespondWithError(w, 400, fmt.Sprintf("Cant create feed: %s", err))
		return
	}

	responses.RespondWithJSON(w, 201, models.DatabaseFeedToFeed(feed))
}
func (apiConfig *ApiConfig) HandlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiConfig.DB.GetFeeds(r.Context())
	if err != nil {
		responses.RespondWithError(w, 400, fmt.Sprintf("Cant get feeds: %s", err))
		return
	}

	responses.RespondWithJSON(w, 201, models.DatabaseFeedsToFeeds(feeds))
}
