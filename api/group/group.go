package group

import (
	"fmt"

	"github.com/moonrhythm/archer/core"
)

// Group spec
type Group struct {
	// one of
	Project string `json:"project"`
	Group   string `json:"group"`

	Description string `json:"description"`
}

const (
	Kind = "Group"
)

func init() {
	core.Register(core.RegistryItem{
		Kind:      Kind,
		Ctrl:      &ctrl{},
		Spec:      &Group{},
		Namespace: true,
	})
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
