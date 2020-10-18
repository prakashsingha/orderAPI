package routers

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/prakashsingha/orderAPI/controllers"
	"github.com/prakashsingha/orderAPI/middleware"
)

// Route structure
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is a collection of Route
type Routes []Route

var routes = Routes{
	Route{
		"CreateOrder",
		strings.ToUpper("Post"),
		"/orders",
		controllers.CreateOrder,
	},

	Route{
		"GetOrders",
		strings.ToUpper("Get"),
		"/orders",
		controllers.GetOrders,
	},

	// Payment
	Route{
		"GetPayment",
		strings.ToUpper("Get"),
		"/payments",
		controllers.GetPayment,
	},
	Route{
		"MakePayment",
		strings.ToUpper("Patch"),
		"/payments",
		controllers.MakePayment,
	},
}

// NewRouter returns the new router instance with all routes
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = middleware.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
