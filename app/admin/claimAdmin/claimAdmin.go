package claimAdmin

import (
	"imageresizerservice/app/admin/adminRoutes"
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/ctx/reqCtx"
	"imageresizerservice/app/home/homeRoutes"
	"imageresizerservice/app/ui/page"
	"imageresizerservice/app/users/userAccount"
	"imageresizerservice/app/users/userAccount/userRole"
	"imageresizerservice/library/static"
	"net/http"
)

func Router(mux *http.ServeMux, ac *appCtx.AppCtx) {
	mux.HandleFunc(adminRoutes.ClaimAdmin, Respond(ac))
}

type Data struct {
	HomeURL string
}

func Respond(ac *appCtx.AppCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rc := reqCtx.FromHttpRequest(ac, r)

		admins, err := ac.UserAccountDB.GetByRole(userRole.Admin)
		if err != nil {
			http.Error(w, "Failed to get admins", http.StatusInternalServerError)
			return
		}

		if len(admins) > 0 {
			http.Redirect(w, r, homeRoutes.HomePage, http.StatusSeeOther)
			return
		}

		// Handle POST request to claim admin
		if r.Method == http.MethodPost && rc.UserAccount != nil {
			handlePost(ac, &rc, w, r)
			return
		}

		// For GET requests or if user is not logged in
		handleGet(w, r)
	}
}

func handlePost(ac *appCtx.AppCtx, rc *reqCtx.ReqCtx, w http.ResponseWriter, r *http.Request) {
	uow, err := ac.UowFactory.Begin()
	if err != nil {
		http.Error(w, "Failed to start transaction", http.StatusInternalServerError)
		return
	}
	defer uow.Rollback()

	// Update the user to be an admin
	userAccountNew := userAccount.UserAccount{
		UserID:       rc.UserAccount.UserID,
		EmailAddress: rc.UserAccount.EmailAddress,
		CreatedAt:    rc.UserAccount.CreatedAt,
		UpdatedAt:    rc.UserAccount.UpdatedAt,
		Role:         userRole.Admin,
	}

	if err = ac.UserAccountDB.Upsert(uow, userAccountNew); err != nil {
		http.Error(w, "Failed to update user as admin", http.StatusInternalServerError)
		return
	}

	if err = uow.Commit(); err != nil {
		http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, homeRoutes.HomePage, http.StatusSeeOther)
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	data := Data{
		HomeURL: homeRoutes.HomePage,
	}

	page.Respond(data, static.GetSiblingPath("claimAdmin.html"))(w, r)
}
