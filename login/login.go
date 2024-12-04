package login

import (
	"image-resizer-service/deps"
	"image-resizer-service/page"
	"net/http"
	"strings"
	"time"
)

type LoginPageData struct {
	Action string
}

func RespondLoginPage(w http.ResponseWriter, r *http.Request) {
	page.Handler("./login/login_page.html", nil)(w, r)
}

func RespondSentLinkPage(w http.ResponseWriter, r *http.Request) {
	page.Handler("./login/sent_link.html", nil)(w, r)
}

func RespondSendLink(d *deps.Deps) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second)
		time.Sleep(time.Second)
		email := r.FormValue("email")
		email = strings.TrimSpace(email)

		if email == "" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
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
