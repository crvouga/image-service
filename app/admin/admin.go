package admin

import (
	"imageService/app/admin/adminPage"
	"imageService/app/admin/claimAdmin"
	"imageService/app/ctx/appCtx"
	"net/http"
)

func Router(mux *http.ServeMux, ac *appCtx.AppCtx) {
	claimAdmin.Router(mux, ac)
	adminPage.Router(mux, ac)
}
