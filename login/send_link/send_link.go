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

		email_input := strings.TrimSpace(r.FormValue("email"))

		err_send := send_link(d, email_input)

		if err_send != nil {
			login_page.RedirectError(w, r, login_page.RedirectErrorArgs{
				Email:      email_input,
				EmailError: err_send.Error(),
			})
			return
		}

		http.Redirect(w, r, login_routes.SentLinkPage, http.StatusSeeOther)
	}
}

func send_link(d *deps.Deps, email_input string) error {
	email_err := email.Validate(email_input)

	if email_err != nil {
		return email_err
	}

	err := d.SendEmail.SendEmail(
		email_input,
		"Login link",
		"Click here to login: http://localhost:8080/login/sent-link",
	)

	if err != nil {
		return err
	}

	return nil
}
