package api

import (
	"imageresizerservice/app/ctx/appContext"
	"net/http"
)

func Router(mux *http.ServeMux, appCtx *appContext.AppCtx) {
	mux.HandleFunc("/api/image/resize", ApiImageResize(appCtx))
}
