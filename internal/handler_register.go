package internal

import (
	"net/http"
)

func RegisterHandlers(routes Routes, mux *http.ServeMux) {
	for route, handler := range routes.RoutesAndHandlers {
		pattern := "/" + routes.Root + route
		mux.HandleFunc(pattern, handler)
	}
}
