package reqCtx

import (
	"imageresizerservice/app/ctx/appContext"
	"imageresizerservice/app/users/userAccount"
	"imageresizerservice/app/users/userSession"
	"imageresizerservice/library/httpRequest"
	"imageresizerservice/library/sessionID"
	"imageresizerservice/library/traceID"
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

func getUserSession(ac *appContext.AppCtx, sessionID sessionID.SessionID) *userSession.UserSession {
	userSession, err := ac.UserSessionDB.GetBySessionID(sessionID)
	if err != nil {
		return nil
	}
	if userSession == nil {
		return nil
	}
	return userSession
}

func getUserAccount(ac *appContext.AppCtx, userSessionInst *userSession.UserSession) *userAccount.UserAccount {
	if userSessionInst == nil {
		return nil
	}
	userAccount, err := ac.UserAccountDB.GetByUserID(userSessionInst.UserID)
	if err != nil {
		return nil
	}
	if userAccount == nil {
		return nil
	}
	return userAccount
}

// FromHttpRequest creates a new ReqCtx from an HTTP request.
func FromHttpRequest(ac *appContext.AppCtx, r *http.Request) ReqCtx {
	sessionIDInst := sessionID.FromSessionIDCookie(r)

	traceIDInst := traceID.FromHttpRequest(r)

	baseURL := httpRequest.GetRequestBaseURL(r)

	logger := slog.Default().With(
		slog.String("traceID", string(traceIDInst)),
	)

	userSessionInst := getUserSession(ac, sessionIDInst)

	userAccountInst := getUserAccount(ac, userSessionInst)

	reqCtxInst := ReqCtx{
		BaseURL:     baseURL,
		SessionID:   sessionIDInst,
		TraceID:     traceIDInst,
		Logger:      logger,
		UserSession: userSessionInst,
		UserAccount: userAccountInst,
	}

	return reqCtxInst
}
