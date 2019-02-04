package environment

import (
	"context"
	"encoding/json"

	"github.com/moonrhythm/archer/core"
	"github.com/moonrhythm/archer/core/namespace"
	"github.com/moonrhythm/archer/core/storage"
)

type ctrl struct{}

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
	ns := namespace.Get(ctx)
	q := storage.Keys{Prefix: ns}
	err := core.Dispatch(ctx, &q)
	if err != nil {
		return nil, err
	}

	ops := make([]interface{}, 0, len(q.Result))
	for _, key := range q.Result {
		ops = append(ops, &storage.Get{Key: key})
	}
	err = core.Dispatch(ctx, ops...)
	if err != nil {
		return nil, err
	}

	objs := make([]*core.Object, 0, len(q.Result))
	for i, key := range q.Result {
		_, name := namespace.Split(key)
		var spec Environment
		err = json.Unmarshal(ops[i].(*storage.Get).Result, &spec)
		if err != nil {
			continue
		}
		objs = append(objs, &core.Object{
			Name: name,
			Spec: &spec,
		})
	}

	return objs, nil
}

func (ctrl) Delete(ctx context.Context, name string) error {
	key := namespace.With(ctx, name)
	{
		q := storage.Exists{Key: key}
		err := core.Dispatch(ctx, &q)
		if err != nil {
			return err
		}
		if !q.Result {
			return core.ErrNotFound
		}
	}
	return core.Dispatch(ctx, &storage.Del{Key: key})
}
