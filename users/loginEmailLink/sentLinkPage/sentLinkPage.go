package sentLinkPage

import (
	"net/http"
	"net/url"

	"imageresizerservice/page"
	"imageresizerservice/static"
	"imageresizerservice/users/loginEmailLink/routes"
)

type Data struct {
	SendAnotherLink string
	Email           string
}

func Router(mux *http.ServeMux) {
	mux.HandleFunc(routes.SentLinkPage, Respond())
}

func Respond() http.HandlerFunc {
	htmlPath := static.GetSiblingPath("sentLinkPage.html")
	return func(w http.ResponseWriter, r *http.Request) {
		data := Data{
			SendAnotherLink: routes.LoginPage,
			Email:           r.URL.Query().Get("Email"),
		}

		page.Respond(htmlPath, data)(w, r)
	}
}

func Redirect(w http.ResponseWriter, r *http.Request, email string) {
	u, _ := url.Parse(routes.SentLinkPage)
	q := u.Query()
	q.Set("Email", email)
	u.RawQuery = q.Encode()
	http.Redirect(w, r, u.String(), http.StatusSeeOther)
}
