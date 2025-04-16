package login

import (
	"net/http"

	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/users/login/sendLink/sendLink"
	"imageresizerservice/app/users/login/useLink/useLink"
)

func Router(mux *http.ServeMux, appCtx *appCtx.AppCtx) {
	sendLink.Router(mux, appCtx)
	useLink.Router(mux, appCtx)
}
