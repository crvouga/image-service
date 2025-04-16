package sendLinkSuccessPage

import (
	"imageresizerservice/app/ui/page"
	"imageresizerservice/app/users/login/loginRoutes"
	"imageresizerservice/library/static"
	"net/http"
	"net/url"
	"strconv"
)

type Data struct {
	SendAnotherLink       string
	Email                 string
	IsSendEmailConfigured string
	LoginLink             string
}

func Router(mux *http.ServeMux) {
	mux.HandleFunc(loginRoutes.SendLinkSuccessPage, Respond())
}

func Respond() http.HandlerFunc {
	htmlPath := static.GetSiblingPath("sendLinkSuccessPage.html")
	return func(w http.ResponseWriter, r *http.Request) {
		data := Data{
			SendAnotherLink:       loginRoutes.SendLinkPage,
			Email:                 r.URL.Query().Get("Email"),
			IsSendEmailConfigured: parseConfiguredFlag(r),
			LoginLink:             r.URL.Query().Get("LoginLink"),
		}

		page.Respond(htmlPath, data)(w, r)
	}
}

func Redirect(w http.ResponseWriter, r *http.Request, email string, isSendEmailConfigured bool, loginLink string) {
	u, _ := url.Parse(loginRoutes.SendLinkSuccessPage)
	q := u.Query()
	q.Set("Email", email)
	q.Set("IsSendEmailConfigured", strconv.FormatBool(isSendEmailConfigured))
	q.Set("LoginLink", loginLink)
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
