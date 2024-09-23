package swagger

import (
	"net/http"

	"github.com/ferizoozoo/sake/internal"
	httpSwagger "github.com/swaggo/http-swagger"
)

type SwaggoHandler struct {
	config *internal.SwaggerConfig
}

func NewSwaggoHandler(config *internal.SwaggerConfig) *SwaggoHandler {
	return &SwaggoHandler{
		config,
	}
}

func (sh *SwaggoHandler) Register(mux *http.ServeMux) {
	mux.Handle("/swagger/", httpSwagger.WrapHandler)
}
