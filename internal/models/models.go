package models

import "time"

type Account struct {
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
