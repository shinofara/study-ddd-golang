package user

import (
	"context"

	"cloud.google.com/go/firestore"
	"gitlab.com/shinofara/alpha/domain/data/type"
	"gitlab.com/shinofara/alpha/domain/data/user"
	"gitlab.com/shinofara/alpha/infrastructure/firestore/internal"
)

type Repository struct {
	cli *firestore.Client
	ctx context.Context
}

const collection = "user"

func New(cli *firestore.Client, ctx context.Context) user.Repository {
	return &Repository{
		cli: cli,
		ctx: ctx,
	}
}

func (r *Repository) Find(id _type.UserID) (*user.User, error) {
	ref, err := r.cli.Collection(collection).Doc(string(id)).Get(r.ctx)
	if err != nil {
		return nil, err
	}

	u := new(user.User)
	if err := internal.Convert(ref, u); err != nil {
		return nil, err
	}

	return u, nil
}

// Add アイテムを追加するKeyは自動で振られる
func (r *Repository) Add(entity *user.User) (*user.User, error) {
	ref, _, err := r.cli.Collection(collection).Add(r.ctx, entity)
	if err != nil {
		return nil, err
	}
	u := *entity
	internal.SetID(&u, ref)
	return &u, nil
}
