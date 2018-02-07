package service

import (
	"context"

	"cloud.google.com/go/firestore"
	"gitlab.com/shinofara/alpha/domain/channel"
	"gitlab.com/shinofara/alpha/domain/message"
	"gitlab.com/shinofara/alpha/domain/type"
	"gitlab.com/shinofara/alpha/domain/user"
)

type Message struct {
	channelRepo *channel.Repository
	userRepo    *user.Repository
	messageRepo *message.Repository
}

func NewMessage(cli *firestore.Client, ctx context.Context) *Message {
	return &Message{
		channelRepo: channel.New(cli, ctx),
		userRepo:    user.New(cli, ctx),
		messageRepo: message.New(cli, ctx),
	}
}

func (m *Message) Post(channelID _type.ChannelID, userID _type.UserID, text string) (*message.Message, error) {
	mess := &message.Message{
		ChannelID: channelID,
		UserID:    userID,
		Text:      text,
	}

	return m.messageRepo.Add(mess)
}
