package users

import (
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/users/login"
	"net/http"
)

func Router(mux *http.ServeMux, appCtx *appCtx.AppCtx) {
	login.Router(mux, appCtx)
}
