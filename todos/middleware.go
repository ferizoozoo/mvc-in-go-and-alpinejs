package todos

import (
	"fmt"
	"net/http"
)

func TodosMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Todos Middleware")
		next.ServeHTTP(w, r)
	})
}
