package useLinkAction

import (
	"errors"
	"net/http"
	"strings"

	"imageresizerservice/app/ctx"
	"imageresizerservice/app/users/loginWithEmailLink/link"
	"imageresizerservice/app/users/loginWithEmailLink/routes"
	"imageresizerservice/app/users/loginWithEmailLink/useLink/useLinkErrorPage"
	"imageresizerservice/app/users/loginWithEmailLink/useLink/useLinkSuccessPage"
)

func Router(mux *http.ServeMux, ctx *ctx.Ctx) {
	mux.HandleFunc(routes.UseLinkAction, Respond(ctx))
}

func Respond(ctx *ctx.Ctx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if err := r.ParseForm(); err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		linkId := strings.TrimSpace(r.FormValue("linkId"))

		err := UseLink(ctx, linkId)

		if err != nil {
			useLinkErrorPage.Redirect(w, r, err.Error())
			return
		}

		useLinkSuccessPage.Redirect(w, r)

	}
}

func UseLink(ctx *ctx.Ctx, linkId string) error {
	cleaned := strings.TrimSpace(linkId)

	if cleaned == "" {
		return errors.New("login link id is required")
	}

	found, err := ctx.LinkDb.GetById(cleaned)

	if err != nil {
		return errors.New("database error: " + err.Error())
	}

	if found == nil {
		return errors.New("no record of login link found")
	}

	if link.WasUsed(found) {
		return errors.New("login link has already been used")
	}

	return nil
}
