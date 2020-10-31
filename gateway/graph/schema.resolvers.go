package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github/Bendomey/peerstronix-store/gateway/graph/generated"
	"github/Bendomey/peerstronix-store/gateway/graph/model"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input *model.CreateUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteUser(ctx context.Context, input model.GetUserFilter) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUsers(ctx context.Context, pagination *model.PaginationInput) ([]*model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	//get users
	var skip, take uint64
	// skip, take = pagination.Skip, pagination.Take
	// if skip == nil {
	// 	skip = int(0)
	// }
	// if take == nil {
	// 	take = int(0)
	// }

	userLists, err := r.userClient.GetUsers(ctx, skip, take)
	if err != nil {
		return nil, err
	}

	var users []*model.User
	//loop and append
	for _, u := range userLists {
		users = append(users, &model.User{
			ID:    u.ID,
			Name:  u.Name,
			Phone: u.Phone,
			Email: u.Email,
		})
	}
	return users, nil

	// r.userClient.GetUsers()
}

func (r *queryResolver) GetUser(ctx context.Context, filter model.GetUserFilter) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{
		Resolver: r,
	}
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{
		Resolver: r,
	}
}

//ToExecutableSchema to make executable schema
func (r *Resolver) ToExecutableSchema() graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: r,
	})
}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
