package core

import (
	"context"

	"github.com/moonrhythm/dispatcher"
)

var bus = dispatcher.NewMux()

// Dispatch dispatches operators
func Dispatch(ctx context.Context, op ...interface{}) error {
	for _, m := range op {
		err := bus.Dispatch(ctx, m)
		if err != nil {
			return err
		}
	}
	return nil
}
