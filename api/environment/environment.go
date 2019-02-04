package environment

import (
	"fmt"

	"github.com/moonrhythm/archer/core"
)

// Environment spec
type Environment struct {
	Project     string            `json:"project"`
	Description string            `json:"description"`
	Data        map[string]string `json:"data"`
}

func init() {
	core.Register("Environment", &ctrl{}, &Environment{})
}

// Valid checks is env spec valid
func (e *Environment) Valid() error {
	if e.Project == "" {
		return fmt.Errorf("requires project")
	}
	return nil
}
