package channel

import (
	"context"

	"cloud.google.com/go/firestore"
	"gitlab.com/shinofara/alpha/domain/post"
	"gitlab.com/shinofara/alpha/domain/type"
	"gitlab.com/shinofara/alpha/domain/user"
)

type Channel struct {
	ID        _type.ChannelID
	Name      string
	OwnerID   _type.UserID
	MemberIDs []*_type.UserID

	ctx context.Context
	cli *firestore.Client
}

func (c *Channel) Posts() ([]*post.Post, error) {
	pRepo := post.New(c.cli, c.ctx)
	return pRepo.FindAllByChannelID(c.ID)
}

func (c *Channel) Owner() (*user.User, error) {
	return user.Find(c.OwnerID)
}

func (c *Channel) Members() ([]*user.User, error) {
	return user.FindByIds(c.MemberIDs)
}
