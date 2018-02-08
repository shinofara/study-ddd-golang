package channel

import "gitlab.com/shinofara/alpha/domain/type"

type Repository interface {
	Find(id _type.ChannelID) (*Channel, error)
	Add(c *Channel) (*Channel, error)
}
