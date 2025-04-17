package confirmationPage

import (
	"imageresizerservice/app/ui/page"
	"imageresizerservice/library/static"
	"net/http"
	"net/url"
)

const (
	ConfirmationPageRoute = "/confirmation"
)

type ConfirmationPage struct {
	Headline    string
	Body        string
	CancelURL   string
	CancelText  string
	ConfirmURL  string
	ConfirmText string
}

func (d ConfirmationPage) ToQueryParams() url.Values {
	q := url.Values{}
	q.Set("headline", d.Headline)
	q.Set("body", d.Body)
	q.Set("cancelURL", d.CancelURL)
	q.Set("cancelText", d.CancelText)
	q.Set("confirmURL", d.ConfirmURL)
	q.Set("confirmText", d.ConfirmText)
	return q
}

func FromQueryParams(query url.Values) ConfirmationPage {
	return ConfirmationPage{
		Headline:    query.Get("headline"),
		Body:        query.Get("body"),
		ConfirmURL:  query.Get("confirmURL"),
		ConfirmText: query.Get("confirmText"),
		CancelURL:   query.Get("cancelURL"),
		CancelText:  query.Get("cancelText"),
	}
}

func (d ConfirmationPage) Redirect(w http.ResponseWriter, r *http.Request) {
	u, _ := url.Parse(ConfirmationPageRoute)
	u.RawQuery = d.ToQueryParams().Encode()
	http.Redirect(w, r, u.String(), http.StatusSeeOther)
}

func Router(mux *http.ServeMux) {
	mux.HandleFunc(ConfirmationPageRoute, func(w http.ResponseWriter, r *http.Request) {
		FromQueryParams(r.URL.Query()).Render(w, r)
	})
}

func (d ConfirmationPage) Render(w http.ResponseWriter, r *http.Request) {
	htmlPath := static.GetSiblingPath("confirmationPage.html")
	page.Respond(d, htmlPath)(w, r)
}
