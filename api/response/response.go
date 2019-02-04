package response

import (
	"github.com/moonrhythm/archer/api"
)

// Response type
type Response struct {
	Data string `json:"data"`
}

func init() {
	api.Register("Response", &Response{})
}
