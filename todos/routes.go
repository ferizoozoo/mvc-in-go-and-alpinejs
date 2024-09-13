package todos

import (
	"net/http"

	"github.com/ferizoozoo/sake/internal"
)

var TodosRoutes = internal.Routes{
	Root: "todos",
	RoutesAndHandlers: map[string]http.HandlerFunc{
		"/":    Home,
		"/all": All,
	},
}
