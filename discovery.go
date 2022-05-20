package sprintsd

import "context"

type Discovery struct{
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

type Registration struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Port    uint   `json:"port"`
}

type Store interface {
	Save(context.Context, *Registration) error
	Query(context.Context, string) (*Registration, error)
}