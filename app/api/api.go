package api

import (
	"imageService/app/ctx/appCtx"
	"net/http"
)

const (
	EndpointApiImageResize = "/api/image/resize"
)

func Router(mux *http.ServeMux, ac *appCtx.AppCtx) {
	mux.HandleFunc(EndpointApiImageResize, ApiImageResize(ac))
}
