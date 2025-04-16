package useLinkErrorPage

import (
	"imageresizerservice/app/ui/page"
	"imageresizerservice/app/users/login/loginRoutes"
	"imageresizerservice/library/static"
	"net/http"
	"net/url"
)

func Router(mux *http.ServeMux) {
	mux.HandleFunc(loginRoutes.UseLinkErrorPage, Respond())
}

type Data struct {
	Error string
}

func Respond() http.HandlerFunc {
	htmlPath := static.GetSiblingPath("useLinkErrorPage.html")
	return func(w http.ResponseWriter, r *http.Request) {
		data := Data{
			Error: r.URL.Query().Get("error"),
		}

		page.Respond(htmlPath, data)(w, r)
	}
}

func Redirect(w http.ResponseWriter, r *http.Request, error string) {
	u, _ := url.Parse(loginRoutes.UseLinkErrorPage)
	q := u.Query()
	q.Set("error", error)
	u.RawQuery = q.Encode()
	http.Redirect(w, r, u.String(), http.StatusSeeOther)
}
