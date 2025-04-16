package logout

import (
	"net/http"

	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/users/logout/logoutAction"
	"imageresizerservice/app/users/logout/logoutPage"
)

func Router(mux *http.ServeMux, appCtx *appCtx.AppCtx) {
	logoutPage.Router(mux, appCtx)
	logoutAction.Router(mux, appCtx)
}
