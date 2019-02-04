package namespace

import (
	"context"
)

type contextKey struct{}

// NewContext creates new context with namespace
func NewContext(parent context.Context, ns string) context.Context {
	return context.WithValue(parent, contextKey{}, ns)
}

// Get gets namespace from context
func Get(ctx context.Context) string {
	ns, _ := ctx.Value(contextKey{}).(string)
	return ns
}
