package useLinkPage

import (
	"imageresizerservice/app/ctx/reqCtx"
	"imageresizerservice/app/ui/page"
	"imageresizerservice/app/users/login/link/linkID"
	"imageresizerservice/app/users/login/loginRoutes"
	"imageresizerservice/library/static"
	"net/http"
	"net/url"
)

func Router(mux *http.ServeMux) {
	mux.HandleFunc(loginRoutes.UseLinkPage, Respond())
}

type Data struct {
	Action string
	LinkID linkID.LinkID
}

func Respond() http.HandlerFunc {
	htmlPath := static.GetSiblingPath("useLinkPage.html")
	return func(w http.ResponseWriter, r *http.Request) {
		data := Data{
			Action: loginRoutes.UseLinkAction,
			LinkID: linkID.New(r.URL.Query().Get("linkID")),
		}

		page.Respond(htmlPath, data)(w, r)
	}
}

func ToUrl(rc *reqCtx.ReqCtx, linkID linkID.LinkID) string {
	path := ToPath(linkID)
	u, _ := url.Parse(rc.BaseURL + path)
	return u.String()
}

func ToPath(linkID linkID.LinkID) string {
	u, _ := url.Parse(loginRoutes.UseLinkPage)
	q := u.Query()
	q.Set("linkID", string(linkID))
	u.RawQuery = q.Encode()
	return u.String()
}
