package userAccount

import "time"

type UserAccount struct {
	ID           string
	EmailAddress string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
