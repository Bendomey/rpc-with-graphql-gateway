package user

import (
	"context"

	"github.com/segmentio/ksuid"
)

// Service inteface holds the functions of this controller
type Service interface {
	CreateUser(ctx context.Context, name string, phone string, email string, password string) (*User, error)
	UpdateUser(ctx context.Context, name string, phone string, email string, id string) (*User, error)
	GetUsers(ctx context.Context, skip uint64, take uint64) ([]User, error)
	GetUser(ctx context.Context, id string) (*User, error)
	DeleteUser(ctx context.Context, id string) (bool, error)
}

// User model
type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsDeleted bool   `json:"isDeleted"`
	// createdAt Date
}

//UserRepository gets repository
type userRepository struct {
	repository Repository
}

// NewService exposed the repository to the functions in the module
func NewService(r Repository) Service {
	return &userRepository{r}
}

func (s *userRepository) CreateUser(ctx context.Context, name string, phone string, email string, password string) (*User, error) {
	u := &User{
		ID:       ksuid.New().String(),
		Name:     name,
		Phone:    phone, // later on I will add 233 infront
		Email:    email,
		Password: password, ///I will has before it saves
	}
	//when I write business logic in repo, it should come here and then I save the data in table

	return u, nil
}

func (s *userRepository) UpdateUser(ctx context.Context, name string, phone string, email string, id string) (*User, error) {
	u := &User{
		ID:    id,
		Name:  name,
		Phone: phone, // later on I will add 233 infront
		Email: email,
	}
	//when I write business logic in repo, it should come here and then I save the data in table
	return u, nil
}

func (s *userRepository) GetUsers(ctx context.Context, take uint64, skip uint64) ([]User, error) {
	var u []User
	//when I write business logic in repo, it should come here and then I save the data in table
	return u, nil
}

func (s *userRepository) GetUser(ctx context.Context, id string) (*User, error) {
	u := &User{}
	//when I write business logic in repo, it should come here and then I save the data in table
	return u, nil
}

func (s *userRepository) DeleteUser(ctx context.Context, id string) (bool, error) {
	// u := &User{}
	//when I write business logic in repo, it should come here and then I save the data in table
	return true, nil
}
