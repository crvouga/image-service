package useLinkAction

import (
	"errors"
	"net/http"
	"strings"

	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/users/loginWithEmailLink/link"
	"imageresizerservice/app/users/loginWithEmailLink/routes"
	"imageresizerservice/app/users/loginWithEmailLink/useLink/useLinkErrorPage"
	"imageresizerservice/app/users/loginWithEmailLink/useLink/useLinkSuccessPage"
)

func Router(mux *http.ServeMux, appCtx *appCtx.AppCtx) {
	mux.HandleFunc(routes.UseLinkAction, Respond(appCtx))
}

func Respond(appCtx *appCtx.AppCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if err := r.ParseForm(); err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		linkId := strings.TrimSpace(r.FormValue("linkId"))

		err := UseLink(appCtx, linkId)

		if err != nil {
			useLinkErrorPage.Redirect(w, r, err.Error())
			return
		}

		useLinkSuccessPage.Redirect(w, r)

	}
}

func UseLink(appCtx *appCtx.AppCtx, linkId string) error {
	cleaned := strings.TrimSpace(linkId)

	if cleaned == "" {
		return errors.New("login link id is required")
	}

	found, err := appCtx.LinkDb.GetById(cleaned)

	if err != nil {
		return newDatabaseError(err)
	}

	if found == nil {
		return errors.New("no record of login link found")
	}

	if link.WasUsed(found) {
		return errors.New("login link has already been used")
	}

	uow, err := appCtx.UowFactory.Begin()

	if err != nil {
		return newDatabaseError(err)
	}

	defer uow.Rollback()

	marked := link.MarkAsUsed(*found)

	if err := appCtx.LinkDb.Upsert(uow, marked); err != nil {
		return newDatabaseError(err)
	}

	if err := uow.Commit(); err != nil {
		return newDatabaseError(err)
	}

	return nil
}

func newDatabaseError(err error) error {
	return errors.New("database error: " + err.Error())
}
