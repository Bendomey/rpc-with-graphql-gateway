package main

import (
	"github/Bendomey/peerstronix-store/gateway/graph"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv, err := graph.NewGraphqlServer("localhost:5000")
	if err != nil {
		log.Fatalf("An error occured %s", err)
	}

	log.Printf("connect to http://localhost:%s/grapql for GraphQL playground", port)
	http.Handle("/", handler.GraphQL(srv.ToExecutableSchema()))
	http.Handle("/graphql", handler.Playground("Peerstronix", "/"))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
