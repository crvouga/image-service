package login_page

import (
	"net/http"
	"net/url"

	"imageresizerservice.com/login/login_routes"
	"imageresizerservice.com/page"
)

type Data struct {
	Action     string
	EmailError string
	Email      string
}

func Router(mux *http.ServeMux) {
	mux.HandleFunc(login_routes.LoginPage, Respond())
}

func Respond() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := Data{
			Action:     login_routes.SendLink,
			Email:      r.URL.Query().Get("Email"),
			EmailError: r.URL.Query().Get("ErrorEmail"),
		}

		page.Respond("./login/login_page/login_page.html", data)(w, r)
	}
}

type RedirectErrorArgs struct {
	Email      string
	EmailError string
}

func RedirectError(w http.ResponseWriter, r *http.Request, args RedirectErrorArgs) {
	u, _ := url.Parse(login_routes.LoginPage)
	q := u.Query()
	q.Set("Email", args.Email)
	q.Set("ErrorEmail", args.EmailError)
	u.RawQuery = q.Encode()
	http.Redirect(w, r, u.String(), http.StatusSeeOther)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, login_routes.LoginPage, http.StatusSeeOther)
}
