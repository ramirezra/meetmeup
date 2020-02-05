package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"

	"github.com/go-pg/pg/v9"

	"github.com/99designs/gqlgen/handler"

	"github.com/ramirezra/meetmeup/domain"
	"github.com/ramirezra/meetmeup/graphql"
	customMW "github.com/ramirezra/meetmeup/middleware"
	"github.com/ramirezra/meetmeup/postgres"
)

const defaultPort = "8080"

func main() {
	DB := postgres.New(&pg.Options{
		Addr:     "172.16.2.201:5432",
		User:     "postgres",
		Password: os.Getenv("PGPW"),
		Database: "meetmeup_dev",
	})

	defer DB.Close()

	DB.AddQueryHook(postgres.DBLogger{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	userRepo := postgres.UsersRepo{DB: DB}
	meetupRepo := postgres.MeetupsRepo{DB: DB}

	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8000", "http:/localhost:8080"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(customMW.AuthMiddleware(userRepo))

	d := domain.NewDomain(userRepo, meetupRepo)
	c := graphql.Config{Resolvers: &graphql.Resolver{
		Domain: d,
	}}

	queryHandler := handler.GraphQL(graphql.NewExecutableSchema(c))

	router.Handle("/", handler.Playground("GraphQL playground", "/query"))
	router.Handle("/query", graphql.DataloaderMiddleware(DB, queryHandler))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
