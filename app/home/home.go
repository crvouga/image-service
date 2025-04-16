package home

import (
	"imageresizerservice/app/ctx/appContext"
	"imageresizerservice/app/home/getHome"

	"net/http"
)

func Router(mux *http.ServeMux, ac *appContext.AppCtx) {
	getHome.Router(mux, ac)
}
