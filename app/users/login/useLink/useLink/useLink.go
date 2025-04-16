package useLink

import (
	"net/http"

	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/users/login/useLink/useLinkAction"
	"imageresizerservice/app/users/login/useLink/useLinkPage"
	"imageresizerservice/app/users/login/useLink/useLinkSuccessPage"
)

func Router(mux *http.ServeMux, appCtx *appCtx.AppCtx) {
	useLinkPage.Router(mux)
	useLinkAction.Router(mux, appCtx)
	useLinkSuccessPage.Router(mux)
}
