package useLink

import (
	"net/http"

	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/users/loginWithEmailLink/useLink/useLinkAction"
	"imageresizerservice/app/users/loginWithEmailLink/useLink/useLinkPage"
	"imageresizerservice/app/users/loginWithEmailLink/useLink/useLinkSuccessPage"
)

func Router(mux *http.ServeMux, appCtx *appCtx.AppCtx) {
	useLinkPage.Router(mux)
	useLinkAction.Router(mux, appCtx)
	useLinkSuccessPage.Router(mux)
}
