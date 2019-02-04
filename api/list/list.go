package list

import (
	"github.com/moonrhythm/archer/core"
)

// List spec
type List []*core.Resource

func init() {
	core.Register("List", &ctrl{}, &List{})
}
