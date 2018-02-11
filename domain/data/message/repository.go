package message

import "gitlab.com/shinofara/alpha/domain/data/type"

type Repository interface {
	Set(key string, entity *Message) error
	Add(c *Message) (*Message, error)
	Find(id _type.MessageID) (*Message, error)
	FindAllByChannelID(id _type.ChannelID) ([]*Message, error)
}
