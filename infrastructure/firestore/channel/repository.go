package channel

import (
	"context"

	"cloud.google.com/go/firestore"
	"gitlab.com/shinofara/alpha/domain/channel"
	"gitlab.com/shinofara/alpha/domain/type"
	internal "gitlab.com/shinofara/alpha/infrastructure/firestore/internal"
)

type Repository struct {
	ctx context.Context
	cli *firestore.Client
}

func New(cli *firestore.Client, ctx context.Context) channel.Repository {
	return &Repository{
		cli: cli,
		ctx: ctx,
	}
}

const collection = "channel"

// Find ChannelIDを元にチャンネル情報を取得
func (r *Repository) Find(id _type.ChannelID) (*channel.Channel, error) {
	ref, err := r.cli.Collection(collection).Doc(string(id)).Get(r.ctx)
	if err != nil {
		return nil, err
	}

	c := new(channel.Channel)
	if err := internal.Convert(ref, c); err != nil {
		return nil, err
	}

	return c, nil
}

// Add 新しいチャンネルを追加
func (r *Repository) Add(c *channel.Channel) (*channel.Channel, error) {
	ref, _, err := r.cli.Collection(collection).Add(r.ctx, c)
	if err != nil {
		return nil, err
	}
	cc := *c
	internal.SetID(&cc, ref)
	return &cc, nil
}
