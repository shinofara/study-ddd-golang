package message

import (
	"gitlab.com/shinofara/alpha/domain/type"
	"gitlab.com/shinofara/alpha/domain/user"
)

type Message struct {
	ID        _type.MessageID `firestore:"-"`
	Text      string
	UserID    _type.UserID
	ChannelID _type.ChannelID

	User *user.User `firestore:"-"`
}

func (m *Message) SetID(id string) {
	m.ID = _type.MessageID(id)
}
