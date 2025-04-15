package sendLink

import (
	"net/http"

	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/users/loginWithEmailLink/sendLink/sendLinkAction"
	"imageresizerservice/app/users/loginWithEmailLink/sendLink/sendLinkPage"
	"imageresizerservice/app/users/loginWithEmailLink/sendLink/sendLinkSuccessPage"
)

func Router(mux *http.ServeMux, appCtx *appCtx.AppCtx) {
	sendLinkPage.Router(mux)
	sendLinkAction.Router(mux, appCtx)
	sendLinkSuccessPage.Router(mux)
}
