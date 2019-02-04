package api

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Resource is the api resource
type Resource struct {
	Kind     string      `json:"kind"`
	Metadata Metadata    `json:"metadata"`
	Spec     interface{} `json:"spec"`
}

type validator interface {
	Valid() error
}

var resourceKinds = make(map[string]reflect.Type)

// Register registers kind
func Register(kind string, d interface{}) {
	resourceKinds[kind] = reflect.TypeOf(d)
}

// UnmarshalJSON implements json.Unmarshaler
func (r *Resource) UnmarshalJSON(b []byte) error {
	var d struct {
		Kind string `json:"kind"`
	}
	if err := json.Unmarshal(b, &d); err != nil {
		return err
	}

	var p struct {
		Kind     string      `json:"kind"`
		Metadata Metadata    `json:"metadata"`
		Spec     interface{} `json:"spec"`
	}
	if typ, ok := resourceKinds[d.Kind]; ok {
		p.Spec = reflect.New(typ)
	} else {
		return fmt.Errorf("unknown kind")
	}
	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	*r = Resource(p)

	return nil
}

// Valid checks is resource valid
func (r *Resource) Valid() error {
	if err := r.Metadata.Valid(); err != nil {
		return fmt.Errorf("metadata; %s", err)
	}

	if spec, ok := r.Spec.(validator); ok {
		if err := spec.Valid(); err != nil {
			return fmt.Errorf("spec; %s", err)
		}
	}

	return nil
}
