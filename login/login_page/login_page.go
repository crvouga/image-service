package login_page

import (
	"image-resizer-service/login/login_routes"
	"image-resizer-service/page"
	"net/http"
	"net/url"
	"strings"
)

type Data struct {
	Action     string
	EmailError string
}

func Respond() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		errorEmail := r.URL.Query().Get("errorEmail")

		data := Data{
			Action:     login_routes.SendLink,
			EmailError: errorEmail,
		}

		page.Respond("./login/login_page/login_page.html", data)(w, r)
	}
}

func RedirectError(w http.ResponseWriter, r *http.Request, errorEmail string) {
	cleaned := strings.ReplaceAll(strings.TrimSpace(errorEmail), " ", "+")
	u, _ := url.Parse(login_routes.LoginPage)
	q := u.Query()
	q.Set("errorEmail", cleaned)
	u.RawQuery = q.Encode()
	http.Redirect(w, r, u.String(), http.StatusSeeOther)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, login_routes.LoginPage, http.StatusSeeOther)
}
