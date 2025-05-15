package projects

import (
	"imageService/app/ctx/appCtx"
	"imageService/app/projects/createProject"
	"imageService/app/projects/deleteProject"
	"imageService/app/projects/editProject"
	"imageService/app/projects/listProjects"
	"imageService/app/projects/projectPage"

	"net/http"
)

func Router(mux *http.ServeMux, ac *appCtx.AppCtx) {
	createProject.Router(mux, ac)
	editProject.Router(mux, ac)
	deleteProject.Router(mux, ac)
	listProjects.Router(mux, ac)
	projectPage.Router(mux, ac)
}
