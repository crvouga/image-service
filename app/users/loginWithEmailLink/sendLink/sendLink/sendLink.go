package sendLink

import (
	"net/http"

	"imageresizerservice/app/ctx"
	"imageresizerservice/app/users/loginWithEmailLink/sendLink/sendLinkAction"
	"imageresizerservice/app/users/loginWithEmailLink/sendLink/sendLinkPage"
	"imageresizerservice/app/users/loginWithEmailLink/sendLink/sendLinkSuccessPage"
)

func Router(mux *http.ServeMux, d *ctx.Ctx) {
	sendLinkPage.Router(mux)
	sendLinkAction.Router(mux, d)
	sendLinkSuccessPage.Router(mux)
}
