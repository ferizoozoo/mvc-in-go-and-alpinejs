package todos

import (
	"alpine-golang-test/backend/internal"
	"net/http"
)

var TodosRoutes = internal.Routes{
	Root: "todos",
	RoutesAndHandlers: map[string]http.HandlerFunc{
		"/":    Home,
		"/all": All,
	},
}
