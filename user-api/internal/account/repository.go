package account

import (
	"database/sql"
	"errors"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) UpdateBalanceTx(tx *sql.Tx, id int, amount float64) error {
	res, err := tx.Exec("UPDATE accounts SET balance = balance + ? WHERE id = ?", amount, id)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("account not found")
	}
	return nil
}

func (r *Repository) GetByID(id int) (Account, error) {
	var a Account
	err := r.db.QueryRow("SELECT id, owner, balance FROM accounts WHERE id = ?", id).
		Scan(&a.ID, &a.Owner, &a.Balance)
	if err == sql.ErrNoRows {
		return Account{}, errors.New("account not found")
	}
	return a, err
}
