package sendLink

import (
	"errors"
	"net/http"
	"sort"
	"strings"

	"imageService/app/ctx/appCtx"
	"imageService/app/ctx/reqCtx"
	"imageService/app/email/sendEmailFactory"
	"imageService/app/home/homeRoutes"
	"imageService/app/ui/errorPage"
	"imageService/app/ui/page"
	"imageService/app/ui/successPage"
	"imageService/app/users/login/link"
	"imageService/app/users/login/link/linkID"
	"imageService/app/users/login/loginRoutes"
	"imageService/app/users/login/useLink"
	"imageService/library/email/email"
	"imageService/library/email/emailAddress"
	"imageService/library/static"
)

type Data struct {
	Email string
}

func Router(mux *http.ServeMux, ac *appCtx.AppCtx) {
	mux.HandleFunc(loginRoutes.SendLinkPage, Respond(ac))
}

func Respond(ac *appCtx.AppCtx) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// Handle GET request - render the form
		if r.Method == http.MethodGet {
			data := Data{
				Email: r.URL.Query().Get("Email"),
			}

			page.Respond(data, static.GetSiblingPath("sendLink.html"))(w, r)
			return
		}

		// Handle POST request - process the form
		if r.Method == http.MethodPost {
			if err := r.ParseForm(); err != nil {
				errorPage.ErrorPage{
					Headline: "Error",
					Body:     "Unable to parse form",
					NextURL:  loginRoutes.ToSendLink(""),
					NextText: "Back",
				}.Render(w, r)
				return
			}

			emailInput := strings.TrimSpace(r.FormValue("email"))
			rc := reqCtx.FromHttpRequest(ac, r)

			errSent := SendLink(ac, &rc, emailInput)

			if errSent != nil {
				errorPage.ErrorPage{
					Headline: "Error",
					Body:     "Unable to send link: " + errSent.Error(),
					NextURL:  loginRoutes.ToSendLink(emailInput),
					NextText: "Back",
				}.Render(w, r)
				return
			}

			isSendEmailConfigured := sendEmailFactory.IsConfigured()
			loginLink := toLoginLink(ac, &rc)

			if !isSendEmailConfigured {
				errorPage.ErrorPage{
					Headline: "Authentication not configured",
					Body:     "The admin of this app has not configured authentication. You can use the email you sent a login link to without authentication.",
					NextURL:  loginLink,
					NextText: "Use login link",
				}.Render(w, r)
				return
			}

			successPage.New(
				"We have sent a login link to "+emailInput+". Please check your email to log in.",
				homeRoutes.HomePage,
				"Home",
			).Redirect(w, r)
			return
		}

		errorPage.New(errors.New("method not allowed")).Redirect(w, r)
	}
}

func SendLink(ac *appCtx.AppCtx, rc *reqCtx.ReqCtx, emailAddressInput string) error {
	emailAddr, err := emailAddress.New(emailAddressInput)
	if err != nil {
		return err
	}

	uow, err := ac.UowFactory.Begin()

	if err != nil {
		return err
	}

	defer uow.Rollback()

	linkNew := link.New(emailAddr, rc.SessionID)

	if err := ac.LinkDB.Upsert(uow, linkNew); err != nil {
		return err
	}

	email := toLoginEmail(rc, emailAddr, linkNew.ID)

	sendEmail := sendEmailFactory.FromReqCtx(rc)

	if err := sendEmail.SendEmail(uow, email); err != nil {
		return err
	}

	if err := uow.Commit(); err != nil {
		return err
	}

	return nil
}

func toLoginEmail(rc *reqCtx.ReqCtx, emailAddress emailAddress.EmailAddress, linkID linkID.LinkID) email.Email {
	return email.Email{
		To:      emailAddress,
		Subject: "Login link",
		Body:    useLink.ToUrl(rc, linkID),
	}
}
func toLoginLink(ac *appCtx.AppCtx, ctx *reqCtx.ReqCtx) string {
	link := toLatestLoginLink(ac, ctx)
	if link == nil {
		return ""
	}
	linkUrl := useLink.ToUrl(ctx, link.ID)
	return linkUrl
}

func toLatestLoginLink(ac *appCtx.AppCtx, ctx *reqCtx.ReqCtx) *link.Link {
	isSendEmailConfigured := sendEmailFactory.IsConfigured()

	if isSendEmailConfigured {
		return nil
	}

	links, err := ac.LinkDB.GetBySessionID(ctx.SessionID)

	if err != nil {
		return nil
	}

	if len(links) == 0 {
		return nil
	}

	latestFirst := make([]*link.Link, len(links))
	copy(latestFirst, links)
	sort.Slice(latestFirst, func(i, j int) bool {
		return latestFirst[i].CreatedAt.After(latestFirst[j].CreatedAt)
	})

	return latestFirst[0]
}

type RedirectErrorArgs struct {
	Email      string
	EmailError string
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, loginRoutes.SendLinkPage, http.StatusSeeOther)
}
