package sendLink

import (
	"net/http"

	"imageresizerservice/app/ctx/appContext"
	"imageresizerservice/app/users/login/sendLink/sendLinkAction"
	"imageresizerservice/app/users/login/sendLink/sendLinkPage"
	"imageresizerservice/app/users/login/sendLink/sendLinkSuccessPage"
)

func Router(mux *http.ServeMux, appCtx *appContext.AppCtx) {
	sendLinkPage.Router(mux)
	sendLinkAction.Router(mux, appCtx)
	sendLinkSuccessPage.Router(mux, appCtx)
}
