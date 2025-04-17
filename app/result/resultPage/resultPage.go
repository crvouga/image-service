package resultPage

import (
	"imageresizerservice/app/ui/page"
	"imageresizerservice/library/static"
	"net/http"
	"net/url"
)

const (
	ResultRoute = "/result"
)

type ResultPageData struct {
	Headline string
	Body     string
	NextURL  string
	NextText string
}

func NewErrorPage(err error) *ResultPageData {
	return &ResultPageData{
		Headline: "Error",
		Body:     err.Error(),
		NextURL:  "/",
		NextText: "Go Home",
	}
}

func NewNotFoundPage(err error) *ResultPageData {
	return &ResultPageData{
		Headline: "Not Found",
		Body:     "The page you are looking for does not exist.",
		NextURL:  "/",
		NextText: "Go Home",
	}
}

func NewSuccessPage(message string, nextURL string, nextText string) *ResultPageData {
	return &ResultPageData{
		Headline: "Success",
		Body:     message,
		NextURL:  nextURL,
		NextText: nextText,
	}
}

func Router(mux *http.ServeMux) {
	mux.HandleFunc(ResultRoute, Respond())
}

func Respond() http.HandlerFunc {
	htmlPath := static.GetSiblingPath("resultPage.html")
	return func(w http.ResponseWriter, r *http.Request) {
		data := ResultPageData{
			Headline: r.URL.Query().Get("headline"),
			Body:     r.URL.Query().Get("body"),
			NextURL:  r.URL.Query().Get("nextURL"),
			NextText: r.URL.Query().Get("nextText"),
		}

		page.Respond(data, htmlPath)(w, r)
	}
}

func (d *ResultPageData) Redirect(w http.ResponseWriter, r *http.Request) {
	u, _ := url.Parse(ResultRoute)
	q := u.Query()
	q.Set("headline", d.Headline)
	q.Set("body", d.Body)
	q.Set("nextURL", d.NextURL)
	q.Set("nextText", d.NextText)
	u.RawQuery = q.Encode()
	http.Redirect(w, r, u.String(), http.StatusSeeOther)
}
