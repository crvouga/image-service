package login

import (
	"net/http"

	"imageresizerservice/app/ctx/appContext"
	"imageresizerservice/app/users/login/sendLink/sendLink"
	"imageresizerservice/app/users/login/useLink/useLink"
)

func Router(mux *http.ServeMux, appCtx *appContext.AppCtx) {
	useLink.Router(mux, appCtx)
}

func RouterLoggedOut(mux *http.ServeMux, appCtx *appContext.AppCtx) {
	sendLink.Router(mux, appCtx)
	useLink.Router(mux, appCtx)
}
