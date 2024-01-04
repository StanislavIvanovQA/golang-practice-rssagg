package handlers

import (
	"StanislavIvanovQA/golang-practice-rssagg/responses"
	"net/http"
)

func HandlerError(w http.ResponseWriter, r *http.Request) {
	responses.RespondWithError(w, 400, "Something went wrong")
}
