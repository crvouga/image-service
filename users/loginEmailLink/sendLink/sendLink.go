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

		emailInput := strings.TrimSpace(r.FormValue("email"))

		errSent := sendLink(d, emailInput)

		if errSent != nil {
			loginPage.RedirectError(w, r, loginPage.RedirectErrorArgs{
				Email:      emailInput,
				EmailError: errSent.Error(),
			})
			return
		}

		sentLinkPage.Redirect(w, r, emailInput)
	}
}

func sendLink(d *deps.Deps, emailInput string) error {
	emailErr := email.Validate(emailInput)

	if emailErr != nil {
		return emailErr
	}

	err := d.SendEmail.SendEmail(
		emailInput,
		"Login link",
		"Click here to login: http://localhost:8080/login-with-email-link/sent-link",
	)

	if err != nil {
		return err
	}

	return nil
}
