package errorPage

import (
	"imageresizerservice/app/ui/page"
	"imageresizerservice/library/static"
	"net/http"
	"net/url"
)

const (
	ErrorRoute = "/error"
)

type ErrorPage struct {
	Headline string
	Body     string
	NextURL  string
	NextText string
}

func (d ErrorPage) Redirect(w http.ResponseWriter, r *http.Request) {
	u, _ := url.Parse(ErrorRoute)
	q := u.Query()
	q.Set("headline", d.Headline)
	q.Set("body", d.Body)
	q.Set("nextURL", d.NextURL)
	q.Set("nextText", d.NextText)
	u.RawQuery = q.Encode()
	http.Redirect(w, r, u.String(), http.StatusFound)
}

func New(err error) *ErrorPage {
	return &ErrorPage{
		Headline: "Error",
		Body:     err.Error(),
		NextURL:  "/",
		NextText: "Home",
	}
}

func Router(mux *http.ServeMux) {
	mux.HandleFunc(ErrorRoute, Respond())
}

func Respond() http.HandlerFunc {
	htmlPath := static.GetSiblingPath("errorPage.html")
	return func(w http.ResponseWriter, r *http.Request) {
		data := ErrorPage{
			Headline: r.URL.Query().Get("headline"),
			Body:     r.URL.Query().Get("body"),
			NextURL:  r.URL.Query().Get("nextURL"),
			NextText: r.URL.Query().Get("nextText"),
		}

		page.Respond(data, htmlPath)(w, r)
	}
}
