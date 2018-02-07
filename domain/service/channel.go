package service

import (
	"context"

	"cloud.google.com/go/firestore"
	"gitlab.com/shinofara/alpha/domain/channel"
	"gitlab.com/shinofara/alpha/domain/message"
	"gitlab.com/shinofara/alpha/domain/type"
	"gitlab.com/shinofara/alpha/domain/user"
)

type Channel struct {
	channelRepo *channel.Repository
	userRepo    *user.Repository
	postRepo    *message.Repository
}

func NewChannel(cli *firestore.Client, ctx context.Context) *Channel {
	return &Channel{
		channelRepo: channel.New(cli, ctx),
		userRepo:    user.New(cli, ctx),
		postRepo:    message.New(cli, ctx),
	}
}

// Create 新しいチャンネルを作成
func (c *Channel) Create(name string, owner *user.User) (*channel.Channel, error) {
	ch := &channel.Channel{
		Name:    name,
		OwnerID: owner.ID,
		Owner:   owner,
	}

	return c.channelRepo.Add(ch)
}

// InitialDisplay channelIDを元に、チャンネル表示に必要な情報をChannel Entityに集約して返す
func (c *Channel) InitialDisplay(channelID _type.ChannelID) (*channel.Channel, error) {
	ch, err := c.channelRepo.Find(channelID)
	if err != nil {
		return nil, err
	}

	ch.Owner, err = c.userRepo.Find(ch.OwnerID)
	if err != nil {
		return nil, err
	}

	ch.Messages, err = c.postRepo.FindAllByChannelID(channelID)
	if err != nil {
		return nil, err
	}

	return ch, nil
}
