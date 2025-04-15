package useLink

import (
	"net/http"

	"imageresizerservice/app/ctx"
	"imageresizerservice/app/users/loginWithEmailLink/useLink/useLinkAction"
	"imageresizerservice/app/users/loginWithEmailLink/useLink/useLinkPage"
	"imageresizerservice/app/users/loginWithEmailLink/useLink/useLinkSuccessPage"
)

func Router(mux *http.ServeMux, d *ctx.Ctx) {
	useLinkPage.Router(mux)
	useLinkAction.Router(mux, d)
	useLinkSuccessPage.Router(mux)
}
