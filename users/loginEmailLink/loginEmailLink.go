package loginEmailLink

import (
	"net/http"

	"imageresizerservice/deps"
	"imageresizerservice/static"
	"imageresizerservice/users/loginEmailLink/routes"
	"imageresizerservice/users/loginEmailLink/sendLink/sendLinkAction"
	"imageresizerservice/users/loginEmailLink/sendLink/sendLinkPage"
	"imageresizerservice/users/loginEmailLink/sendLink/sentLinkPage"
)

func Router(mux *http.ServeMux, d *deps.Deps) {
	sendLinkPage.Router(mux)
	sendLinkAction.Router(mux, d)
	sentLinkPage.Router(mux)

	mux.HandleFunc(removeTrailingSlash(routes.Prefix), respondNotFound)
	mux.HandleFunc(routes.Prefix, respondNotFound)
}

func respondNotFound(w http.ResponseWriter, r *http.Request) {
	err := static.ServeStaticAssets(w, r)
	if err != nil {
		sendLinkPage.Redirect(w, r)
	}
}

func removeTrailingSlash(url string) string {
	if url[len(url)-1] == '/' {
		return url[:len(url)-1]
	}
	return url
}
