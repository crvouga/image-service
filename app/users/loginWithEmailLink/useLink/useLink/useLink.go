package useLink

import (
	"net/http"

	"imageresizerservice/app/deps"
	"imageresizerservice/app/users/loginWithEmailLink/useLink/useLinkAction"
	"imageresizerservice/app/users/loginWithEmailLink/useLink/useLinkPage"
	"imageresizerservice/app/users/loginWithEmailLink/useLink/useLinkResultPage"
)

func Router(mux *http.ServeMux, d *deps.Deps) {
	useLinkPage.Router(mux)
	useLinkAction.Router(mux, d)
	useLinkResultPage.Router(mux)
}
