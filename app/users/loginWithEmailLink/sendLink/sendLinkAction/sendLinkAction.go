package sendLinkAction

import (
	"net/http"
	"strings"

	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/ctx/reqCtx"
	"imageresizerservice/app/users/loginWithEmailLink/link"
	"imageresizerservice/app/users/loginWithEmailLink/routes"
	"imageresizerservice/app/users/loginWithEmailLink/sendLink/sendLinkPage"
	"imageresizerservice/app/users/loginWithEmailLink/sendLink/sendLinkSuccessPage"
	"imageresizerservice/app/users/loginWithEmailLink/useLink/useLinkPage"
	"imageresizerservice/library/email/email"
)

func Router(mux *http.ServeMux, appCtx *appCtx.AppCtx) {
	mux.HandleFunc(routes.SendLinkAction, Respond(appCtx))
}

func Respond(appCtx *appCtx.AppCtx) http.HandlerFunc {
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

		reqCtx := reqCtx.FromHttpRequest(r)

		errSent := SendLink(appCtx, &reqCtx, emailInput)

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

func SendLink(appCtx *appCtx.AppCtx, reqCtx *reqCtx.ReqCtx, emailAddressInput string) error {
	if err := email.ValidateEmailAddress(emailAddressInput); err != nil {
		return err
	}

	uow, err := appCtx.UowFactory.Begin()

	if err != nil {
		return err
	}

	defer uow.Rollback()

	linkNew := link.New(emailAddressInput)

	if err := appCtx.LinkDb.Upsert(uow, linkNew); err != nil {
		return err
	}

	email := toLoginEmail(reqCtx, emailAddressInput, linkNew.Id)

	if err := appCtx.SendEmail.SendEmail(uow, email); err != nil {
		return err
	}

	if err := uow.Commit(); err != nil {
		return err
	}

	return nil
}
func toLoginEmail(reqCtx *reqCtx.ReqCtx, emailAddress string, linkId string) email.Email {
	return email.Email{
		To:      emailAddress,
		From:    "noreply@imageresizer.com",
		Subject: "Login link",
		Body:    useLinkPage.ToUrl(reqCtx, linkId),
	}
}
