package environment

import (
	"fmt"

	"github.com/moonrhythm/archer/api"
)

// Environment spec
type Environment struct {
	Project string            `json:"project"`
	Data    map[string]string `json:"data"`
}

func init() {
	api.Register("Environment", &Environment{})
}

// Valid checks is env spec valid
func (e *Environment) Valid() error {
	if e.Project == "" {
		return fmt.Errorf("requires project")
	}
	return nil
}
