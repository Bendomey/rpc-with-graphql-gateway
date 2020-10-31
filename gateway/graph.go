package gateway

import (
	"github/Bendomey/peerstronix-store/user"

	"github.com/99designs/gqlgen/graphql"
)

//Server get the services
type Server struct {
	userClient *user.Client
}

// NewGraphqlServer creates a graphql server of all microservices
func NewGraphqlServer(userURL string) (*Server, error) {
	// connect to user service
	userClient, err := user.NewClient(userURL)
	if err != nil {
		return nil, err
	}

	return &Server{
		userClient: userClient,
	}, nil
}

//ToExecutableSchema to make executable schema
func (s *Server) ToExecutableSchema() graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: s,
	})
}
