package routes

import (
	"net/http"

	ctrl "github.com/alvesRenan/rest-golang-test/controllers"
)

type route struct {
	Path   string
	Method string
	Func   http.HandlerFunc
}

// Routes is a collection of route
type Routes []route

var routes = Routes{
	// scenario routes
	route{
		Path:   "/scneario/list",
		Method: "GET",
		Func:   ctrl.GetScenarios,
	},
	route{
		Path:   "/scenario/create",
		Method: "POST",
		Func:   ctrl.CreateScenario,
	},
	route{
		Path:   "/scenario/delete/{name}",
		Method: "DELETE",
		Func:   ctrl.DeleteScenario,
	},

	// container routes
	route{
		Path:   "/container/list",
		Method: "GET",
		Func:   ctrl.GetContainers,
	},
	route{
		Path:   "/container/create",
		Method: "POST",
		Func:   ctrl.CreateContainer,
	},
	route{
		Path:   "/container/delete/{name}",
		Method: "DELETE",
		Func:   ctrl.DeleteContainer,
	},
}
