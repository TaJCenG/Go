package user

import (
	"Day1Utils/UserManagementService/Internal/database"
	"database/sql"
	"errors"
)

// CreateUser inserts a new user into MySQL
func CreateUser(u User) error {
	_, err := database.DB.Exec(
		"INSERT INTO users (id, name, email) VALUES (?, ?, ?)",
		u.ID, u.Name, u.Email,
	)
	if err != nil {
		return err
	}
	return nil
}

// GetUsers retrieves all users
func GetUsers() ([]User, error) {
	rows, err := database.DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

// GetUser retrieves a user by ID
func GetUser(id int) (User, error) {
	var u User
	err := database.DB.QueryRow(
		"SELECT id, name, email FROM users WHERE id = ?",
		id,
	).Scan(&u.ID, &u.Name, &u.Email)

	if err == sql.ErrNoRows {
		return User{}, errors.New("user not found")
	}
	if err != nil {
		return User{}, err
	}
	return u, nil
}

// UpdateUser updates an existing user
func UpdateUser(id int, u User) error {
	res, err := database.DB.Exec(
		"UPDATE users SET name = ?, email = ? WHERE id = ?",
		u.Name, u.Email, id,
	)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

// DeleteUser deletes a user by ID
func DeleteUser(id int) error {
	res, err := database.DB.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}
