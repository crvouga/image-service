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

func (d ConfirmationPage) Redirect(w http.ResponseWriter, r *http.Request) {
	u, _ := url.Parse(ConfirmationPageRoute)
	q := u.Query()
	q.Set("headline", d.Headline)
	q.Set("body", d.Body)
	q.Set("cancelURL", d.CancelURL)
	q.Set("cancelText", d.CancelText)
	q.Set("confirmURL", d.ConfirmURL)
	q.Set("confirmText", d.ConfirmText)
	u.RawQuery = q.Encode()
	http.Redirect(w, r, u.String(), http.StatusSeeOther)
}

func Router(mux *http.ServeMux) {
	mux.HandleFunc(ConfirmationPageRoute, Respond())
}

func Respond() http.HandlerFunc {
	htmlPath := static.GetSiblingPath("confirmationPage.html")
	return func(w http.ResponseWriter, r *http.Request) {
		data := ConfirmationPage{
			Headline:    r.URL.Query().Get("headline"),
			Body:        r.URL.Query().Get("body"),
			ConfirmURL:  r.URL.Query().Get("confirmURL"),
			ConfirmText: r.URL.Query().Get("confirmText"),
			CancelURL:   r.URL.Query().Get("cancelURL"),
			CancelText:  r.URL.Query().Get("cancelText"),
		}
		page.Respond(data, htmlPath)(w, r)
	}
}
