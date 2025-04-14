package sendLinkAction

import (
	"net/http"
	"strings"

	"imageresizerservice/deps"
	"imageresizerservice/email/email"
	"imageresizerservice/users/loginWithEmailLink/link"
	"imageresizerservice/users/loginWithEmailLink/routes"
	"imageresizerservice/users/loginWithEmailLink/sendLink/sendLinkPage"
	"imageresizerservice/users/loginWithEmailLink/sendLink/sendLinkSuccessPage"
)

func Router(mux *http.ServeMux, d *deps.Deps) {
	mux.HandleFunc(routes.SendLinkAction, Respond(d))
}

func Respond(d *deps.Deps) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if err := r.ParseForm(); err != nil {
			sendLinkPage.RedirectError(w, r, sendLinkPage.RedirectErrorArgs{
				Email:      "",
				EmailError: "Unable to parse form",
			})
			return
		}

		emailInput := strings.TrimSpace(r.FormValue("email"))

		errSent := SendLink(d, emailInput)

		if errSent != nil {
			sendLinkPage.RedirectError(w, r, sendLinkPage.RedirectErrorArgs{
				Email:      emailInput,
				EmailError: errSent.Error(),
			})
			return
		}

		sendLinkSuccessPage.Redirect(w, r, emailInput)
	}
}

func SendLink(d *deps.Deps, emailAddressInput string) error {
	if err := email.ValidateEmailAddress(emailAddressInput); err != nil {
		return err
	}

	uow, err := d.UowFactory.Begin()

	if err != nil {
		return err
	}

	defer uow.Rollback()

	linkNew := link.New(emailAddressInput)

	if err := d.LinkDb.Upsert(uow, linkNew); err != nil {
		return err
	}

	email := email.Email{
		To:      emailAddressInput,
		From:    "noreply@imageresizer.com",
		Subject: "Login link",
		Body:    "Click here to login: http://localhost:8080/login-with-email-link/use-link-page?linkId=" + linkNew.Id,
	}

	if err := d.EmailOutbox.Add(uow, email); err != nil {
		return err
	}

	if err := uow.Commit(); err != nil {
		return err
	}

	return nil
}
