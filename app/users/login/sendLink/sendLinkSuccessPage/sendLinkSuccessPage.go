package sendLinkSuccessPage

import (
	"imageresizerservice/app/ctx/appContext"
	"imageresizerservice/app/ctx/reqCtx"
	"imageresizerservice/app/email/sendEmailFactory"
	"imageresizerservice/app/ui/page"
	"imageresizerservice/app/users/login/link"
	"imageresizerservice/app/users/login/loginRoutes"
	"imageresizerservice/app/users/login/useLink/useLinkPage"
	"imageresizerservice/library/static"
	"net/http"
	"net/url"
	"sort"
	"strconv"
)

type Data struct {
	SendAnotherLink       string
	Email                 string
	IsSendEmailConfigured string
	LoginLink             string
}

func Router(mux *http.ServeMux, ac *appContext.AppCtx) {
	mux.HandleFunc(loginRoutes.SendLinkSuccessPage, Respond(ac))
}

func Respond(ac *appContext.AppCtx) http.HandlerFunc {
	htmlPath := static.GetSiblingPath("sendLinkSuccessPage.html")
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := reqCtx.FromHttpRequest(ac, r)

		isSendEmailConfigured := sendEmailFactory.IsConfigured()
		loginLink := toLoginLink(ac, &ctx)

		data := Data{
			SendAnotherLink:       loginRoutes.SendLinkPage,
			Email:                 r.URL.Query().Get("Email"),
			IsSendEmailConfigured: strconv.FormatBool(isSendEmailConfigured),
			LoginLink:             loginLink,
		}

		page.Respond(htmlPath, data)(w, r)
	}
}

func toLoginLink(ac *appContext.AppCtx, ctx *reqCtx.ReqCtx) string {
	link := toLatestLoginLink(ac, ctx)
	if link == nil {
		return ""
	}
	linkUrl := useLinkPage.ToUrl(ctx, link.ID)
	return linkUrl
}

func toLatestLoginLink(ac *appContext.AppCtx, ctx *reqCtx.ReqCtx) *link.Link {
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

func Redirect(w http.ResponseWriter, r *http.Request, email string) {
	u, _ := url.Parse(loginRoutes.SendLinkSuccessPage)
	q := u.Query()
	q.Set("Email", email)
	u.RawQuery = q.Encode()
	http.Redirect(w, r, u.String(), http.StatusSeeOther)
}

func parseConfiguredFlag(r *http.Request) string {
	configStr := r.URL.Query().Get("IsSendEmailConfigured")
	isConfigured, err := strconv.ParseBool(configStr)
	if configStr == "" || err != nil {
		isConfigured = true
	}
	return strconv.FormatBool(isConfigured)
}
