package user

import (
	"gitlab.com/shinofara/alpha/domain/type"
)

type User struct {
	ID   _type.UserID `firestore:"-"`
	Name string
}

func (u *User) SetID(id string) {
	u.ID = _type.UserID(id)
}
