package group

import (
	"fmt"

	"github.com/moonrhythm/archer/api"
)

// Group spec
//
// Group must contain one of project or group
type Group struct {
	Project string `json:"project"`
	Group   string `json:"group"`
}

func init() {
	api.Register("Group", &Group{})
}

// Valid checks is group spec valid
func (g *Group) Valid() error {
	if g.Project == "" && g.Group == "" {
		return fmt.Errorf("requires one of project or group")
	}
	if g.Project != "" && g.Group != "" {
		return fmt.Errorf("only allow one of project or group")
	}
	return nil
}
