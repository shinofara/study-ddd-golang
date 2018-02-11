package user

import (
	"gitlab.com/shinofara/alpha/domain/data/type"
)

type User struct {
	ID   _type.UserID `firestore:"-"`
	Name string
}

func (u *User) SetID(id string) {
	u.ID = _type.UserID(id)
}
