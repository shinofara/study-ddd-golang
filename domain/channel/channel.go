package channel

import (
	"context"

	"cloud.google.com/go/firestore"
	"gitlab.com/shinofara/alpha/domain/internal"
	"gitlab.com/shinofara/alpha/domain/message"
	"gitlab.com/shinofara/alpha/domain/type"
	"gitlab.com/shinofara/alpha/domain/user"
)

const collection = "channel"

type Channel struct {
	ID      _type.ChannelID `firestore:"-"`
	Name    string
	OwnerID _type.UserID

	Owner    *user.User         `firestore:"-"`
	Messages []*message.Message `firestore:"-"`
	Members  []*user.User       `firestore:"-"`
}

func (c *Channel) SetID(id string) {
	c.ID = _type.ChannelID(id)
}

type Repository struct {
	ctx context.Context
	cli *firestore.Client
}

func New(cli *firestore.Client, ctx context.Context) *Repository {
	return &Repository{
		cli: cli,
		ctx: ctx,
	}
}

// Find ChannelIDを元にチャンネル情報を取得
func (r *Repository) Find(id _type.ChannelID) (*Channel, error) {
	ref, err := r.cli.Collection(collection).Doc(string(id)).Get(r.ctx)
	if err != nil {
		return nil, err
	}

	c := new(Channel)
	if err := internal.Convert(ref, c); err != nil {
		return nil, err
	}

	return c, nil
}

// Add 新しいチャンネルを追加
func (r *Repository) Add(c *Channel) (*Channel, error) {
	ref, _, err := r.cli.Collection(collection).Add(r.ctx, c)
	if err != nil {
		return nil, err
	}
	cc := *c
	internal.SetID(&cc, ref)
	return &cc, nil
}
