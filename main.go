package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/the-att-21/rssagg/internal/database"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	fmt.Println("Hello World")

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT not found in environment variables")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL not found in environment variables")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	apiCnf := apiConfig{
		DB: database.New(db),
	}

	go startScraping(apiCnf.DB, 10, 10)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/error", handlerErr)

	v1Router.Post("/users", apiCnf.handlerCreateUser)
	v1Router.Get("/users", apiCnf.middlewareAuth(apiCnf.handlerGetUser))

	v1Router.Post("/feeds", apiCnf.middlewareAuth(apiCnf.handlerCreateFeed))
	v1Router.Get("/feeds", apiCnf.handlerGetFeeds)

	v1Router.Post("/feed_follows", apiCnf.middlewareAuth(apiCnf.handlerCreateFeedFollows))
	v1Router.Get("/feed_follows", apiCnf.middlewareAuth(apiCnf.handlerGetFeedFollows))
	v1Router.Delete("/feed_follows/{feedFollowID}", apiCnf.middlewareAuth(apiCnf.handlerDeleteFeedFollows))

	v1Router.Get("/posts", apiCnf.middlewareAuth(apiCnf.handlerGetPostForUser))

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
