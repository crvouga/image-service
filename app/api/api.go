package api

import (
	"imageresizerservice/app/ctx/appCtx"
	"net/http"
)

func Router(mux *http.ServeMux, appCtx *appCtx.AppCtx) {
	mux.HandleFunc("/api/image/resize", ApiImageResize(appCtx))
}
