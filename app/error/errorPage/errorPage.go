package errorPage

import (
	"imageresizerservice/app/error/errorRoutes"
	"imageresizerservice/app/home/homeRoutes"
	"imageresizerservice/app/ui/page"
	"imageresizerservice/library/static"
	"net/http"
	"net/url"
)

func Router(mux *http.ServeMux) {
	mux.HandleFunc(errorRoutes.ServerError, Respond())
}

type Data struct {
	ErrorMessage string
	HomeURL      string
}

func Respond() http.HandlerFunc {
	htmlPath := static.GetSiblingPath("errorPage.html")
	return func(w http.ResponseWriter, r *http.Request) {
		data := Data{
			ErrorMessage: r.URL.Query().Get("error"),
			HomeURL:      homeRoutes.HomePage,
		}

		page.Respond(data, htmlPath)(w, r)
	}
}

func Redirect(w http.ResponseWriter, r *http.Request, errorMessage string) {
	u, _ := url.Parse(errorRoutes.ServerError)
	q := u.Query()
	q.Set("error", errorMessage)
	u.RawQuery = q.Encode()
	http.Redirect(w, r, u.String(), http.StatusSeeOther)
}
