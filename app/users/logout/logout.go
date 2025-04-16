package logout

import (
	"net/http"

	"imageresizerservice/app/ctx/appContext"
	"imageresizerservice/app/users/logout/logoutAction"
	"imageresizerservice/app/users/logout/logoutPage"
)

func Router(mux *http.ServeMux, appCtx *appContext.AppCtx) {
	logoutPage.Router(mux, appCtx)
	logoutAction.Router(mux, appCtx)
}
