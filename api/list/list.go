package list

import (
	"github.com/moonrhythm/archer/api"
)

// List spec
type List []*api.Resource

func init() {
	api.Register("List", &List{})
}
