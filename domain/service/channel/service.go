package channel

import (
	"gitlab.com/shinofara/alpha/domain/data/channel"
	"gitlab.com/shinofara/alpha/domain/data/message"
	"gitlab.com/shinofara/alpha/domain/data/type"
	"gitlab.com/shinofara/alpha/domain/data/user"
)

type Service struct {
	channelRepo channel.Repository
	userRepo    user.Repository
	messageRepo message.Repository
}

func New(
	channelRepo channel.Repository,
	userRepo user.Repository,
	messageRepo message.Repository) *Service {

	return &Service{
		channelRepo: channelRepo,
		userRepo:    userRepo,
		messageRepo: messageRepo,
	}
}

// Create 新しいチャンネルを作成
func (c *Service) Create(name string, owner *user.User) (*channel.Channel, error) {
	ch := &channel.Channel{
		Name:    name,
		OwnerID: owner.ID,
		Owner:   owner,
	}

	return c.channelRepo.Add(ch)
}

// InitialDisplay channelIDを元に、チャンネル表示に必要な情報をChannel Entityに集約して返す
func (c *Service) InitialDisplay(channelID _type.ChannelID) (*channel.Channel, error) {
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
