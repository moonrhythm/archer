package request

import (
	"fmt"

	"github.com/moonrhythm/archer/api"
)

// Request type
type Request struct {
	// one of
	Project string `json:"project"`
	Group   string `json:"group"`

	Method  string   `json:"method"`
	URL     string   `json:"url"`
	Headers []Header `json:"headers"`
	Body    string   `json:"body"`
}

func init() {
	api.Register("Request", &Request{})
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
