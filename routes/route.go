package route

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"

	"github.com/JamesnaW/go-boilerplate/middlewares"
	"github.com/JamesnaW/go-boilerplate/models"
)

var routes model.Routes = concatRoute(
	Todo(),
)

func NewFastRouter() *fasthttprouter.Router {
	router := fasthttprouter.New()

	for _, route := range routes {
		var handler fasthttp.RequestHandler
		handler = route.HandlerFunc
		handler = middleware.LoggerFastHttp(handler, route.Name)

		router.Handle(route.Method, route.Pattern, handler)
	}

	return router
}

func concatRoute(r ...model.Routes) model.Routes {
	var routes model.Routes
	for _, arr := range r {
		for i := 0; i < len(arr); i++ {
			routes = append(routes, arr[i])
		}
	}

	return routes
}
