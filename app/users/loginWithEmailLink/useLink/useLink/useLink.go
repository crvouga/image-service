package useLink

import (
	"net/http"

	"imageresizerservice/app/ctx"
	"imageresizerservice/app/users/loginWithEmailLink/useLink/useLinkAction"
	"imageresizerservice/app/users/loginWithEmailLink/useLink/useLinkPage"
	"imageresizerservice/app/users/loginWithEmailLink/useLink/useLinkSuccessPage"
)

func Router(mux *http.ServeMux, appCtx *ctx.AppCtx) {
	useLinkPage.Router(mux)
	useLinkAction.Router(mux, appCtx)
	useLinkSuccessPage.Router(mux)
}
