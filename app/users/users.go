package users

import (
	"imageresizerservice/app/ctx"
	"imageresizerservice/app/users/loginWithEmailLink"
	"net/http"
)

func Router(mux *http.ServeMux, appCtx *ctx.AppCtx) {
	loginWithEmailLink.Router(mux, appCtx)
}
