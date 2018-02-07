package channel

import (
	"context"

	"cloud.google.com/go/firestore"
	"gitlab.com/shinofara/alpha/domain/internal"
	"gitlab.com/shinofara/alpha/domain/post"
	"gitlab.com/shinofara/alpha/domain/type"
	"gitlab.com/shinofara/alpha/domain/user"
)

const Collection = "channel"

type Channel struct {
	ID        _type.ChannelID
	Name      string
	OwnerID   _type.UserID
	MemberIDs []*_type.UserID

	Owner   *user.User   `firestore:"-"`
	Posts   []*post.Post `firestore:"-"`
	Members []*user.User `firestore:"-"`
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

func (r *Repository) Find(id _type.ChannelID) (*Channel, error) {
	ref, err := r.cli.Collection(Collection).Doc(string(id)).Get(r.ctx)
	if err != nil {
		return nil, err
	}

	c := new(Channel)
	if err := internal.Convert(ref, &c); err != nil {
		return nil, err
	}

	return c, nil
}
