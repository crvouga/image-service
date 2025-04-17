package useLinkSuccessPage

import (
	"imageresizerservice/app/home/homeRoutes"
	"imageresizerservice/app/ui/page"
	"imageresizerservice/app/users/login/loginRoutes"
	"imageresizerservice/library/static"
	"net/http"
)

func Router(mux *http.ServeMux) {
	mux.HandleFunc(loginRoutes.UseLinkSuccessPage, Respond())
}

type Data struct {
	Home string
}

func Respond() http.HandlerFunc {
	htmlPath := static.GetSiblingPath("useLinkSuccessPage.html")
	return func(w http.ResponseWriter, r *http.Request) {
		data := Data{
			Home: homeRoutes.HomePage,
		}

		page.Respond(data, htmlPath)(w, r)
	}
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, loginRoutes.UseLinkSuccessPage, http.StatusSeeOther)
}
