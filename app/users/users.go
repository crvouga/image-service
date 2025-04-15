package users

import (
	"imageresizerservice/app/deps"
	"imageresizerservice/app/users/loginWithEmailLink"
	"net/http"
)

func Router(mux *http.ServeMux, d *deps.Deps) {
	loginWithEmailLink.Router(mux, d)
}
