package message

import (
	"context"

	"cloud.google.com/go/firestore"
	"gitlab.com/shinofara/alpha/domain/data/message"
	"gitlab.com/shinofara/alpha/domain/data/type"
	"gitlab.com/shinofara/alpha/infrastructure/firestore/internal"
)

type Repository struct {
	cli *firestore.Client
	ctx context.Context
}

const collection = "message"

func New(cli *firestore.Client, ctx context.Context) message.Repository {
	return &Repository{
		cli: cli,
		ctx: ctx,
	}
}

// Set アイテムを追加する
func (r *Repository) Set(key string, entity *message.Message) error {
	_, err := r.cli.Collection(collection).Doc(key).Set(r.ctx, entity)

	return err
}

// Add アイテムを追加するKeyは自動で振られる
func (r *Repository) Add(entity *message.Message) (*message.Message, error) {
	ref, _, err := r.cli.Collection(collection).Add(r.ctx, entity)
	if err != nil {
		return nil, err
	}
	m := *entity
	internal.SetID(&m, ref)
	return &m, nil
}

// Find IDを元にメッセージを取得
func (r *Repository) Find(id _type.MessageID) (*message.Message, error) {
	ref, err := r.cli.Collection(collection).Doc(string(id)).Get(r.ctx)
	if err != nil {
		return nil, err
	}

	c := new(message.Message)
	if err := internal.Convert(ref, c); err != nil {
		return nil, err
	}

	return c, nil
}

// FindAllByChannelID channelIDでチャンネル内のメッセージを取得
func (r *Repository) FindAllByChannelID(id _type.ChannelID) ([]*message.Message, error) {
	var messages []*message.Message

	m := r.cli.Collection(collection).Where("ChannelID", "==", id).Documents(r.ctx)
	docs, err := m.GetAll()
	if err != nil {
		return nil, err
	}

	for _, doc := range docs {
		c := new(message.Message)
		if err := internal.Convert(doc, c); err != nil {
			return nil, err
		}
		messages = append(messages, c)
	}

	return messages, nil
}
