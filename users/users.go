package users

import (
	"imageresizerservice/deps"
	"imageresizerservice/users/loginEmailLink"
	"net/http"
)

func Router(mux *http.ServeMux, d *deps.Deps) {
	loginEmailLink.Router(mux, d)
}
