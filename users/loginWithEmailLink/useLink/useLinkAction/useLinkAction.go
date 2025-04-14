package useLinkAction

import (
	"errors"
	"imageresizerservice/deps"
	"net/http"
	"strings"

	"imageresizerservice/users/loginWithEmailLink/link"
	"imageresizerservice/users/loginWithEmailLink/routes"
)

func Router(mux *http.ServeMux, d *deps.Deps) {
	mux.HandleFunc(routes.UseLinkAction, Respond(d))
}

func Respond(d *deps.Deps) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		linkId := strings.TrimSpace(r.URL.Query().Get("linkId"))

		err := UseLink(d, linkId)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		http.Redirect(w, r, routes.SendLinkPage, http.StatusSeeOther)

	}
}

func UseLink(d *deps.Deps, linkId string) error {
	cleaned := strings.TrimSpace(linkId)

	if cleaned == "" {
		return errors.New("login link id is required")
	}

	found, err := d.LinkDb.GetById(cleaned)

	if err != nil {
		return err
	}

	if found == nil {
		return errors.New("no record of login link found")
	}

	if link.WasUsed(found) {
		return errors.New("login link has already been used")
	}

	return nil
}
