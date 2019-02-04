package namespace

import (
	"github.com/moonrhythm/archer/core"
)

// Namespace spec
type Namespace struct{}

func init() {
	core.Register("Namespace", &ctrl{}, &Namespace{})
}
