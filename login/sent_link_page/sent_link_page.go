package sent_link_page

import (
	"net/http"

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
			Email:           r.URL.Query().Get("email"),
		}
		page.Respond("./login/sent_link_page/sent_link_page.html", data)(w, r)
	}
}
