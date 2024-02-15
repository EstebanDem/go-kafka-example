package user

import "time"

type User struct {
	ID             string
	Name           string
	Email          string
	PhoneNumber    string
	EmailConfirmed bool
	PhoneConfirmed bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
