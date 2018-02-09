package message

import (
	"errors"

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
func (m *Service) Post(channelID _type.ChannelID, userID _type.UserID, msg string, spec Specification) (*Message, error) {
	mess := NewMessage(channelID, userID, msg)
	if ok := spec.IsSatisfiedBy(mess); !ok {
		return nil, errors.New("メッセージ内容が仕様を満たしていません")
	}

	return m.messageRepo.Add(mess)
}
