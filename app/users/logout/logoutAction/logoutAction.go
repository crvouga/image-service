package logoutAction

import (
	"errors"
	"net/http"

	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/ctx/reqCtx"
	"imageresizerservice/app/ui/errorPage"
	"imageresizerservice/app/users/logout/logoutRoutes"
)

func Router(mux *http.ServeMux, ac *appCtx.AppCtx) {
	mux.HandleFunc(logoutRoutes.LogoutAction, Respond(ac))
}

func Respond(ac *appCtx.AppCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			errorPage.New(errors.New("method not allowed")).Redirect(w, r)
			return
		}

		req := reqCtx.FromHttpRequest(ac, r)

		if err := Logout(ac, &req); err != nil {
			errorPage.New(errors.New("failed to logout")).Redirect(w, r)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func Logout(ac *appCtx.AppCtx, rc *reqCtx.ReqCtx) error {
	if rc.UserSession == nil {
		return nil
	}

	uow, err := ac.UowFactory.Begin()
	if err != nil {
		return err
	}
	defer uow.Rollback()

	if err := ac.UserSessionDB.ZapBySessionID(uow, rc.SessionID); err != nil {
		return err
	}

	if err := uow.Commit(); err != nil {
		return err
	}

	return nil
}
