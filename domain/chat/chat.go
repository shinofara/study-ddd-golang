package chat

import (
	"context"

	"cloud.google.com/go/firestore"
)

type Chat struct {
	Message string
	User    string
}

type Repository struct {
	cli *firestore.Client
	ctx context.Context
}

const Collection = "chat"

func New(cli *firestore.Client, ctx context.Context) *Repository {
	return &Repository{
		cli: cli,
		ctx: ctx,
	}
}
func (r *Repository) Add(entity *Chat) error {
	_, err := r.cli.Collection(Collection).Doc("aaa").Set(r.ctx, entity)

	return err
}
