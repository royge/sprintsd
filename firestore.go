package sprintsd

import (
	"context"

	"cloud.google.com/go/firestore"
)

type Firestore struct {
	Client     *firestore.Client
	Collection string
}

func NewFirestore(client *firestore.Client, collection string) *Firestore {
	if client == nil {
		panic("unexpected nil firestore client")
	}

	if collection == "" {
		panic("unexpected empty collection name")
	}

	return &Firestore{
		Client:     client,
		Collection: collection,
	}
}

func (fs *Firestore) Save(context.Context, *Registration) error {
	return nil
}

func (fs *Firestore) Query(context.Context, string) (*Registration, error) {
	return nil, nil
}

func (fs *Firestore) Delete(context.Context, string) error {
	return nil
}
