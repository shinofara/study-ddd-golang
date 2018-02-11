package channel

import (
	"gitlab.com/shinofara/alpha/domain/data/message"
	"gitlab.com/shinofara/alpha/domain/data/type"
	"gitlab.com/shinofara/alpha/domain/data/user"
)

type Channel struct {
	ID      _type.ChannelID `firestore:"-"`
	Name    string
	OwnerID _type.UserID

	Owner    *user.User         `firestore:"-"`
	Messages []*message.Message `firestore:"-"`
	Members  []*user.User       `firestore:"-"`
}

func NewChannel(ownerID _type.UserID, name string) *Channel {
	c := &Channel{
		OwnerID: ownerID,
		Name:    name,
	}

	return c
}

func (c *Channel) SetID(id string) {
	c.ID = _type.ChannelID(id)
}
