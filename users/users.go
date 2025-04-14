package users

import (
	"imageresizerservice/deps"
	"imageresizerservice/users/loginWithEmailLink"
	"net/http"
)

func Router(mux *http.ServeMux, d *deps.Deps) {
	loginWithEmailLink.Router(mux, d)
}
