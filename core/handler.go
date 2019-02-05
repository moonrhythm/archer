package core

import (
	"net/http"
	"strings"
)

// Handler returns core's handler
func Handler() http.Handler {
	mux := http.NewServeMux()

	nsMux := http.NewServeMux()
	mux.Handle("/namespaces/", http.StripPrefix("/namespaces", nsMux))

	for name, it := range registry {
		name = strings.ToLower(name)
		h := http.StripPrefix("/"+name, it.handler())

		var m *http.ServeMux
		if it.Namespace {
			m = nsMux
		} else {
			m = mux
		}
		m.Handle("/"+name, h)
		m.Handle("/"+name+"/", h)
	}

	return mux
}
