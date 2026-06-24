package user

import "errors"

var users = make(map[int]User)

// Add user
func CreateUser(u User) error {
	if _, exists := users[u.ID]; exists {
		return errors.New("user already exists")
	}
	users[u.ID] = u
	return nil
}

// Get all users
func GetUsers() []User {
	list := []User{}
	for _, u := range users {
		list = append(list, u)
	}
	return list
}

// Get user by ID
func GetUser(id int) (User, error) {
	u, exists := users[id]
	if !exists {
		return User{}, errors.New("user not found")
	}
	return u, nil
}

// Update user
func UpdateUser(id int, u User) error {
	if _, exists := users[id]; !exists {
		return errors.New("user not found")
	}
	users[id] = u
	return nil
}

// Delete user
func DeleteUser(id int) error {
	if _, exists := users[id]; !exists {
		return errors.New("user not found")
	}
	delete(users, id)
	return nil
}
