package server

import (
	"fmt"
	"net/http"
)

func RegisterEndpoints() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello you've requested: %s\n", r.URL.Path)
	})

	fs := http.FileServer(http.Dir("server/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}

func StartServer(port string) error {
	if port == "" {
		port = "8080"
	}
	return http.ListenAndServe(":"+port, nil)
}
