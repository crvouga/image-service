package sent_link_page

import (
	"image-resizer-service/login/login_routes"
	"image-resizer-service/page"
	"net/http"
)

type Data struct {
	SendAnotherLink string
	Email           string
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
