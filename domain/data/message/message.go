package message

import (
	"gitlab.com/shinofara/alpha/domain/data/type"
	"gitlab.com/shinofara/alpha/domain/data/user"
)

type Message struct {
	ID        _type.MessageID `firestore:"-"`
	Text      string
	UserID    _type.UserID
	ChannelID _type.ChannelID

	User *user.User `firestore:"-"`
}

func NewMessage(channelID _type.ChannelID, userID _type.UserID, msg string) *Message {
	mess := &Message{
		ChannelID: channelID,
		UserID:    userID,
		Text:      msg,
	}

	return mess
}

func (m *Message) SetID(id string) {
	m.ID = _type.MessageID(id)
}
