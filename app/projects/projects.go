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

func Router(mux *http.ServeMux, ac *appContext.AppCtx) {
	createProject.Router(mux, ac)
	editProject.Router(mux, ac)
	deleteProject.Router(mux, ac)
	listProjects.Router(mux, ac)
	getProject.Router(mux, ac)
}
