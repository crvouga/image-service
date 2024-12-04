package send_link

import (
	"image-resizer-service/deps"
	"image-resizer-service/email"
	"image-resizer-service/login/login_page"
	"image-resizer-service/login/login_routes"
	"net/http"
	"strings"
)

func Router(mux *http.ServeMux, d *deps.Deps) {
	mux.HandleFunc(login_routes.SendLink, Respond(d))
}

func Respond(d *deps.Deps) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if err := r.ParseForm(); err != nil {
			login_page.RedirectError(w, r, login_page.RedirectErrorArgs{
				Email:      "",
				EmailError: "Unable to parse form",
			})
			return
		}

		emailValue := strings.TrimSpace(r.FormValue("email"))

		emailErr := email.Validate(emailValue)

		if emailErr != nil {
			login_page.RedirectError(w, r, login_page.RedirectErrorArgs{
				Email:      emailValue,
				EmailError: emailErr.Error(),
			})
			return
		}

		err := d.SendEmail.SendEmail(
			emailValue,
			"Login link",
			"Click here to login: http://localhost:8080/login/sent-link",
		)

		if err != nil {
			login_page.RedirectError(w, r, login_page.RedirectErrorArgs{
				Email:      emailValue,
				EmailError: "Unable to send email",
			})
			return
		}

		http.Redirect(w, r, login_routes.SentLinkPage, http.StatusSeeOther)
	}
}
