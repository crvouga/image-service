package logout

import (
	"net/http"

	"imageresizerservice/app/ctx/appCtx"

	"imageresizerservice/app/users/logout/logoutPage"
)

func Router(mux *http.ServeMux, ac *appCtx.AppCtx) {
	logoutPage.Router(mux, ac)
}
