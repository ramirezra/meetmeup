package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-pg/pg/v9"

	"github.com/99designs/gqlgen/handler"

	"github.com/ramirezra/meetmeup/graphql"
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

	c := graphql.Config{Resolvers: &graphql.Resolver{
		MeetupsRepo: postgres.MeetupsRepo{DB: DB},
		UsersRepo:   postgres.UsersRepo{DB: DB},
	}}

	queryHandler := handler.GraphQL(graphql.NewExecutableSchema(c))

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", graphql.DataloaderMiddleware(DB, queryHandler))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
