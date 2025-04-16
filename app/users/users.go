package users

import (
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/users/login"
	"imageresizerservice/app/users/logout"
	"net/http"
)

func Router(mux *http.ServeMux, appCtx *appCtx.AppCtx) {
	logout.Router(mux, appCtx)
	login.Router(mux, appCtx)
}

func RouterLoggedOut(mux *http.ServeMux, appCtx *appCtx.AppCtx) {
	login.RouterLoggedOut(mux, appCtx)

}
