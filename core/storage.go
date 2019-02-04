package core

import (
	"context"

	"github.com/pkg/errors"
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
