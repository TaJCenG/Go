package user

import (
	"context"
	"database/sql"
	"errors"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(u User) error {
	_, err := r.db.Exec("INSERT INTO users (id, name, email) VALUES (?, ?, ?)", u.ID, u.Name, u.Email)
	return err
}

func (r *Repository) GetAll() ([]User, error) {
	rows, err := r.db.Query("SELECT id, name, email FROM users")
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
	return users, nil
}

func (r *Repository) GetByID(id int) (User, error) {
	var u User
	err := r.db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id).Scan(&u.ID, &u.Name, &u.Email)
	if err == sql.ErrNoRows {
		return User{}, errors.New("user not found")
	}
	return u, err
}

func (r *Repository) Update(id int, u User) error {
	res, err := r.db.Exec("UPDATE users SET name=?, email=? WHERE id=?", u.Name, u.Email, id)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (r *Repository) Delete(id int) error {
	res, err := r.db.Exec("DELETE FROM users WHERE id=?", id)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (r *Repository) GetByIDCtx(ctx context.Context, id int) (User, error) {
	var u User
	err := r.db.QueryRowContext(ctx,
		"SELECT id, name, email FROM users WHERE id = ?", id,
	).Scan(&u.ID, &u.Name, &u.Email)

	if err == sql.ErrNoRows {
		return User{}, errors.New("user not found")
	}
	return u, err
}
