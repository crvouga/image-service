package projects

import (
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/projects/createProject"
	"imageresizerservice/app/projects/deleteProject"
	"imageresizerservice/app/projects/editProject"
	"imageresizerservice/app/projects/getProject"
	"imageresizerservice/app/projects/listProjects"

	"net/http"
)

func Router(mux *http.ServeMux, appCtx *appCtx.AppCtx) {
	createProject.Router(mux, appCtx)
	editProject.Router(mux, appCtx)
	deleteProject.Router(mux, appCtx)
	listProjects.Router(mux, appCtx)
	getProject.Router(mux, appCtx)
}
