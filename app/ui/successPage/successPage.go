package successPage

import (
	"imageresizerservice/app/ui/page"
	"imageresizerservice/library/static"
	"net/http"
	"net/url"
)

const (
	SuccessRoute = "/success"
)

type SuccessPage struct {
	Headline string
	Body     string
	NextURL  string
	NextText string
}

func (d SuccessPage) Redirect(w http.ResponseWriter, r *http.Request) {
	u, _ := url.Parse(SuccessRoute)
	q := u.Query()
	q.Set("headline", d.Headline)
	q.Set("body", d.Body)
	q.Set("nextURL", d.NextURL)
	q.Set("nextText", d.NextText)
	u.RawQuery = q.Encode()
	http.Redirect(w, r, u.String(), http.StatusSeeOther)
}

func New(message string, nextURL string, nextText string) *SuccessPage {
	return &SuccessPage{
		Headline: "Success",
		Body:     message,
		NextURL:  nextURL,
		NextText: nextText,
	}
}

func Router(mux *http.ServeMux) {
	mux.HandleFunc(SuccessRoute, Respond())
}

func Respond() http.HandlerFunc {
	htmlPath := static.GetSiblingPath("successPage.html")
	return func(w http.ResponseWriter, r *http.Request) {
		data := SuccessPage{
			Headline: r.URL.Query().Get("headline"),
			Body:     r.URL.Query().Get("body"),
			NextURL:  r.URL.Query().Get("nextURL"),
			NextText: r.URL.Query().Get("nextText"),
		}

		page.Respond(data, htmlPath)(w, r)
	}
}
