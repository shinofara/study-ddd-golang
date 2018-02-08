package service

import (
	"gitlab.com/shinofara/alpha/domain/channel"
	"gitlab.com/shinofara/alpha/domain/message"
	"gitlab.com/shinofara/alpha/domain/type"
	"gitlab.com/shinofara/alpha/domain/user"
)

// Message メッセージを操作する為に必要な、Repositoryなどを管理
type Message struct {
	channelRepo *channel.Repository
	userRepo    *user.Repository
	messageRepo *message.Repository
}

func NewMessage(messageRepo *message.Repository) *Message {
	return &Message{
		messageRepo: messageRepo,
	}
}

// Post 指定したチャンネルにメッセージをポストする
func (m *Message) Post(channelID _type.ChannelID, userID _type.UserID, msg string) (*message.Message, error) {
	mess := &message.Message{
		ChannelID: channelID,
		UserID:    userID,
		Text:      msg,
	}

	return m.messageRepo.Add(mess)
}
