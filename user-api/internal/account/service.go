package account

import "database/sql"

type Service struct {
	repo *Repository
	db   *sql.DB
}

func NewService(repo *Repository, db *sql.DB) *Service {
	return &Service{repo: repo, db: db}
}

// TransferMoney ensures atomic debit/credit
func (s *Service) TransferMoney(fromID, toID int, amount float64) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	// Step 1: Deduct from sender
	if err := s.repo.UpdateBalanceTx(tx, fromID, -amount); err != nil {
		tx.Rollback()
		return err
	}

	// Step 2: Add to receiver
	if err := s.repo.UpdateBalanceTx(tx, toID, amount); err != nil {
		tx.Rollback()
		return err
	}

	// Step 3: Commit if both succeed
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
