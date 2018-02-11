package message

import (
	"errors"

	"gitlab.com/shinofara/alpha/domain/data/message"
	"gitlab.com/shinofara/alpha/domain/data/type"
	"gitlab.com/shinofara/alpha/domain/data/user"
)

// Message メッセージを操作する為に必要な、Repositoryなどを管理
type Service struct {
	userRepo    user.Repository
	messageRepo message.Repository
}

func New(messageRepo message.Repository) *Service {
	return &Service{
		messageRepo: messageRepo,
	}
}

// Post 指定したチャンネルにメッセージをポストする
func (m *Service) Post(channelID _type.ChannelID, userID _type.UserID, msg string, spec message.Specification) (*message.Message, error) {
	mess := message.NewMessage(channelID, userID, msg)
	if ok := spec.IsSatisfiedBy(mess); !ok {
		return nil, errors.New("メッセージ内容が仕様を満たしていません")
	}

	return m.messageRepo.Add(mess)
}
