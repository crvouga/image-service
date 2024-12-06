package sendLink

import (
	"net/http"
	"strings"

	"imageresizerservice/deps"
	"imageresizerservice/email"
	"imageresizerservice/users/loginEmailLink/loginPage"
	"imageresizerservice/users/loginEmailLink/routes"
	"imageresizerservice/users/loginEmailLink/sentLinkPage"
)

func Router(mux *http.ServeMux, d *deps.Deps) {
	mux.HandleFunc(routes.SendLink, Respond(d))
}

func Respond(d *deps.Deps) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if err := r.ParseForm(); err != nil {
			loginPage.RedirectError(w, r, loginPage.RedirectErrorArgs{
				Email:      "",
				EmailError: "Unable to parse form",
			})
			return
		}

		email_input := strings.TrimSpace(r.FormValue("email"))

		err_send := send_link(d, email_input)

		if err_send != nil {
			loginPage.RedirectError(w, r, loginPage.RedirectErrorArgs{
				Email:      email_input,
				EmailError: err_send.Error(),
			})
			return
		}

		sentLinkPage.Redirect(w, r, email_input)
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
		"Click here to login: http://localhost:8080/login-with-email-link/sent-link",
	)

	if err != nil {
		return err
	}

	return nil
}
