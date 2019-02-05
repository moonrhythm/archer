package namespace

import (
	"github.com/moonrhythm/archer/core"
)

// Namespace spec
type Namespace struct{}

// Kind value
const Kind = "Namespace"

func init() {
	core.Register(core.RegistryItem{
		Kind: Kind,
		Ctrl: &ctrl{},
		Spec: &Namespace{},
	})
}
