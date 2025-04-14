package loginWithEmailLink

import (
	"net/http"

	"imageresizerservice/deps"
	"imageresizerservice/users/loginWithEmailLink/sendLink/sendLink"
	"imageresizerservice/users/loginWithEmailLink/useLink/useLink"
)

func Router(mux *http.ServeMux, d *deps.Deps) {
	sendLink.Router(mux, d)
	useLink.Router(mux, d)
}
