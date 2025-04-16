package logoutAction

import (
	"net/http"

	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/ctx/reqCtx"
	"imageresizerservice/app/users/logout/logoutRoutes"
)

func Router(mux *http.ServeMux, appCtx *appCtx.AppCtx) {
	mux.HandleFunc(logoutRoutes.LogoutAction, Respond(appCtx))
}

func Respond(appCtx *appCtx.AppCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		req := reqCtx.FromHttpRequest(appCtx, r)

		if err := Logout(appCtx, &req); err != nil {
			http.Error(w, "Failed to logout", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func Logout(appCtx *appCtx.AppCtx, reqCtx *reqCtx.ReqCtx) error {
	if reqCtx.UserSession == nil {
		return nil
	}

	uow, err := appCtx.UowFactory.Begin()
	if err != nil {
		return err
	}
	defer uow.Rollback()

	if err := appCtx.UserSessionDb.ZapBySessionID(uow, reqCtx.SessionID); err != nil {
		return err
	}

	if err := uow.Commit(); err != nil {
		return err
	}

	return nil
}
