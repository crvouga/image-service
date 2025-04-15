package useLinkResultPage

import (
	"imageresizerservice/page"
	"imageresizerservice/static"
	"imageresizerservice/users/loginWithEmailLink/routes"
	"net/http"
)

func Router(mux *http.ServeMux) {
	mux.HandleFunc(routes.UseLinkResultPage, Respond())
}

type Data struct {
	Action string
}

func Respond() http.HandlerFunc {
	htmlPath := static.GetSiblingPath("useLinkResultPage.html")
	return func(w http.ResponseWriter, r *http.Request) {
		data := Data{
			Action: routes.UseLinkAction,
		}

		page.Respond(htmlPath, data)(w, r)
	}
}
