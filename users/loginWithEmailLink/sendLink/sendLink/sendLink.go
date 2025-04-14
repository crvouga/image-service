package sendLink

import (
	"net/http"

	"imageresizerservice/deps"
	"imageresizerservice/users/loginWithEmailLink/sendLink/sendLinkAction"
	"imageresizerservice/users/loginWithEmailLink/sendLink/sendLinkPage"
	"imageresizerservice/users/loginWithEmailLink/sendLink/sendLinkSuccessPage"
)

func Router(mux *http.ServeMux, d *deps.Deps) {
	sendLinkPage.Router(mux)
	sendLinkAction.Router(mux, d)
	sendLinkSuccessPage.Router(mux)
}
