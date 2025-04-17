package sendLink

import (
	"errors"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/ctx/reqCtx"
	"imageresizerservice/app/email/sendEmailFactory"
	"imageresizerservice/app/home/homeRoutes"
	"imageresizerservice/app/ui/errorPage"
	"imageresizerservice/app/ui/page"
	"imageresizerservice/app/ui/successPage"
	"imageresizerservice/app/users/login/link"
	"imageresizerservice/app/users/login/link/linkID"
	"imageresizerservice/app/users/login/loginRoutes"
	"imageresizerservice/app/users/login/useLink"
	"imageresizerservice/library/email/email"
	"imageresizerservice/library/email/emailAddress"
	"imageresizerservice/library/static"
)

type Data struct {
	Action     string
	EmailError string
	Email      string
	JsPath     string
}

func Router(mux *http.ServeMux, ac *appCtx.AppCtx) {
	mux.HandleFunc(loginRoutes.SendLinkPage, Respond(ac))
}

func Respond(ac *appCtx.AppCtx) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// Handle GET request - render the form
		if r.Method == http.MethodGet {
			data := Data{
				Action:     loginRoutes.SendLinkPage,
				Email:      r.URL.Query().Get("Email"),
				EmailError: r.URL.Query().Get("ErrorEmail"),
			}

			page.Respond(data, static.GetSiblingPath("sendLink.html"))(w, r)
			return
		}

		// Handle POST request - process the form
		if r.Method == http.MethodPost {
			if err := r.ParseForm(); err != nil {
				RedirectError(w, r, RedirectErrorArgs{
					Email:      "",
					EmailError: "Unable to parse form",
				})
				return
			}

			emailInput := strings.TrimSpace(r.FormValue("email"))
			rc := reqCtx.FromHttpRequest(ac, r)

			errSent := SendLink(ac, &rc, emailInput)

			if errSent != nil {
				RedirectError(w, r, RedirectErrorArgs{
					Email:      emailInput,
					EmailError: errSent.Error(),
				})
				return
			}

			isSendEmailConfigured := sendEmailFactory.IsConfigured()
			loginLink := toLoginLink(ac, &rc)

			if !isSendEmailConfigured {
				errorPage.ErrorPage{
					Headline: "Authentication is not enabled",
					Body:     "The admin of this app has not enabled authentication. You can use the email you sent a login link to without authentication.",
					NextURL:  loginLink,
					NextText: "Use login link",
				}.Redirect(w, r)
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

func RedirectError(w http.ResponseWriter, r *http.Request, args RedirectErrorArgs) {
	u, _ := url.Parse(loginRoutes.SendLinkPage)
	q := u.Query()
	q.Set("Email", args.Email)
	q.Set("ErrorEmail", args.EmailError)
	u.RawQuery = q.Encode()
	http.Redirect(w, r, u.String(), http.StatusSeeOther)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, loginRoutes.SendLinkPage, http.StatusSeeOther)
}
