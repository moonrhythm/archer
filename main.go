package main

import (
	"log"
	"net/http"

	// inject apis
	_ "github.com/moonrhythm/archer/api/environment"
	_ "github.com/moonrhythm/archer/api/group"
	_ "github.com/moonrhythm/archer/api/list"
	_ "github.com/moonrhythm/archer/api/namespace"
	_ "github.com/moonrhythm/archer/api/request"

	"github.com/moonrhythm/archer/core"
	"github.com/moonrhythm/archer/store/postgresql"
)

func main() {
	// TODO: accept db strategy from flag
	storage, err := postgresql.New("postgres://localhost/archer?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	core.SetStorage(storage)

	mux := http.NewServeMux()
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", core.Handler()))

	http.ListenAndServe(":8080", mux)
}
