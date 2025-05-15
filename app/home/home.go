package home

import (
	"imageService/app/ctx/appCtx"
	"imageService/app/home/homePage"

	"net/http"
)

func Router(mux *http.ServeMux, ac *appCtx.AppCtx) {
	homePage.Router(mux, ac)
}
