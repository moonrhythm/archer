package core

// Create creates new resource, error if resource already exists
type Create struct {
	*Resource
}

// Apply creates or updates if resource already exists
type Apply struct {
	*Resource
}

// Update updates resource
type Update struct {
	*Resource
}

// Get get a resource
type Get struct {
	Name Name `json:"name"`

	Result *Resource `json:"-"`
}

// List lists resources
type List struct{}

// Delete deletes a resource
type Delete struct {
	Kind string `json:"kind"`
	Name Name   `json:"name"`
}
