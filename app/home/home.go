package home

import (
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/home/getHome"

	"net/http"
)

func Router(mux *http.ServeMux, ac *appCtx.AppCtx) {
	getHome.Router(mux, ac)
}
