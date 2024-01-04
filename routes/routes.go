package routes

import (
	"StanislavIvanovQA/golang-practice-rssagg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func CreateRouter(apiConfig *handlers.ApiConfig) chi.Router {
	router := chi.NewRouter()

	router.Use(cors.Handler(
		cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"*"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300,
		},
	))

	v1router := chi.NewRouter()
	v1router.Get("/healthcheck", handlers.HandlerReadiness)
	v1router.Get("/err", handlers.HandlerError)

	v1router.Post("/users", apiConfig.HandlerCreateUser)
	v1router.Get("/users", apiConfig.MiddlewareAuth(apiConfig.HandlerGetUser))

	v1router.Get("/posts", apiConfig.MiddlewareAuth(apiConfig.HandlerGetPostsForUser))

	v1router.Get("/feeds", apiConfig.HandlerGetFeeds)
	v1router.Post("/feeds", apiConfig.MiddlewareAuth(apiConfig.HandlerCreateFeed))

	v1router.Get("/feed_follows", apiConfig.MiddlewareAuth(apiConfig.HandlerGetFeedFollows))
	v1router.Post("/feed_follows", apiConfig.MiddlewareAuth(apiConfig.HandlerCreateFeedFollow))
	v1router.Delete("/feed_follows/{id}", apiConfig.MiddlewareAuth(apiConfig.HandlerDeleteFeedFollow))

	router.Mount("/v1", v1router)

	return router
}
