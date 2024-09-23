package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/ferizoozoo/sake/cache"
	"github.com/ferizoozoo/sake/internal"
	"github.com/ferizoozoo/sake/swagger"
	"github.com/ferizoozoo/sake/todos"
)

const SERVER_PORT = 4444
const SERVER_HOST = ""
const keyServerAddr = "serverAddr"

func main() {
	mux := http.NewServeMux()
	server_addr := fmt.Sprintf("%s:%d", SERVER_HOST, SERVER_PORT)

	// Controllers
	internal.RegisterHandlers(todos.TodosRoutes, mux)

	// Middlewares
	wrappedMux := internal.NewMiddlewareRegistrar().
		Add(todos.TodosMiddleware).
		Register(mux)

	internal.LoadEnvironmentVariables()

	swagger.NewSwaggoHandler(&internal.SwaggerConfig{
		Host: server_addr,
	}).Register(mux)

	internal.ServeStaticFiles(mux)

	internal.GetCache().SetProvider(cache.NewRedisCacheProvider(internal.CacheOptions{
		Address:  "localhost:6379",
		Password: "enter your cache password here",
	}))

	// TODO: Add jwt support

	ctx := context.Background()

	server := http.Server{
		Addr:    server_addr,
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
