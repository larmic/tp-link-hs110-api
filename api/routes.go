package api

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		Name:        "Open Api documentation (yaml)",
		Method:      "GET",
		Pattern:     "/",
		HandlerFunc: OpenApiDocumentation,
	},
	Route{
		Name:        "Open Api documentation (yaml)",
		Method:      "GET",
		Pattern:     "/api",
		HandlerFunc: OpenApiDocumentation,
	},
	Route{
		Name:        "System",
		Method:      "GET",
		Pattern:     "/api/{ip}",
		HandlerFunc: GetSystemHandler,
	},
	Route{
		Name:        "Consumption",
		Method:      "GET",
		Pattern:     "/api/{ip}/consumption",
		HandlerFunc: ConsumptionHandler,
	},
}
