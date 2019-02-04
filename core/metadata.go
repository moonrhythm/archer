package core

import (
	"fmt"
	"regexp"
)

// Metadata is the resource's metadata
type Metadata struct {
	Name      Name `json:"name"`
	Namespace Name `json:"namespace"`
}

// Valid checks is metadata valid
func (m *Metadata) Valid() error {
	if err := m.Name.Valid(); err != nil {
		return fmt.Errorf("name; %s", err)
	}
	if err := m.Namespace.Valid(); err != nil {
		return fmt.Errorf("namespace; %s", err)
	}
	return nil
}

// Name is the resource name
type Name string

var reName = regexp.MustCompile(`^[a-z0-9\-.]+$`)

func (n Name) Valid() error {
	if n == "" {
		return fmt.Errorf("required")
	}
	if len(n) > 255 {
		return fmt.Errorf("length allow only 255 bytes")
	}
	if !reName.MatchString(string(n)) {
		return fmt.Errorf("invalid format")
	}
	return nil
}
