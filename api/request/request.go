package request

import (
	"fmt"

	"github.com/moonrhythm/archer/core"
)

// Request type
type Request struct {
	// one of
	Project string `json:"project"`
	Group   string `json:"group"`

	Description string   `json:"description"`
	Method      string   `json:"method"`
	URL         string   `json:"url"`
	Headers     []Header `json:"headers"`
	Body        string   `json:"body"`
}

func init() {
	core.Register("Request", &ctrl{}, &Request{})
}

// Valid checks is request spec valid
func (r *Request) Valid() error {
	if r.Project == "" && r.Group == "" {
		return fmt.Errorf("requires one of project or group")
	}
	if r.Project != "" && r.Group != "" {
		return fmt.Errorf("only allow one of project or group")
	}
	return nil
}
