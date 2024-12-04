package send_link

import (
	"image-resizer-service/deps"
	"image-resizer-service/login/login_page"
	"image-resizer-service/login/login_routes"

	"net/http"
	"strings"
)

func Respond(d *deps.Deps) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if err := r.ParseForm(); err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		email := strings.TrimSpace(r.FormValue("email"))

		if email == "" {
			login_page.RedirectError(w, r, "Email is required")
			return
		}

		err := d.SendEmail.SendEmail(
			email,
			"Login link",
			"Click here to login: http://localhost:8080/login/sent-link",
		)

		if err != nil {
			login_page.RedirectError(w, r, "Unable to send email")
			return
		}

		http.Redirect(w, r, login_routes.SentLinkPage, http.StatusSeeOther)
	}
}
