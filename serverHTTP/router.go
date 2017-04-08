package serverHTTP

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/8tomat8/webAppSkeleton/environment"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

type Handlers struct {
	env *environment.Env
}

var handlers Handlers

//NewRouter create application router
func NewRouter(env *environment.Env) *mux.Router {
	handlers = Handlers{env}
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
		Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	router.PathPrefix("/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	return router
}

var routes = Routes{
	Route{"Index", "GET", "/", handlers.Index},
}
