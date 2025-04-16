package projects

import (
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/projects/projectCreate"
	"net/http"
)

func Router(mux *http.ServeMux, appCtx *appCtx.AppCtx) {
	projectCreate.Router(mux, appCtx)
}
