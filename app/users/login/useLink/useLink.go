package useLink

import (
	"errors"
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/ctx/reqCtx"
	"imageresizerservice/app/home/homeRoutes"
	"imageresizerservice/app/ui/confirmationPage"
	"imageresizerservice/app/ui/errorPage"
	"imageresizerservice/app/ui/successPage"
	"imageresizerservice/app/users/login/link"
	"imageresizerservice/app/users/login/link/linkID"
	"imageresizerservice/app/users/login/loginRoutes"
	"imageresizerservice/app/users/userAccount"
	"imageresizerservice/app/users/userID"
	"imageresizerservice/app/users/userSession"
	"imageresizerservice/app/users/userSession/userSessionID"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func Router(mux *http.ServeMux, ac *appCtx.AppCtx) {
	mux.HandleFunc(loginRoutes.UseLinkPage, Respond(ac))
}

type Data struct {
	Action string
	LinkID linkID.LinkID
}

func Respond(ac *appCtx.AppCtx) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			confirmationPage.ConfirmationPage{
				Headline:    "Login",
				Body:        "Use this link to login to your account",
				ConfirmURL:  loginRoutes.UseLinkPage,
				ConfirmText: "Login",
				CancelURL:   homeRoutes.HomePage,
				CancelText:  "Cancel",
				HiddenForm: map[string]string{
					"linkID": r.URL.Query().Get("linkID"),
				},
			}.Render(w, r)
			return
		}

		if r.Method == http.MethodPost {
			// Handle POST request - process the link
			rc := reqCtx.FromHttpRequest(ac, r)

			if err := r.ParseForm(); err != nil {
				errorPage.New(errors.New("failed to parse form")).Redirect(w, r)
				return
			}

			linkID := strings.TrimSpace(r.FormValue("linkID"))

			if err := UseLink(ac, &rc, linkID); err != nil {
				errorPage.New(err).Redirect(w, r)
				return
			}

			successPage.New("Login successful", homeRoutes.HomePage, "Home").Redirect(w, r)
			return
		}

		errorPage.New(errors.New("method not allowed")).Redirect(w, r)
	}
}

func UseLink(ac *appCtx.AppCtx, rc *reqCtx.ReqCtx, maybeLinkID string) error {
	logger := rc.Logger.With(slog.String("operation", "UseLink"))

	logger.Info("Starting login with email link process", "linkID", maybeLinkID)

	cleaned := strings.TrimSpace(maybeLinkID)

	if cleaned == "" {
		logger.Warn("Empty link ID provided")
		return errors.New("login link id is required")
	}

	logger.Info("Fetching link from database", "linkID", cleaned)

	linkID := linkID.New(cleaned)

	found, err := ac.LinkDB.GetByLinkID(linkID)

	if err != nil {
		logger.Error("Error fetching link", "error", err.Error())
		return newDatabaseError(err)
	}

	if found == nil {
		logger.Warn("No link found with provided ID", "linkID", cleaned)
		return errors.New("no record of login link found")
	}

	if link.WasUsed(found) {
		logger.Warn("Link has already been used", "linkID", cleaned)
		return errors.New("login link has already been used")
	}

	logger.Info("Beginning database transaction")
	uow, err := ac.UowFactory.Begin()

	if err != nil {
		logger.Error("Failed to begin transaction", "error", err.Error())
		return newDatabaseError(err)
	}

	defer uow.Rollback()

	logger.Info("Marking link as used", "linkID", cleaned)
	marked := link.MarkAsUsed(*found)

	if err := ac.LinkDB.Upsert(uow, marked); err != nil {
		logger.Error("Failed to mark link as used", "error", err.Error())
		return newDatabaseError(err)
	}

	logger.Info("Looking up user account by email", "email", found.EmailAddress)
	account, err := ac.UserAccountDB.GetByEmailAddress(found.EmailAddress)

	if err != nil {
		logger.Error("Error looking up user account", "error", err.Error())
		return newDatabaseError(err)
	}

	if account == nil {
		logger.Info("Creating new user account", "email", found.EmailAddress)
		account = &userAccount.UserAccount{
			UserID:       userID.Gen(),
			EmailAddress: found.EmailAddress,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}
	} else {
		logger.Info("Found existing user account", "userID", account.UserID)
	}

	if err := ac.UserAccountDB.Upsert(uow, *account); err != nil {
		logger.Error("Failed to save user account", "error", err.Error())
		return newDatabaseError(err)
	}

	logger.Info("Creating new user session", "userID", account.UserID)
	sessionNew := userSession.UserSession{
		ID:        userSessionID.Gen(),
		UserID:    account.UserID,
		CreatedAt: time.Now(),
		SessionID: rc.SessionID,
	}

	if err := ac.UserSessionDB.Upsert(uow, sessionNew); err != nil {
		logger.Error("Failed to create user session", "error", err.Error())
		return newDatabaseError(err)
	}

	logger.Info("Committing transaction")
	if err := uow.Commit(); err != nil {
		logger.Error("Failed to commit transaction", "error", err.Error())
		return newDatabaseError(err)
	}

	logger.Info("Successfully completed login process", "userID", account.UserID)
	return nil
}

func newDatabaseError(err error) error {
	return errors.New("database error: " + err.Error())
}

func ToUrl(rc *reqCtx.ReqCtx, linkID linkID.LinkID) string {
	path := ToPath(linkID)
	u, _ := url.Parse(rc.BaseURL + path)
	return u.String()
}

func ToPath(linkID linkID.LinkID) string {
	u, _ := url.Parse(loginRoutes.UseLinkPage)
	q := u.Query()
	q.Set("linkID", string(linkID))
	u.RawQuery = q.Encode()
	return u.String()
}
