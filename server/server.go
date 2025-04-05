package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var htmlStr = `
<!DOCTYPE HTML>
<html>
<head>
	<title>My Go App</title>
</head>
<body>
	%s
</body>
</html>
`

var router *mux.Router

func GetRouter() *mux.Router {
	if router == nil {
		router = mux.NewRouter()
	}
	return router
}

/*
 */
func RegisterEndpoints(router *mux.Router) http.Handler {

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		htmlContent := `
		<h2>Hello world</h2>
		<p>requested url: %s</p>
		`
		htmlContent = fmt.Sprintf(htmlStr, htmlContent)
		fmt.Fprintf(w, htmlContent, r.URL.Path)
	}).Methods(("GET"))

	router.HandleFunc("/page/{name}/{id}", func(w http.ResponseWriter, r *http.Request) {
		routeVars := mux.Vars(r)
		htmlContent := `Hello page requested: %s, id: %s`

		htmlContent = fmt.Sprintf(htmlStr, htmlContent)
		fmt.Fprintf(w, htmlContent, routeVars["name"], routeVars["id"])
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
