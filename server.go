package main

import (
	"log"
	"net/http"
	"os"
	"todogql/db"
	"todogql/graph/resolvers"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db.Init()

	srv := handler.NewDefaultServer(resolvers.NewSchema(db.Client))

	srv.Use(entgql.Transactioner{TxOpener: db.Client})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
