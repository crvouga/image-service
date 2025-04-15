package sendLinkSuccessPage

import (
	"net/http"
	"net/url"

	"imageresizerservice/app/ui/page"
	"imageresizerservice/app/users/loginWithEmailLink/routes"
	"imageresizerservice/library/static"
)

type Data struct {
	SendAnotherLink string
	Email           string
}

func Router(mux *http.ServeMux) {
	mux.HandleFunc(routes.SendLinkSuccessPage, Respond())
}

func Respond() http.HandlerFunc {
	htmlPath := static.GetSiblingPath("sendLinkSuccessPage.html")
	return func(w http.ResponseWriter, r *http.Request) {
		data := Data{
			SendAnotherLink: routes.SendLinkPage,
			Email:           r.URL.Query().Get("Email"),
		}

		page.Respond(htmlPath, data)(w, r)
	}
}

func Redirect(w http.ResponseWriter, r *http.Request, email string) {
	u, _ := url.Parse(routes.SendLinkSuccessPage)
	q := u.Query()
	q.Set("Email", email)
	u.RawQuery = q.Encode()
	http.Redirect(w, r, u.String(), http.StatusSeeOther)
}
