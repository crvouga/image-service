package login

import (
	"net/http"

	"imageresizerservice.com/deps"
	"imageresizerservice.com/login/login_page"
	"imageresizerservice.com/login/login_routes"
	"imageresizerservice.com/login/send_link"
	"imageresizerservice.com/login/sent_link_page"
	"imageresizerservice.com/static"
)

func Router(mux *http.ServeMux, d *deps.Deps) {
	login_page.Router(mux)
	send_link.Router(mux, d)
	sent_link_page.Router(mux)
	routerFallback(mux)
}

func routerFallback(mux *http.ServeMux) {
	mux.HandleFunc(RemoveTrailingSlash(login_routes.Prefix), RespondNotFound)
	mux.HandleFunc(login_routes.Prefix, RespondNotFound)
}

func RespondNotFound(w http.ResponseWriter, r *http.Request) {
	err := static.ServeStaticAssets(w, r)
	if err != nil {
		login_page.Redirect(w, r)
	}
}

func RemoveTrailingSlash(url string) string {
	if url[len(url)-1] == '/' {
		return url[:len(url)-1]
	}
	return url
}
