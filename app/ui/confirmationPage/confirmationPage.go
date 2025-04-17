package confirmationPage

import (
	"imageresizerservice/app/ui/breadcrumbs"
	"imageresizerservice/app/ui/page"
	"imageresizerservice/library/static"
	"net/http"
	"net/url"
	"strings"
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
	HiddenForm  map[string]string
	Breadcrumbs []breadcrumbs.Breadcrumb
}

func (d ConfirmationPage) ToQueryParams() url.Values {
	q := url.Values{}
	q.Set("headline", d.Headline)
	q.Set("body", d.Body)
	q.Set("cancelURL", d.CancelURL)
	q.Set("cancelText", d.CancelText)
	q.Set("confirmURL", d.ConfirmURL)
	q.Set("confirmText", d.ConfirmText)

	for key, value := range d.HiddenForm {
		q.Set("hidden_"+key, value)
	}

	breadcrumbParams := breadcrumbs.ToQueryParams(d.Breadcrumbs)
	for key, values := range breadcrumbParams {
		for _, value := range values {
			q.Add(key, value)
		}
	}
	return q
}

func FromQueryParams(query url.Values) ConfirmationPage {
	hiddenForm := make(map[string]string)
	for key, values := range query {
		if strings.HasPrefix(key, "hidden_") {
			hiddenForm[strings.TrimPrefix(key, "hidden_")] = values[0]
		}
	}

	breadcrumbs := breadcrumbs.FromQueryParams(query)

	return ConfirmationPage{
		Headline:    query.Get("headline"),
		Body:        query.Get("body"),
		ConfirmURL:  query.Get("confirmURL"),
		ConfirmText: query.Get("confirmText"),
		CancelURL:   query.Get("cancelURL"),
		CancelText:  query.Get("cancelText"),
		HiddenForm:  hiddenForm,
		Breadcrumbs: breadcrumbs,
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
