package home

import (
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/home/homePage"

	"net/http"
)

func Router(mux *http.ServeMux, ac *appCtx.AppCtx) {
	homePage.Router(mux, ac)
}
