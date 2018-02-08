package internal

import "cloud.google.com/go/firestore"

type Entity interface {
	SetID(id string)
}

func Convert(snapshot *firestore.DocumentSnapshot, obj Entity) error {
	if err := snapshot.DataTo(obj); err != nil {
		return err
	}

	SetID(obj, snapshot.Ref)

	return nil
}

func SetID(obj Entity, ref *firestore.DocumentRef) {
	obj.SetID(ref.ID)
}
