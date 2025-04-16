package login

import (
	"net/http"

	"imageresizerservice/app/ctx/appContext"
	"imageresizerservice/app/users/login/sendLink/sendLink"
	"imageresizerservice/app/users/login/useLink/useLink"
)

func Router(mux *http.ServeMux, ac *appContext.AppCtx) {
	useLink.Router(mux, ac)
}

func RouterLoggedOut(mux *http.ServeMux, ac *appContext.AppCtx) {
	sendLink.Router(mux, ac)
	useLink.Router(mux, ac)
}
