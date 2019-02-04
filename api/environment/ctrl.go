package environment

import (
	"context"

	"github.com/moonrhythm/archer/core"
)

type ctrl struct{}

func (ctrl) Create(ctx context.Context, obj *core.Object) error {
	panic("implement me")
}

func (ctrl) Update(ctx context.Context, obj *core.Object) error {
	panic("implement me")
}

func (ctrl) Get(ctx context.Context, obj *core.Object) error {
	panic("implement me")
}

func (ctrl) Delete(ctx context.Context, name string) error {
	panic("implement me")
}
