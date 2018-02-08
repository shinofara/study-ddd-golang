package user

import "gitlab.com/shinofara/alpha/domain/type"

type Repository interface {
	Add(entity *User) (*User, error)
	Find(id _type.UserID) (*User, error)
}
