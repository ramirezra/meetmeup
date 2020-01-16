package main

import (
	"github.com/go-pg/pg/v9"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"

	"github.com/ramirezra/meetmeup"
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

	c := meetmeup.Config{Resolvers: &meetmeup.Resolver{
		MeetupsRepo: postgres.MeetupsRepo{DB: DB},
		UsersRepo:   postgres.UsersRepo{DB: DB},
	}}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(meetmeup.NewExecutableSchema(c)))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
