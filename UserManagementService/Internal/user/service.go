package user

import (
	"errors"
)

// UserService wraps repository functions with business logic
type UserService struct{}

func (s *UserService) CreateUser(u User) error {
	// Business validation
	if u.Name == "" || u.Email == "" {
		return errors.New("name and email are required")
	}
	return CreateUser(u)
}

func (s *UserService) GetUsers() ([]User, error) {
	return GetUsers()
}

func (s *UserService) GetUser(id int) (User, error) {
	return GetUser(id)
}

func (s *UserService) UpdateUser(id int, u User) error {
	if u.Name == "" || u.Email == "" {
		return errors.New("name and email are required")
	}
	return UpdateUser(id, u)
}

func (s *UserService) DeleteUser(id int) error {
	return DeleteUser(id)
}
