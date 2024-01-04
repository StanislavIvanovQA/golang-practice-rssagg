package handlers

import (
	"StanislavIvanovQA/golang-practice-rssagg/internal/auth"
	"StanislavIvanovQA/golang-practice-rssagg/internal/database"
	"StanislavIvanovQA/golang-practice-rssagg/responses"
	"fmt"
	"net/http"
)

type authedHandler func(w http.ResponseWriter, r *http.Request, user database.User)

func (apiConfig *ApiConfig) MiddlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		apiKey, err := auth.GetApiKey(request.Header)
		if err != nil {
			responses.RespondWithError(writer, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apiConfig.DB.GetUserByAPIKey(request.Context(), apiKey)
		if err != nil {
			responses.RespondWithError(writer, 400, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}

		handler(writer, request, user)
	}
}
