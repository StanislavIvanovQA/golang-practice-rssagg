package main

import (
	"StanislavIvanovQA/golang-practice-rssagg/handlers"
	"StanislavIvanovQA/golang-practice-rssagg/internal/database"
	"StanislavIvanovQA/golang-practice-rssagg/routes"
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT should be specified in environment variables")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL should be specified in environment variables")
	}

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Can't connect to database:", err)
	}

	db := database.New(conn)
	apiConfig := handlers.ApiConfig{
		DB: db,
	}

	go startScraping(db, 10, time.Minute)

	router := routes.CreateRouter(&apiConfig)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Server is failed to start", err)
	}
}
