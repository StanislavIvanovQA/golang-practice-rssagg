package handlers

import (
	"StanislavIvanovQA/golang-practice-rssagg/responses"
	"net/http"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	responses.RespondWithJSON(w, 200, struct{}{})
}
