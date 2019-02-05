package request

import (
	"context"

	"github.com/moonrhythm/archer/core"
	"github.com/moonrhythm/archer/core/namespace"
)

type ctrl struct{}

func (c ctrl) Key(ctx context.Context, name string) string {
	return "namespaces/" + namespace.Get(ctx) + "/request/" + name
}

func (ctrl) Create(ctx context.Context, obj *core.Object) error {
	panic("implement me")
}

func (ctrl) Update(ctx context.Context, obj *core.Object) error {
	panic("implement me")
}

func (ctrl) Get(ctx context.Context, name string) (*core.Object, error) {
	panic("implement me")
}

func (ctrl) List(ctx context.Context) ([]*core.Object, error) {
	panic("implement me")
}

func (ctrl) Delete(ctx context.Context, name string) error {
	panic("implement me")
}
