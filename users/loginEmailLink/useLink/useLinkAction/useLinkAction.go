package useLinkAction

import (
	"errors"
	"imageresizerservice/deps"
	"net/http"
	"strings"

	"imageresizerservice/users/loginEmailLink/loginLink"
	"imageresizerservice/users/loginEmailLink/routes"
)

func Router(mux *http.ServeMux, d *deps.Deps) {
	mux.HandleFunc(routes.UseLinkAction, Respond(d))
}

func Respond(d *deps.Deps) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		loginLinkId := strings.TrimSpace(r.URL.Query().Get("loginLinkId"))

		err := useLoginLink(d, loginLinkId)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		http.Redirect(w, r, routes.SendLinkPage, http.StatusSeeOther)

	}
}

func useLoginLink(d *deps.Deps, loginLinkId string) error {
	cleaned := strings.TrimSpace(loginLinkId)

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

	if loginLink.WasUsed(found) {
		return errors.New("login link has already been used")
	}

	return nil
}
