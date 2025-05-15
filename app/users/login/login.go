package login

import (
	"net/http"

	"imageService/app/ctx/appCtx"
	"imageService/app/users/login/sendLink"
	"imageService/app/users/login/useLink"
)

func Router(mux *http.ServeMux, ac *appCtx.AppCtx) {
	useLink.Router(mux, ac)
}

func RouterLoggedOut(mux *http.ServeMux, ac *appCtx.AppCtx) {
	sendLink.Router(mux, ac)
	useLink.Router(mux, ac)
}
