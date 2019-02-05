package namespace

import (
	"context"
	"net/http"
	"strings"
)

type contextKey struct{}

// NewContext creates new context with namespace
func NewContext(parent context.Context, ns string) context.Context {
	return context.WithValue(parent, contextKey{}, ns)
}

const defaultNs = "default"

// Get gets namespace from context
func Get(ctx context.Context) string {
	ns, _ := ctx.Value(contextKey{}).(string)
	if ns == "" {
		ns = defaultNs
	}
	return ns
}

// With returns namespace with resource name
func With(ctx context.Context, name string) string {
	return Get(ctx) + "/" + name
}

// Split splits namespace and name
func Split(s string) (ns string, name string) {
	xs := strings.Split(s, "/")
	switch len(xs) {
	case 1:
		return xs[0], ""
	case 2:
		return xs[0], xs[1]
	default:
		return "", ""
	}
}

// Middleware strips prefix path as namespace
func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		index := strings.Index(r.URL.Path[1:], "/")
		r = r.WithContext(NewContext(r.Context(), r.URL.Path[:index+1]))
		r.URL.Path = r.URL.Path[index+1:]
		h.ServeHTTP(w, r)
	})
}
