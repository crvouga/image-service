package login

import (
	"image-resizer-service/deps"
	"image-resizer-service/login/login_page"
	"image-resizer-service/login/login_routes"
	"image-resizer-service/login/send_link"
	"image-resizer-service/login/sent_link_page"
	"image-resizer-service/static"
	"net/http"
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
