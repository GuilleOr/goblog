package service

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter from http://thenewstack.io/make-a-restful-json-api-go/
func NewRouter(serviceName string) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
