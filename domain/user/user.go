package user

import (
	"context"

	"cloud.google.com/go/firestore"
	"gitlab.com/shinofara/alpha/domain/internal"
	"gitlab.com/shinofara/alpha/domain/type"
)

type User struct {
	ID   _type.UserID `firestore:"-"`
	Name string
}

func (u *User) SetID(id string) {
	u.ID = _type.UserID(id)
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

const collection = "user"

func New(cli *firestore.Client, ctx context.Context) *Repository {
	return &Repository{
		cli: cli,
		ctx: ctx,
	}
}

func (r *Repository) Find(id _type.UserID) (*User, error) {
	ref, err := r.cli.Collection(collection).Doc(string(id)).Get(r.ctx)
	if err != nil {
		return nil, err
	}

	u := new(User)
	if err := internal.Convert(ref, u); err != nil {
		return nil, err
	}

	return u, nil
}

// Add アイテムを追加するKeyは自動で振られる
func (r *Repository) Add(entity *User) (*User, error) {
	ref, _, err := r.cli.Collection(collection).Add(r.ctx, entity)
	if err != nil {
		return nil, err
	}
	u := *entity
	internal.SetID(&u, ref)
	return &u, nil
}
