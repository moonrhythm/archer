package core

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
)

// Errors
var (
	ErrNotSupport = errors.New("not support")
)

// Controller is the resource controller
type Controller interface {
	Create(ctx context.Context, obj *Object) error
	Update(ctx context.Context, obj *Object) error
	Get(ctx context.Context, name string) (*Object, error)
	List(ctx context.Context) ([]*Object, error)
	Delete(ctx context.Context, name string) error
}

// Object type
type Object struct {
	Name string
	Spec interface{}
}

type registryItem struct {
	Ctrl Controller
	Spec reflect.Type
}

var registry = make(map[string]*registryItem)

// Register registers kind
func Register(kind string, c Controller, d interface{}) {
	registry[kind] = &registryItem{c, reflect.TypeOf(d)}
}
