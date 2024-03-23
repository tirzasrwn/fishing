package repository

import "github.com/tirzasrwn/fishing/internal/models"

type DatabaseRepo interface {
	InsertAccount(models.Account) error
	AllAccounts() ([]models.Account, error)
}
