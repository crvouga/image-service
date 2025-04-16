package users

import (
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/users/login"
	"imageresizerservice/app/users/logout"
	"imageresizerservice/app/users/userAccount/getUserAccount"
	"net/http"
)

func Router(mux *http.ServeMux, ac *appCtx.AppCtx) {
	logout.Router(mux, ac)
	login.Router(mux, ac)
	getUserAccount.Router(mux, ac)
}

func RouterLoggedOut(mux *http.ServeMux, ac *appCtx.AppCtx) {
	login.RouterLoggedOut(mux, ac)
}
