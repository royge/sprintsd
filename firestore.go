package sprintsd

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
)

type Firestore struct {
	client     *firestore.Client
	collection string
}

func NewFirestore(client *firestore.Client, collection string) *Firestore {
	if client == nil {
		panic("unexpected nil firestore client")
	}

	if collection == "" {
		panic("unexpected empty collection name")
	}

	return &Firestore{
		client:     client,
		collection: collection,
	}
}

func (fs *Firestore) Save(ctx context.Context, reg *Registration) error {
	doc := fs.client.Collection(fs.collection).Doc(reg.ID)
	return fs.client.RunTransaction(ctx, func(_ context.Context, tx *firestore.Transaction) error {
		return tx.Set(doc, reg)
	})
}

func (fs *Firestore) Query(context.Context, string) (*Registration, error) {
	return nil, errors.New("not available")
}

func (fs *Firestore) Delete(context.Context, string) error {
	return errors.New("not available")
}
