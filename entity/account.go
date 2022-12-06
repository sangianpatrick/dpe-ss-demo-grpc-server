package entity

import "time"

type Account struct {
	Id        int64
	Email     string
	Password  string
	Name      string
	CreatedAt time.Time
	UpdatedAt *time.Time
}
