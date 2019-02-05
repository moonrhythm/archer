package namespace

import (
	"context"

	"github.com/moonrhythm/archer/core"
	"github.com/moonrhythm/archer/core/namespace"
	"github.com/moonrhythm/archer/core/storage"
)

type ctrl struct{}

func (c ctrl) Key(ctx context.Context, name string) string {
	return "namespace/" + name
}

func (c ctrl) Create(ctx context.Context, obj *core.Object) error {
	return core.Dispatch(ctx, &storage.Set{
		Key:   c.Key(ctx, obj.Name),
		Value: []byte{},
	})
}

func (c ctrl) Update(ctx context.Context, obj *core.Object) error {
	return core.ErrNotSupport
}

func (c ctrl) Get(ctx context.Context, name string) (*core.Object, error) {
	q := storage.Get{Key: c.Key(ctx, name)}
	err := core.Dispatch(ctx, &q)
	if err != nil {
		return nil, err
	}

	return &core.Object{
		Name: name,
		Spec: &Namespace{},
	}, nil
}

func (c ctrl) List(ctx context.Context) ([]*core.Object, error) {
	var keys []string
	{
		q := storage.Keys{Prefix: c.Key(ctx, "")}
		err := core.Dispatch(ctx, &q)
		if err != nil {
			return nil, err
		}
		keys = q.Result
	}

	objs := make([]*core.Object, 0, len(keys))
	for _, key := range keys {
		q := storage.Get{Key: key}
		err := core.Dispatch(ctx, &q)
		if err != nil {
			return nil, err
		}
		_, name := namespace.Split(key)
		objs = append(objs, &core.Object{
			Name: name,
			Spec: &Namespace{},
		})
	}

	return objs, nil
}

func (c ctrl) Delete(ctx context.Context, name string) error {
	return core.Dispatch(ctx, &storage.Del{
		Key: c.Key(ctx, name),
	})
}
