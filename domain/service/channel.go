package service

import (
	"context"

	"cloud.google.com/go/firestore"
	"gitlab.com/shinofara/alpha/domain/channel"
	"gitlab.com/shinofara/alpha/domain/post"
	"gitlab.com/shinofara/alpha/domain/type"
	"gitlab.com/shinofara/alpha/domain/user"
)

type Channel struct {
	channelRepo *channel.Repository
	userRepo    *user.Repository
	postRepo    *post.Repository
}

func New(cli *firestore.Client, ctx context.Context) *Channel {
	return &Channel{
		channelRepo: channel.New(cli, ctx),
		userRepo:    user.New(cli, ctx),
		postRepo:    post.New(cli, ctx),
	}
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

	ch.Posts, err = c.postRepo.FindAllByChannelID(channelID)
	if err != nil {
		return nil, err
	}

	return ch, nil
}
