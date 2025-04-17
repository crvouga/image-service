package admin

import (
	"imageresizerservice/app/admin/claimAdmin"
	"imageresizerservice/app/ctx/appCtx"
	"net/http"
)

func Router(mux *http.ServeMux, ac *appCtx.AppCtx) {
	claimAdmin.Router(mux, ac)
}
