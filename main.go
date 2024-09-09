package main

import (
	"alpine-golang-test/backend/internal"
	"alpine-golang-test/backend/todos"
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
)

const SERVER_PORT = 4444
const keyServerAddr = "serverAddr"

func main() {
	mux := http.NewServeMux()
	internal.RegisterHandlers(todos.TodosRoutes, mux)
	wrappedMux := internal.NewMiddlewareRegistrar().
		Add(todos.TodosMiddleware).
		Register(mux)

	ctx := context.Background()

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", SERVER_PORT),
		Handler: wrappedMux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}
	if err := server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("error running http server: %s\n", err)
		}
	}
}
