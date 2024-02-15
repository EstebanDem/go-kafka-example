package storage

import "registration-service/internal/domain/user"

type UserMemoryRepository struct {
	users map[string]*user.User
}

func NewUserInMemoryRepository() *UserMemoryRepository {
	return &UserMemoryRepository{
		users: map[string]*user.User{},
	}
}

func (ur *UserMemoryRepository) GetByID(id string) (*user.User, error) {
	userFound, ok := ur.users[id]
	if !ok {
		return nil, user.ErrUserNotFound
	}
	return userFound, nil
}

func (ur *UserMemoryRepository) Save(u user.User) error {
	ur.users[u.ID] = &u
	return nil
}

func (ur *UserMemoryRepository) GetAll() ([]user.User, error) {
	var allUsers []user.User
	for _, u := range ur.users {
		allUsers = append(allUsers, *u)
	}
	return allUsers, nil
}
