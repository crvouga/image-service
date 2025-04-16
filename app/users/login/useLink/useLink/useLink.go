package useLink

import (
	"net/http"

	"imageresizerservice/app/ctx/appContext"
	"imageresizerservice/app/users/login/useLink/useLinkAction"
	"imageresizerservice/app/users/login/useLink/useLinkPage"
	"imageresizerservice/app/users/login/useLink/useLinkSuccessPage"
)

func Router(mux *http.ServeMux, ac *appContext.AppCtx) {
	useLinkPage.Router(mux)
	useLinkAction.Router(mux, ac)
	useLinkSuccessPage.Router(mux)
}
