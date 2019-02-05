package core

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"reflect"

	"github.com/moonrhythm/archer/core/namespace"
	"github.com/moonrhythm/archer/core/storage"
)

// Errors
var (
	ErrNotSupport = errors.New("not support")
)

// Controller is the resource controller
type Controller interface {
	Key(ctx context.Context, name string) string
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

// Resource converts object to resource
func (obj *Object) Resource(ctx context.Context, kind string, ns string) *Resource {
	return &Resource{
		Kind: kind,
		Metadata: Metadata{
			Name:      Name(obj.Name),
			Namespace: Name(ns),
		},
		Spec: obj.Spec,
	}
}

var registry = make(map[string]*RegistryItem)

// Register registers kind
func Register(it RegistryItem) {
	it.specType = reflect.TypeOf(it.Spec)
	registry[it.Kind] = &it
}

// RegistryItem type
type RegistryItem struct {
	Kind      string
	Ctrl      Controller
	Spec      interface{}
	Namespace bool

	specType reflect.Type
}

func (it *RegistryItem) handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// normalize
		if r.URL.Path == "" {
			r.URL.Path = "/"
		}

		// list
		if r.URL.Path == "/" && r.Method == http.MethodGet || r.Method == http.MethodHead {
			ctx := r.Context()
			objs, err := it.Ctrl.List(ctx)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			var ns string
			if it.Namespace {
				ns = namespace.Get(ctx)
			}

			resources := make([]*Resource, 0, len(objs))
			for _, obj := range objs {
				resources = append(resources, obj.Resource(ctx, it.Kind, ns))
			}

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(resources)
			return
		}

		// create
		if r.URL.Path == "/" && r.Method == http.MethodPost {
			var obj Object
			err := json.NewDecoder(r.Body).Decode(&obj)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			ctx := r.Context()

			// ensure object not exists
			{
				key := it.Ctrl.Key(ctx, obj.Name)
				if key != "" {
					q := storage.Exists{Key: key}
					err := Dispatch(ctx, &q)
					if err != nil {
						http.Error(w, err.Error(), http.StatusBadRequest)
						return
					}
					if q.Result {
						http.Error(w, "object exists", http.StatusBadRequest)
						return
					}
				}
			}

			err = it.Ctrl.Create(ctx, &obj)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}

		// update
		if r.Method == http.MethodPut {
			name := Name(r.URL.Path[1:])
			if err := name.Valid(); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			var obj Object
			err := json.NewDecoder(r.Body).Decode(&obj)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			ctx := r.Context()
			err = it.Ctrl.Update(ctx, &obj)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}

		// get
		if r.Method == http.MethodGet {
			name := Name(r.URL.Path[1:])
			if err := name.Valid(); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			ctx := r.Context()
			obj, err := it.Ctrl.Get(ctx, name.String())
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			var ns string
			if it.Namespace {
				ns = namespace.Get(ctx)
			}

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(obj.Resource(ctx, it.Kind, ns))
			return
		}

		// delete
		if r.Method == http.MethodDelete {
			name := Name(r.URL.Path[1:])
			if err := name.Valid(); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			ctx := r.Context()
			err := it.Ctrl.Delete(ctx, name.String())
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}

		http.NotFound(w, r)
	})
}
