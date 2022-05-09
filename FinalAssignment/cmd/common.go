package cmd

import (
	_ "embed"
	"net/http"

	"github.com/flowchartsman/swaggerui"
)

//go:embed swagger.json
var spec []byte

func CreateCommonMux(h http.Handler) http.Handler {
	r := http.NewServeMux()
	fs := http.FileServer(http.Dir("./client/build/static"))
	r.Handle("/swagger/", http.StripPrefix("/swagger", swaggerui.Handler(spec)))
	r.HandleFunc("/app", index)
	r.HandleFunc("api/lists", index)
	r.HandleFunc("api/tasks/", index)
	r.HandleFunc("api/tasks", index)
	r.HandleFunc("api/lists/{id}/tasks", index)
	r.Handle("/static/", http.StripPrefix("/static/", fs))

	r.Handle("/api", h)
	r.Handle("/api/lists", h)
	r.Handle("/api/tasks/", h)
	r.Handle("/api/tasks", h)
	r.Handle("/api/lists/{id}/tasks", h)

	return r
}

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./client/build/index.html")
}
