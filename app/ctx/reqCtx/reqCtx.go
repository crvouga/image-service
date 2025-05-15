package reqCtx

import (
	"imageService/app/ctx/appCtx"
	"imageService/app/users/userAccount"
	"imageService/app/users/userSession"
	"imageService/library/httpRequest"
	"imageService/library/sessionID"
	"imageService/library/traceID"
	"log/slog"
	"net/http"
)

type ReqCtx struct {
	BaseURL     string
	SessionID   sessionID.SessionID
	TraceID     traceID.TraceID
	Logger      *slog.Logger
	UserSession *userSession.UserSession
	UserAccount *userAccount.UserAccount
}

func getUserSession(ac *appCtx.AppCtx, sessionID sessionID.SessionID) *userSession.UserSession {
	userSession, err := ac.UserSessionDB.GetBySessionID(sessionID)
	if err != nil {
		return nil
	}
	if userSession == nil {
		return nil
	}
	return userSession
}

func getUserAccount(ac *appCtx.AppCtx, userSessionVar *userSession.UserSession) *userAccount.UserAccount {
	if userSessionVar == nil {
		return nil
	}
	userAccountVar, err := ac.UserAccountDB.GetByUserID(userSessionVar.UserID)
	if err != nil {
		return nil
	}
	if userAccountVar == nil {
		return nil
	}
	return userAccountVar
}

// FromHttpRequest creates a new ReqCtx from an HTTP request.
func FromHttpRequest(ac *appCtx.AppCtx, r *http.Request) ReqCtx {
	sessionIDVar := sessionID.FromSessionIDCookie(r)

	traceIDVar := traceID.FromHttpRequest(r)

	baseURL := httpRequest.GetRequestBaseURL(r)

	logger := slog.Default().With(
		slog.String("traceID", string(traceIDVar)),
	)

	userSessionVar := getUserSession(ac, sessionIDVar)

	userAccountVar := getUserAccount(ac, userSessionVar)

	rc := ReqCtx{
		BaseURL:     baseURL,
		SessionID:   sessionIDVar,
		TraceID:     traceIDVar,
		Logger:      logger,
		UserSession: userSessionVar,
		UserAccount: userAccountVar,
	}

	return rc
}
