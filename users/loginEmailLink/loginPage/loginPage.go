package loginPage

import (
	"log"
	"net/http"
	"net/url"

	"imageresizerservice/page"
	"imageresizerservice/static"
	"imageresizerservice/users/loginEmailLink/routes"
)

type Data struct {
	Action     string
	EmailError string
	Email      string
}

func Router(mux *http.ServeMux) {
	mux.HandleFunc(routes.LoginPage, Respond())
}

func Respond() http.HandlerFunc {
	htmlPath := static.GetSiblingPath("loginPage.html")
	log.Println("htmlPath", htmlPath)
	jsPath := static.GetSiblingPath("loginPage.js")
	log.Println("jsPath", jsPath)
	return func(w http.ResponseWriter, r *http.Request) {
		data := Data{
			Action:     routes.SendLink,
			Email:      r.URL.Query().Get("Email"),
			EmailError: r.URL.Query().Get("ErrorEmail"),
		}

		page.Respond(htmlPath, data)(w, r)
	}
}

type RedirectErrorArgs struct {
	Email      string
	EmailError string
}

func RedirectError(w http.ResponseWriter, r *http.Request, args RedirectErrorArgs) {
	u, _ := url.Parse(routes.LoginPage)
	q := u.Query()
	q.Set("Email", args.Email)
	q.Set("ErrorEmail", args.EmailError)
	u.RawQuery = q.Encode()
	http.Redirect(w, r, u.String(), http.StatusSeeOther)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, routes.LoginPage, http.StatusSeeOther)
}
