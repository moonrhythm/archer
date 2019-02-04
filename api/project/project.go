package project

import (
	"github.com/moonrhythm/archer/api"
)

// Project spec
type Project struct{}

func init() {
	api.Register("Project", &Project{})
}

// Valid checks is project spec valid
func (p *Project) Valid() error {
	return nil
}
