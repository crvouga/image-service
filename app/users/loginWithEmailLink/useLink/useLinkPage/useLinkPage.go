package useLinkPage

import (
	"imageresizerservice/app/ctx/reqCtx"
	"imageresizerservice/app/ui/page"
	"imageresizerservice/app/users/loginWithEmailLink/routes"
	"imageresizerservice/library/static"
	"net/http"
	"net/url"
)

func Router(mux *http.ServeMux) {
	mux.HandleFunc(routes.UseLinkPage, Respond())
}

type Data struct {
	Action string
	LinkId string
}

func Respond() http.HandlerFunc {
	htmlPath := static.GetSiblingPath("useLinkPage.html")
	return func(w http.ResponseWriter, r *http.Request) {
		data := Data{
			Action: routes.UseLinkAction,
			LinkId: r.URL.Query().Get("linkId"),
		}

		page.Respond(htmlPath, data)(w, r)
	}
}

func ToUrl(reqCtx *reqCtx.ReqCtx, linkId string) string {
	path := ToPath(linkId)
	u, _ := url.Parse(reqCtx.BaseURL + path)
	return u.String()
}

func ToPath(linkId string) string {
	u, _ := url.Parse(routes.UseLinkPage)
	q := u.Query()
	q.Set("linkId", linkId)
	u.RawQuery = q.Encode()
	return u.String()
}
