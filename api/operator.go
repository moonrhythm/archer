package api

// Create creates new resource, error if resource already exists
type Create struct {
	Resource
}

// Apply creates or updates if resource already exists
type Apply struct {
	Resource
}

// Update updates resource
type Update struct {
	Resource
}

// Delete deletes resource
type Delete struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
}
