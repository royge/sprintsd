package sprintsd_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/royge/sprintsd"
	"github.com/stretchr/testify/mock"
)

func Test_NewDiscovery(t *testing.T) {
	store := new(mockStore)
	d := sprintsd.NewDiscovery(store)
	if d == nil {
		t.Fatalf("unexpected nil discovery")
	}
}

func Test_Discovery_Enroll(t *testing.T) {
	store := new(mockStore)
	defer store.AssertExpectations(t)

	d := sprintsd.NewDiscovery(store)
	ctx := context.Background()

	reg := &sprintsd.Registration{
		Name:    "myservice",
		Address: "http://localhost",
		Port:    80,
	}

	t.Run("ok", func(t *testing.T) {
		store.On("Save", ctx, reg).Return(nil).Once()

		err := d.Enroll(ctx, reg)
		if err != nil {
			t.Fatalf("registration failed: %v", err)
		}
	})

	t.Run("error", func(t *testing.T) {
		store.On("Save", ctx, reg).Return(errors.New("store error")).Once()

		err := d.Enroll(ctx, reg)
		if err == nil {
			t.Fatal("unexpected nil error")
		}
	})
}

func Test_Discovery_Locate(t *testing.T) {
	store := new(mockStore)
	defer store.AssertExpectations(t)

	d := sprintsd.NewDiscovery(store)
	ctx := context.Background()

	reg := &sprintsd.Registration{
		Name:    "myservice",
		Address: "http://localhost",
		Port:    80,
	}

	t.Run("found", func(t *testing.T) {
		store.On("Query", ctx, reg.Name).Return(reg, nil).Once()

		got, err := d.Locate(ctx, reg.Name)
		if err != nil {
			t.Fatalf("unable to locate enrolled service: %v", err)
		}

		if !reflect.DeepEqual(reg, got) {
			t.Errorf("want %#v, got %#v", reg, got)
		}
	})

	t.Run("not found", func(t *testing.T) {
		store.On("Query", ctx, reg.Name).Return(nil, errors.New("not found")).Once()

		got, err := d.Locate(ctx, reg.Name)
		if err == nil {
			t.Fatal("unexpected nil error")
		}

		if got != nil {
			t.Error("expected nil registration")
		}
	})
}

type mockStore struct {
	mock.Mock
}

func (m *mockStore) Save(ctx context.Context, reg *sprintsd.Registration) error {
	return m.Called(ctx, reg).Error(0)
}

func (m *mockStore) Query(ctx context.Context, name string) (
	*sprintsd.Registration, error,
) {
	args := m.Called(ctx, name)
	err := args.Error(1)

	if err != nil {
		return nil, err
	}

	return args.Get(0).(*sprintsd.Registration), err
}
