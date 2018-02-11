package internal

import (
	"cloud.google.com/go/firestore"
	"gitlab.com/shinofara/alpha/domain"
)

func Convert(snapshot *firestore.DocumentSnapshot, obj domain.Entity) error {
	if err := snapshot.DataTo(obj); err != nil {
		return err
	}

	SetID(obj, snapshot.Ref)

	return nil
}

func SetID(obj domain.Entity, ref *firestore.DocumentRef) {
	obj.SetID(ref.ID)
}
