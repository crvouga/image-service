package logout

import (
	"net/http"

	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/users/logout/logoutAction"
	"imageresizerservice/app/users/logout/logoutPage"
)

func Router(mux *http.ServeMux, ac *appCtx.AppCtx) {
	logoutPage.Router(mux, ac)
	logoutAction.Router(mux, ac)
}
