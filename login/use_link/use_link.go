package use_link

import (
	"errors"

	"imageresizerservice.com/deps"
	"imageresizerservice.com/login/login_link"
	"imageresizerservice.com/login/login_routes"

	"net/http"
	"strings"
)

func Router(mux *http.ServeMux, d *deps.Deps) {
	mux.HandleFunc(login_routes.UseLink, Respond(d))
}

func Respond(d *deps.Deps) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		login_link_id := strings.TrimSpace(r.URL.Query().Get("loginLinkId"))

		err := use_login_link(d, login_link_id)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		http.Redirect(w, r, login_routes.LoginPage, http.StatusSeeOther)

	}
}

func use_login_link(d *deps.Deps, login_link_id string) error {
	cleaned := strings.TrimSpace(login_link_id)

	if cleaned == "" {
		return errors.New("login link id is required")
	}

	found, err := d.LoginLinkDb.GetById(cleaned)

	if err != nil {
		return err
	}

	if found == nil {
		return errors.New("no record of login link found")
	}

	if login_link.WasUsed(found) {
		return errors.New("login link has already been used")
	}

	return nil
}
