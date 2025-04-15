package useLinkSuccessPage

import (
	"imageresizerservice/app/ui/page"
	"imageresizerservice/app/users/loginWithEmailLink/routes"
	"imageresizerservice/library/static"
	"net/http"
)

func Router(mux *http.ServeMux) {
	mux.HandleFunc(routes.UseLinkSuccessPage, Respond())
}

type Data struct {
}

func Respond() http.HandlerFunc {
	htmlPath := static.GetSiblingPath("useLinkSuccessPage.html")
	return func(w http.ResponseWriter, r *http.Request) {
		data := Data{}

		page.Respond(htmlPath, data)(w, r)
	}
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, routes.UseLinkSuccessPage, http.StatusSeeOther)
}
