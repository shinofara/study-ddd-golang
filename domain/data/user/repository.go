package user

import "gitlab.com/shinofara/alpha/domain/data/type"

type Repository interface {
	Add(entity *User) (*User, error)
	Find(id _type.UserID) (*User, error)
}
