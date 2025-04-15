package users

import (
	"imageresizerservice/app/ctx"
	"imageresizerservice/app/users/loginWithEmailLink"
	"net/http"
)

func Router(mux *http.ServeMux, d *ctx.Ctx) {
	loginWithEmailLink.Router(mux, d)
}
