package users

import (
	"imageService/app/ctx/appCtx"
	"imageService/app/users/login"
	"imageService/app/users/logout"
	"imageService/app/users/userAccount/accountPage"
	"net/http"
)

func Router(mux *http.ServeMux, ac *appCtx.AppCtx) {
	logout.Router(mux, ac)
	login.Router(mux, ac)
	accountPage.Router(mux, ac)
}

func RouterLoggedOut(mux *http.ServeMux, ac *appCtx.AppCtx) {
	login.RouterLoggedOut(mux, ac)
}
