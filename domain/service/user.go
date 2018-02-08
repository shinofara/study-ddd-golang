package service

import (
	"gitlab.com/shinofara/alpha/domain/user"
)

// User メッセージを操作する為に必要な、Repositoryなどを管理
type User struct {
	userRepo user.Repository
}

func NewUser(userRepo user.Repository) *User {
	return &User{
		userRepo: userRepo,
	}
}

func (m *User) Register(name string) (*user.User, error) {
	u := &user.User{
		Name: name,
	}

	return m.userRepo.Add(u)
}
