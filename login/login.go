package login

import (
	"image-resizer-service/deps"
	"image-resizer-service/login/login_page"
	"image-resizer-service/login/login_routes"
	"image-resizer-service/login/send_link"
	"image-resizer-service/login/sent_link_page"
	"net/http"
)

func Router(d deps.Deps) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(login_routes.LoginPage, login_page.Respond())
	mux.HandleFunc(login_routes.SendLink, send_link.Respond(&d))
	mux.HandleFunc(login_routes.SentLinkPage, sent_link_page.Respond())
	mux.HandleFunc(RemoveTrailingSlash(login_routes.Prefix), login_page.Redirect)
	mux.HandleFunc(login_routes.Prefix, login_page.Redirect)
	return mux
}

func RemoveTrailingSlash(url string) string {
	if url[len(url)-1] == '/' {
		return url[:len(url)-1]
	}
	return url
}
