package gateway

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/handler"
)

func main() {
	s, err := NewGraphqlServer("localhost:5000")
	if err != nil {
		log.Fatalf("Something happened => %v", err)
	}

	http.Handle("/graphql", handler.GraphQL(s.ToExecutableSchema()))
	http.Handle("/playground", handler.Playground("Peerstronix", "/graphql"))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
