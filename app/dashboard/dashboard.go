package dashboard

import (
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/dashboard/dashboardPage"
	"net/http"
)

func Router(mux *http.ServeMux, appCtx *appCtx.AppCtx) {
	dashboardPage.Router(mux, appCtx)
}
