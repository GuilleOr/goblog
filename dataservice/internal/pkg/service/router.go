package service

import (
	"github.com/callistaenterprise/goblog/common/monitoring"
	"github.com/callistaenterprise/goblog/common/router"
	"net/http"

	"github.com/callistaenterprise/goblog/common/tracing"
	"github.com/gorilla/mux"
)

// NewRouter creates a mux.Router pointer.
func NewRouter(serviceName string) *mux.Router {

	muxRouter := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		summaryVec := monitoring.BuildSummaryVec(serviceName, route.Name, route.Method+" "+route.Pattern)

		// Add route to muxRouter, including middleware chaining and passing the summaryVec to the WithMonitoring func.
		muxRouter.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(monitoring.WithMonitoring(withTracing(route.HandlerFunc, route), route, summaryVec))

	}
	return muxRouter
}

func withTracing(next http.Handler, route router.Route) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		span := tracing.StartHTTPTrace(req, route.Name)
		defer span.Finish()

		ctx := tracing.UpdateContext(req.Context(), span)
		next.ServeHTTP(rw, req.WithContext(ctx))
	})
}
