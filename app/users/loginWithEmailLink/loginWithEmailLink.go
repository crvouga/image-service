package loginWithEmailLink

import (
	"net/http"

	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/users/loginWithEmailLink/sendLink/sendLink"
	"imageresizerservice/app/users/loginWithEmailLink/useLink/useLink"
)

func Router(mux *http.ServeMux, appCtx *appCtx.AppCtx) {
	sendLink.Router(mux, appCtx)
	useLink.Router(mux, appCtx)
}
