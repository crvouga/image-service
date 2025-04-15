package loginWithEmailLink

import (
	"net/http"

	"imageresizerservice/app/deps"
	"imageresizerservice/app/users/loginWithEmailLink/sendLink/sendLink"
	"imageresizerservice/app/users/loginWithEmailLink/useLink/useLink"
)

func Router(mux *http.ServeMux, d *deps.Deps) {
	sendLink.Router(mux, d)
	useLink.Router(mux, d)
}
