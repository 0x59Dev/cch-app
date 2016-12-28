package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/methlab669/cch-app/app/handlers"
)

type RouteConf struct {
	Name    string
	Method  string
	Path    string
	Handler http.HandlerFunc
}

//Routes is an slice of Routes
type Routes []RouteConf

//Name, method, path, handler function, add every URL here
var routerPaths = Routes{
	RouteConf{
		"register",
		"POST",
		"/register",
		handlers.Register,
	},
}

//NewRouter Initializes a Router used for some wise routing used in ListenAndServe
func NewRouter() *mux.Router {
	fmt.Println("Initializing new router")
	//initializes new Gorilla router , now handlers entry point is our r
	r := mux.NewRouter()
	//Ranges over routerPaths specified in slice and registers them with field specified in struct
	for _, route := range routerPaths {
		handler := route.Handler
		log.Println("Registered : ", route.Name, route.Path)

		r.Methods(route.Method).Path(route.Path).Name(route.Name).Handler(handler)
	}
	return r

}
