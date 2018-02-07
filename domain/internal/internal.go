package internal

import "cloud.google.com/go/firestore"

func Convert(snapshot *firestore.DocumentSnapshot, obj interface{}) error {
	if err := snapshot.DataTo(obj); err != nil {
		return err
	}
	return nil
}
