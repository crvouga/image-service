package loginWithEmailLink

import (
	"net/http"

	"imageresizerservice/app/ctx"
	"imageresizerservice/app/users/loginWithEmailLink/sendLink/sendLink"
	"imageresizerservice/app/users/loginWithEmailLink/useLink/useLink"
)

func Router(mux *http.ServeMux, d *ctx.Ctx) {
	sendLink.Router(mux, d)
	useLink.Router(mux, d)
}
