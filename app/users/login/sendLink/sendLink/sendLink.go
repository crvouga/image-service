package sendLink

import (
	"net/http"

	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/users/login/sendLink/sendLinkAction"
	"imageresizerservice/app/users/login/sendLink/sendLinkPage"
	"imageresizerservice/app/users/login/sendLink/sendLinkSuccessPage"
)

func Router(mux *http.ServeMux, appCtx *appCtx.AppCtx) {
	sendLinkPage.Router(mux)
	sendLinkAction.Router(mux, appCtx)
	sendLinkSuccessPage.Router(mux, appCtx)
}
