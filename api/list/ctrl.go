package list

import (
	"context"

	"github.com/moonrhythm/archer/core"
)

type ctrl struct{}

func (c ctrl) Key(ctx context.Context, name string) string {
	return ""
}

func (ctrl) Create(ctx context.Context, obj *core.Object) error {
	spec := obj.Spec.(*List)

	var ops []interface{}
	for _, m := range *spec {
		ops = append(ops, &core.Create{
			Resource: m,
		})
	}

	return core.Dispatch(ctx, ops...)
}

func (ctrl) Update(ctx context.Context, obj *core.Object) error {
	spec := obj.Spec.(*List)

	var ops []interface{}
	for _, m := range *spec {
		ops = append(ops, &core.Update{
			Resource: m,
		})
	}

	return core.Dispatch(ctx, ops...)
}

func (ctrl) Get(ctx context.Context, name string) (*core.Object, error) {
	return nil, core.ErrNotSupport
}

func (ctrl) List(ctx context.Context) ([]*core.Object, error) {
	return nil, core.ErrNotSupport
}

func (ctrl) Delete(ctx context.Context, name string) error {
	return core.ErrNotSupport
}
