package internal

import "net/http"

type Middleware func(http.Handler) http.Handler

type MiddlewareRegistrar struct {
	middlewares []Middleware
}

func NewMiddlewareRegistrar() *MiddlewareRegistrar {
	return &MiddlewareRegistrar{
		[]Middleware{},
	}
}

func (mr *MiddlewareRegistrar) Add(middleware Middleware) *MiddlewareRegistrar {
	mr.middlewares = append(mr.middlewares, middleware)
	return mr
}

func (mr *MiddlewareRegistrar) Register(handler http.Handler) http.Handler {
	for _, middleware := range mr.middlewares {
		handler = middleware(handler)
	}

	return handler
}
