package user

import "gitlab.com/shinofara/alpha/domain/type"

type User struct {
	ID   _type.UserID
	Name string
}

func Find(id _type.UserID) (*User, error) {
	return &User{
		ID: id,
	}, nil
}

func FindByIds(ids []*_type.UserID) ([]*User, error) {
	return []*User{
		{
			ID: *ids[0],
		},
	}, nil
}
