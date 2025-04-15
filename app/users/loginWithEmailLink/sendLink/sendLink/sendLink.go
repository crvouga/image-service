package sendLink

import (
	"net/http"

	"imageresizerservice/app/deps"
	"imageresizerservice/app/users/loginWithEmailLink/sendLink/sendLinkAction"
	"imageresizerservice/app/users/loginWithEmailLink/sendLink/sendLinkPage"
	"imageresizerservice/app/users/loginWithEmailLink/sendLink/sendLinkSuccessPage"
)

func Router(mux *http.ServeMux, d *deps.Deps) {
	sendLinkPage.Router(mux)
	sendLinkAction.Router(mux, d)
	sendLinkSuccessPage.Router(mux)
}
