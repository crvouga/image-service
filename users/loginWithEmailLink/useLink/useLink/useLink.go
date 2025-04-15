package useLink

import (
	"net/http"

	"imageresizerservice/deps"
	"imageresizerservice/users/loginWithEmailLink/useLink/useLinkAction"
	"imageresizerservice/users/loginWithEmailLink/useLink/useLinkPage"
	"imageresizerservice/users/loginWithEmailLink/useLink/useLinkResultPage"
)

func Router(mux *http.ServeMux, d *deps.Deps) {
	useLinkPage.Router(mux)
	useLinkAction.Router(mux, d)
	useLinkResultPage.Router(mux)
}
