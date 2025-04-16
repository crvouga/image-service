package projects

import (
	"imageresizerservice/app/ctx/appContext"
	"imageresizerservice/app/projects/createProject"
	"imageresizerservice/app/projects/deleteProject"
	"imageresizerservice/app/projects/editProject"
	"imageresizerservice/app/projects/getProject"
	"imageresizerservice/app/projects/listProjects"

	"net/http"
)

func Router(mux *http.ServeMux, appCtx *appContext.AppCtx) {
	createProject.Router(mux, appCtx)
	editProject.Router(mux, appCtx)
	deleteProject.Router(mux, appCtx)
	listProjects.Router(mux, appCtx)
	getProject.Router(mux, appCtx)
}
