package sendLinkAction

import (
	"net/http"
	"strings"

	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/ctx/reqCtx"
	"imageresizerservice/app/email/sendEmailFactory"
	"imageresizerservice/app/users/login/link"
	"imageresizerservice/app/users/login/link/linkID"
	"imageresizerservice/app/users/login/loginRoutes"
	"imageresizerservice/app/users/login/sendLink/sendLinkPage"
	"imageresizerservice/app/users/login/sendLink/sendLinkSuccessPage"
	"imageresizerservice/app/users/login/useLink/useLinkPage"
	"imageresizerservice/library/email/email"
	"imageresizerservice/library/email/emailAddress"
)

func Router(mux *http.ServeMux, appCtx *appCtx.AppCtx) {
	mux.HandleFunc(loginRoutes.SendLinkAction, Respond(appCtx))
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

		reqCtxInst := reqCtx.FromHttpRequest(appCtx, r)

		errSent := SendLink(appCtx, &reqCtxInst, emailInput)

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
	emailAddress, err := emailAddress.New(emailAddressInput)
	if err != nil {
		return err
	}

	uow, err := appCtx.UowFactory.Begin()

	if err != nil {
		return err
	}

	defer uow.Rollback()

	linkNew := link.New(emailAddress, reqCtx.SessionID)

	if err := appCtx.LinkDB.Upsert(uow, linkNew); err != nil {
		return err
	}

	email := toLoginEmail(reqCtx, emailAddress, linkNew.ID)

	sendEmail := sendEmailFactory.FromReqCtx(reqCtx)

	if err := sendEmail.SendEmail(uow, email); err != nil {
		return err
	}

	if err := uow.Commit(); err != nil {
		return err
	}

	return nil
}
func toLoginEmail(reqCtx *reqCtx.ReqCtx, emailAddress emailAddress.EmailAddress, linkID linkID.LinkID) email.Email {
	return email.Email{
		To:      emailAddress,
		Subject: "Login link",
		Body:    useLinkPage.ToUrl(reqCtx, linkID),
	}
}
