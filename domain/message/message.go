package message

import (
	"context"

	"gitlab.com/shinofara/alpha/domain/user"

	"cloud.google.com/go/firestore"
	"gitlab.com/shinofara/alpha/domain/internal"
	"gitlab.com/shinofara/alpha/domain/type"
)

type Message struct {
	ID        _type.MessageID `firestore:"-"`
	Text      string
	UserID    _type.UserID
	ChannelID _type.ChannelID

	User *user.User `firestore:"-"`
}

func (m *Message) SetID(id string) {
	m.ID = _type.MessageID(id)
}

type Repository struct {
	cli *firestore.Client
	ctx context.Context
}

const Collection = "post"

func New(cli *firestore.Client, ctx context.Context) *Repository {
	return &Repository{
		cli: cli,
		ctx: ctx,
	}
}

// Set アイテムを追加する
func (r *Repository) Set(key string, entity *Message) error {
	_, err := r.cli.Collection(Collection).Doc(key).Set(r.ctx, entity)

	return err
}

// Add アイテムを追加するKeyは自動で振られる
func (r *Repository) Add(entity *Message) (*Message, error) {
	ref, _, err := r.cli.Collection(Collection).Add(r.ctx, entity)
	if err != nil {
		return nil, err
	}
	m := *entity
	internal.SetID(&m, ref)
	return &m, nil
}

func (r *Repository) Find(key string) (*Message, error) {
	ref, err := r.cli.Collection(Collection).Doc(key).Get(r.ctx)
	if err != nil {
		return nil, err
	}

	c := new(Message)
	if err := internal.Convert(ref, c); err != nil {
		return nil, err
	}

	return c, nil
}

func (r *Repository) FindAllByChannelID(id _type.ChannelID) ([]*Message, error) {
	var messages []*Message

	m := r.cli.Collection(Collection).Where("ChannelID", "==", id).Documents(r.ctx)
	docs, err := m.GetAll()
	if err != nil {
		return nil, err
	}

	for _, doc := range docs {
		c := new(Message)
		if err := internal.Convert(doc, c); err != nil {
			return nil, err
		}
		messages = append(messages, c)
	}

	return messages, nil
}
