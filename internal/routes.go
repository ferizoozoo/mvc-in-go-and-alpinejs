package internal

import "net/http"

type Routes struct {
	Root              string
	RoutesAndHandlers map[string]http.HandlerFunc
}
