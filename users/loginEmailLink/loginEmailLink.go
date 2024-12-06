package loginEmailLink

import (
	"net/http"

	"imageresizerservice/deps"
	"imageresizerservice/static"
	"imageresizerservice/users/loginEmailLink/loginPage"
	"imageresizerservice/users/loginEmailLink/routes"
	"imageresizerservice/users/loginEmailLink/sendLink"
	"imageresizerservice/users/loginEmailLink/sentLinkPage"
)

func Router(mux *http.ServeMux, d *deps.Deps) {
	loginPage.Router(mux)
	sendLink.Router(mux, d)
	sentLinkPage.Router(mux)
	routerFallback(mux)
}

func routerFallback(mux *http.ServeMux) {
	mux.HandleFunc(RemoveTrailingSlash(routes.Prefix), RespondNotFound)
	mux.HandleFunc(routes.Prefix, RespondNotFound)
}

func RespondNotFound(w http.ResponseWriter, r *http.Request) {
	err := static.ServeStaticAssets(w, r)
	if err != nil {
		loginPage.Redirect(w, r)
	}
}

func RemoveTrailingSlash(url string) string {
	if url[len(url)-1] == '/' {
		return url[:len(url)-1]
	}
	return url
}
