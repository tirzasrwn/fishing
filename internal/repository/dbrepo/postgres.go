package dbrepo

import (
	"context"
	"time"

	"github.com/tirzasrwn/fishing/internal/models"
)

func (m *postgresDBRepo) InsertAccount(account models.Account) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
INSERT INTO
  accounts (email, password)
VALUES
  ($1, $2);
`
	var id int
	err := m.DB.QueryRowContext(ctx, stmt, account.Email, account.Password).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}

func (m *postgresDBRepo) AllAccounts() ([]models.Account, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var accounts []models.Account
	stmt := `SELECT email, password FROM accounts`
	rows, err := m.DB.QueryContext(ctx, stmt)
	if err != nil {
		return accounts, err
	}
	defer rows.Close()

	for rows.Next() {
		var a models.Account
		err := rows.Scan(&a.Email, &a.Password)
		if err != nil {
			return accounts, err
		}
		accounts = append(accounts, a)
	}
	if err = rows.Err(); err != nil {
		return accounts, err
	}
	return accounts, nil
}
