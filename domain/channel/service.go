package channel

import (
	"gitlab.com/shinofara/alpha/domain/message"
	"gitlab.com/shinofara/alpha/domain/type"
	"gitlab.com/shinofara/alpha/domain/user"
)

type Service struct {
	channelRepo Repository
	userRepo    user.Repository
	messageRepo message.Repository
}

func NewService(
	channelRepo Repository,
	userRepo user.Repository,
	messageRepo message.Repository) *Service {

	return &Service{
		channelRepo: channelRepo,
		userRepo:    userRepo,
		messageRepo: messageRepo,
	}
}

// Create 新しいチャンネルを作成
func (c *Service) Create(name string, owner *user.User) (*Channel, error) {
	ch := &Channel{
		Name:    name,
		OwnerID: owner.ID,
		Owner:   owner,
	}

	return c.channelRepo.Add(ch)
}

// InitialDisplay channelIDを元に、チャンネル表示に必要な情報をChannel Entityに集約して返す
func (c *Service) InitialDisplay(channelID _type.ChannelID) (*Channel, error) {
	ch, err := c.channelRepo.Find(channelID)
	if err != nil {
		return nil, err
	}

	ch.Owner, err = c.userRepo.Find(ch.OwnerID)
	if err != nil {
		return nil, err
	}

	ch.Messages, err = c.messageRepo.FindAllByChannelID(channelID)
	if err != nil {
		return nil, err
	}

	return ch, nil
}
