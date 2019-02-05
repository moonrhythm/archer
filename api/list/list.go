package list

import (
	"github.com/moonrhythm/archer/core"
)

// List spec
type List []*core.Resource

const (
	Kind = "List"
)

func init() {
	core.Register(core.RegistryItem{
		Kind:      Kind,
		Ctrl:      &ctrl{},
		Spec:      &List{},
		Namespace: true, // TODO: maybe false ?
	})
}
