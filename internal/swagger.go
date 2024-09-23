package internal

import "net/http"

type SwaggerConfig struct {
	Host string
}

func NewSwaggerConfig(host string) *SwaggerConfig {
	return &SwaggerConfig{
		Host: host,
	}
}

type SwaggerHandler interface {
	Register(mux *http.ServeMux)
}
