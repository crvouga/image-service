package sent_link_page

import (
	"net/http"
	"net/url"

	"imageresizerservice.com/login/login_routes"
	"imageresizerservice.com/page"
)

type Data struct {
	SendAnotherLink string
	Email           string
}

func Router(mux *http.ServeMux) {
	mux.HandleFunc(login_routes.SentLinkPage, Respond())
}

func Respond() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := Data{
			SendAnotherLink: login_routes.LoginPage,
			Email:           r.URL.Query().Get("Email"),
		}
		page.Respond("./login/sent_link_page/sent_link_page.html", data)(w, r)
	}
}

func Redirect(w http.ResponseWriter, r *http.Request, email string) {
	u, _ := url.Parse(login_routes.SentLinkPage)
	q := u.Query()
	q.Set("Email", email)
	u.RawQuery = q.Encode()
	http.Redirect(w, r, u.String(), http.StatusSeeOther)
}
