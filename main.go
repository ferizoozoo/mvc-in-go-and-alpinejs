package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/ferizoozoo/sake/internal"
	"github.com/ferizoozoo/sake/todos"
)

const SERVER_PORT = 4444
const keyServerAddr = "serverAddr"

func main() {
	mux := http.NewServeMux()

	// Controllers
	internal.RegisterHandlers(todos.TodosRoutes, mux)

	// Middlewares
	wrappedMux := internal.NewMiddlewareRegistrar().
		Add(todos.TodosMiddleware).
		Register(mux)

	internal.LoadEnvironmentVariables()
	//internal.GenerateSwaggerFiles()
	internal.ServeStaticFiles(mux)

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
