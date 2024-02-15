package user

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserRepository interface {
	GetByID(id string) (*User, error)
	Save(u User) error
	GetAll() ([]User, error)
}
