package projects

import (
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/projects/projectCreate"
	"imageresizerservice/app/projects/projectDelete"
	"imageresizerservice/app/projects/projectEdit"
	"imageresizerservice/app/projects/projectPage"
	"net/http"
)

func Router(mux *http.ServeMux, appCtx *appCtx.AppCtx) {
	projectCreate.Router(mux, appCtx)
	projectEdit.Router(mux, appCtx)
	projectDelete.Router(mux, appCtx)
	projectPage.Router(mux, appCtx)
}
