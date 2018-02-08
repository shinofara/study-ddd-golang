package service

import (
	"context"

	"cloud.google.com/go/firestore"
	"gitlab.com/shinofara/alpha/domain/user"
)

// User メッセージを操作する為に必要な、Repositoryなどを管理
type User struct {
	userRepo *user.Repository
}

func NewUser(cli *firestore.Client, ctx context.Context) *User {
	return &User{
		userRepo: user.New(cli, ctx),
	}
}

func (m *User) Register(name string) (*user.User, error) {
	u := &user.User{
		Name: name,
	}

	return m.userRepo.Add(u)
}
