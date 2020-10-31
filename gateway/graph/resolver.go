package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"github/Bendomey/peerstronix-store/user"
)

//Resolver get the services
type Resolver struct {
	userClient *user.Client
}

// NewGraphqlServer creates a graphql server of all microservices
func NewGraphqlServer(userURL string) (*Resolver, error) {
	// connect to user service
	userClient, err := user.NewClient(userURL)
	if err != nil {
		return nil, err
	}

	return &Resolver{
		userClient: userClient,
	}, nil
}
