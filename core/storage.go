package core

import (
	"context"

	"github.com/pkg/errors"

	"github.com/moonrhythm/archer/core/storage"
)

// Errors
var (
	ErrNotFound = errors.New("not found")
)

// Storage type
type Storage interface {
	Set(ctx context.Context, key string, value []byte) error
	Get(ctx context.Context, key string) ([]byte, error)
	Del(ctx context.Context, key string) error
	Keys(ctx context.Context, prefix string) ([]string, error)
	Exists(ctx context.Context, key string) (bool, error)
}

// SetStorage sets storage
func SetStorage(s Storage) {
	bus.Register(func(ctx context.Context, m *storage.Set) error {
		return s.Set(ctx, m.Key, m.Value)
	})

	bus.Register(func(ctx context.Context, m *storage.Get) (err error) {
		m.Result, err = s.Get(ctx, m.Key)
		return
	})

	bus.Register(func(ctx context.Context, m *storage.Del) error {
		return s.Del(ctx, m.Key)
	})

	bus.Register(func(ctx context.Context, m *storage.Keys) (err error) {
		m.Result, err = s.Keys(ctx, m.Prefix)
		return err
	})

	bus.Register(func(ctx context.Context, m *storage.Exists) (err error) {
		m.Result, err = s.Exists(ctx, m.Key)
		return err
	})
}
