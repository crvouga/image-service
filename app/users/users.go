package users

import (
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/users/loginWithEmailLink"
	"net/http"
)

func Router(mux *http.ServeMux, appCtx *appCtx.AppCtx) {
	loginWithEmailLink.Router(mux, appCtx)
}
