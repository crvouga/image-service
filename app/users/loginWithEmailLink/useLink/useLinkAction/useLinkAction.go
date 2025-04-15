package useLinkAction

import (
	"errors"
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/ctx/reqCtx"
	"imageresizerservice/app/users/loginWithEmailLink/link"
	"imageresizerservice/app/users/loginWithEmailLink/routes"
	"imageresizerservice/app/users/loginWithEmailLink/useLink/useLinkErrorPage"
	"imageresizerservice/app/users/loginWithEmailLink/useLink/useLinkSuccessPage"
	"imageresizerservice/app/users/userAccount"
	"imageresizerservice/app/users/userID"
	"imageresizerservice/app/users/userSession"
	"imageresizerservice/app/users/userSession/userSessionID"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

func Router(mux *http.ServeMux, appCtx *appCtx.AppCtx) {
	mux.HandleFunc(routes.UseLinkAction, Respond(appCtx))
}

func Respond(appCtx *appCtx.AppCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		reqCtx := reqCtx.FromHttpRequest(appCtx, r)

		if err := r.ParseForm(); err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		linkId := strings.TrimSpace(r.FormValue("linkId"))

		if err := UseLink(appCtx, &reqCtx, linkId); err != nil {
			useLinkErrorPage.Redirect(w, r, err.Error())
			return
		}

		useLinkSuccessPage.Redirect(w, r)

	}
}

func UseLink(appCtx *appCtx.AppCtx, reqCtx *reqCtx.ReqCtx, linkId string) error {
	logger := reqCtx.Logger.With(slog.String("operation", "UseLink"))

	logger.Info("Starting login with email link process", "linkId", linkId)

	cleaned := strings.TrimSpace(linkId)

	if cleaned == "" {
		logger.Warn("Empty link ID provided")
		return errors.New("login link id is required")
	}

	logger.Info("Fetching link from database", "linkId", cleaned)
	found, err := appCtx.LinkDb.GetById(cleaned)

	if err != nil {
		logger.Error("Error fetching link", "error", err.Error())
		return newDatabaseError(err)
	}

	if found == nil {
		logger.Warn("No link found with provided ID", "linkId", cleaned)
		return errors.New("no record of login link found")
	}

	if link.WasUsed(found) {
		logger.Warn("Link has already been used", "linkId", cleaned)
		return errors.New("login link has already been used")
	}

	logger.Info("Beginning database transaction")
	uow, err := appCtx.UowFactory.Begin()

	if err != nil {
		logger.Error("Failed to begin transaction", "error", err.Error())
		return newDatabaseError(err)
	}

	defer uow.Rollback()

	logger.Info("Marking link as used", "linkId", cleaned)
	marked := link.MarkAsUsed(*found)

	if err := appCtx.LinkDb.Upsert(uow, marked); err != nil {
		logger.Error("Failed to mark link as used", "error", err.Error())
		return newDatabaseError(err)
	}

	logger.Info("Looking up user account by email", "email", found.EmailAddress)
	account, err := appCtx.UserAccountDb.GetByEmailAddress(found.EmailAddress)

	if err != nil {
		logger.Error("Error looking up user account", "error", err.Error())
		return newDatabaseError(err)
	}

	if account == nil {
		logger.Info("Creating new user account", "email", found.EmailAddress)
		account = &userAccount.UserAccount{
			ID:           userID.Gen(),
			EmailAddress: found.EmailAddress,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}
	} else {
		logger.Info("Found existing user account", "userId", account.ID)
	}

	if err := appCtx.UserAccountDb.Upsert(uow, *account); err != nil {
		logger.Error("Failed to save user account", "error", err.Error())
		return newDatabaseError(err)
	}

	logger.Info("Creating new user session", "userId", account.ID)
	sessionNew := userSession.UserSession{
		ID:        userSessionID.Gen(),
		UserID:    account.ID,
		CreatedAt: time.Now(),
		SessionID: reqCtx.SessionID,
	}

	if err := appCtx.UserSessionDb.Upsert(uow, sessionNew); err != nil {
		logger.Error("Failed to create user session", "error", err.Error())
		return newDatabaseError(err)
	}

	logger.Info("Committing transaction")
	if err := uow.Commit(); err != nil {
		logger.Error("Failed to commit transaction", "error", err.Error())
		return newDatabaseError(err)
	}

	logger.Info("Successfully completed login process", "userId", account.ID)
	return nil
}

func newDatabaseError(err error) error {
	return errors.New("database error: " + err.Error())
}
