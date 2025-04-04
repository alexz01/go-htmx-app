package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterEndpoints() http.Handler {

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello you've requested: %s\n", r.URL.Path)
	})

	router.HandleFunc("/page/{name}/{id}", func(w http.ResponseWriter, r *http.Request) {
		routeVars := mux.Vars(r)

		fmt.Fprintf(w, "Hello page requested: %s, id: %s\n", routeVars["name"], routeVars["id"])
	})

	fs := http.FileServer(http.Dir("server/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	return router
}

func StartServer(handler http.Handler, port string) error {
	if port == "" {
		port = "8080"
	}
	return http.ListenAndServe(":"+port, handler)
}
