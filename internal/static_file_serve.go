package internal

import "net/http"

func ServeStaticFiles(mux *http.ServeMux) {
	fs := http.FileServer(http.Dir("./public"))
	mux.Handle("/", fs)
}
