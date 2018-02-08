package message

import (
	"gitlab.com/shinofara/alpha/domain/type"
	"gitlab.com/shinofara/alpha/domain/user"
)

// Message メッセージを操作する為に必要な、Repositoryなどを管理
type Service struct {
	userRepo    user.Repository
	messageRepo Repository
}

func NewService(messageRepo Repository) *Service {
	return &Service{
		messageRepo: messageRepo,
	}
}

// Post 指定したチャンネルにメッセージをポストする
func (m *Service) Post(channelID _type.ChannelID, userID _type.UserID, msg string) (*Message, error) {
	mess := &Message{
		ChannelID: channelID,
		UserID:    userID,
		Text:      msg,
	}

	return m.messageRepo.Add(mess)
}
