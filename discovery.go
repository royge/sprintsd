package sprintsd

import "context"

type Discovery struct {
	store Store
}

func NewDiscovery(store Store) *Discovery {
	return &Discovery{
		store: store,
	}
}

func (d *Discovery) Enroll(ctx context.Context, reg *Registration) error {
	return d.store.Save(ctx, reg)
}

func (d *Discovery) Locate(ctx context.Context, name string) (
	r *Registration, err error,
) {
	r, err = d.store.Query(ctx, name)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (d *Discovery) Forget(ctx context.Context, name string) error {
	return d.store.Delete(ctx, name)
}

type Registration struct {
	ID      string `json:"id" firestore:"id"`
	Name    string `json:"name" firestore:"name"`
	Address string `json:"address" firestore:"address"`
	Port    int    `json:"port" firestore:"port"`
}

type Store interface {
	Save(context.Context, *Registration) error
	Query(context.Context, string) (*Registration, error)
	Delete(context.Context, string) error
}
