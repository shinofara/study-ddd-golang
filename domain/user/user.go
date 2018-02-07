package user

import (
	"context"

	"cloud.google.com/go/firestore"
	"gitlab.com/shinofara/alpha/domain/internal"
	"gitlab.com/shinofara/alpha/domain/type"
)

type User struct {
	ID   _type.UserID
	Name string
}

func Find(id _type.UserID) (*User, error) {
	return &User{
		ID: id,
	}, nil
}

type Repository struct {
	cli *firestore.Client
	ctx context.Context
}

const Collection = "user"

func New(cli *firestore.Client, ctx context.Context) *Repository {
	return &Repository{
		cli: cli,
		ctx: ctx,
	}
}

func (r *Repository) Find(id _type.UserID) (*User, error) {
	ref, err := r.cli.Collection(Collection).Doc(string(id)).Get(r.ctx)
	if err != nil {
		return nil, err
	}

	u := new(User)
	if err := internal.Convert(ref, &u); err != nil {
		return nil, err
	}

	return u, nil
}
