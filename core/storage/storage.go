package storage

// Set operator
type Set struct {
	Key   string
	Value []byte
}

// Get operator
type Get struct {
	Key string

	Result []byte
}

// Del operator
type Del struct {
	Key string
}

// Keys operator
type Keys struct {
	Prefix string

	Result []string
}

// Exists operator
type Exists struct {
	Key string

	Result bool
}
