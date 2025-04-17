package notFoundPage

import (
	"imageresizerservice/app/ui/page"
	"imageresizerservice/library/static"
	"net/http"
	"net/url"
)

const (
	NotFoundRoute = "/not-found"
)

type NotFoundPageData struct {
	Headline string
	Body     string
	NextURL  string
	NextText string
}

func (d NotFoundPageData) Redirect(w http.ResponseWriter, r *http.Request) {
	u, _ := url.Parse(NotFoundRoute)
	q := u.Query()
	q.Set("headline", d.Headline)
	q.Set("body", d.Body)
	q.Set("nextURL", d.NextURL)
	q.Set("nextText", d.NextText)
	u.RawQuery = q.Encode()
	http.Redirect(w, r, u.String(), http.StatusSeeOther)
}

func New(nextURL string, nextText string) *NotFoundPageData {
	return &NotFoundPageData{
		Headline: "Not Found",
		Body:     "The page you are looking for does not exist.",
		NextURL:  nextURL,
		NextText: nextText,
	}
}

func Router(mux *http.ServeMux) {
	mux.HandleFunc(NotFoundRoute, Respond())
}

func Respond() http.HandlerFunc {
	htmlPath := static.GetSiblingPath("notFoundPage.html")
	return func(w http.ResponseWriter, r *http.Request) {
		data := NotFoundPageData{
			Headline: r.URL.Query().Get("headline"),
			Body:     r.URL.Query().Get("body"),
			NextURL:  r.URL.Query().Get("nextURL"),
			NextText: r.URL.Query().Get("nextText"),
		}

		page.Respond(data, htmlPath)(w, r)
	}
}
