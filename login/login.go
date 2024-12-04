package login

import (
	"image-resizer-service/deps"
	"image-resizer-service/page"
	"net/http"
	"strings"
)

type LoginPageData struct {
	Action     string
	EmailError string
}

func RespondLoginPage(w http.ResponseWriter, r *http.Request) {
	errorEmail := r.URL.Query().Get("errorEmail")

	pageData := LoginPageData{
		Action:     "/login/send-link",
		EmailError: errorEmail,
	}
	page.Handler("./login/login_page.html", pageData)(w, r)
}

func RespondSentLinkPage(w http.ResponseWriter, r *http.Request) {
	page.Handler("./login/sent_link.html", nil)(w, r)
}

func RespondSendLink(d *deps.Deps) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if err := r.ParseForm(); err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		email := strings.TrimSpace(r.FormValue("email"))

		if email == "" {
			http.Redirect(w, r, "/login?errorEmail=Email+is+required", http.StatusSeeOther)
			return
		}

		err := d.SendEmail.SendEmail(
			email,
			"Login link",
			"Click here to login: http://localhost:8080/login/sent-link",
		)

		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/login/sent-link", http.StatusSeeOther)
	}
}
