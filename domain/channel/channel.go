package channel

import (
	"gitlab.com/shinofara/alpha/domain/message"
	"gitlab.com/shinofara/alpha/domain/type"
	"gitlab.com/shinofara/alpha/domain/user"
)

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
