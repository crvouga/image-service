package reqCtx

import (
	"imageresizerservice/app/ctx/appCtx"
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

func getUserSession(appCtx *appCtx.AppCtx, sessionID sessionID.SessionID) *userSession.UserSession {
	userSession, err := appCtx.UserSessionDB.GetBySessionID(sessionID)
	if err != nil {
		return nil
	}
	if userSession == nil {
		return nil
	}
	return userSession
}

func getUserAccount(appCtx *appCtx.AppCtx, userSessionInst *userSession.UserSession) *userAccount.UserAccount {
	if userSessionInst == nil {
		return nil
	}
	userAccount, err := appCtx.UserAccountDB.GetByUserID(userSessionInst.UserID)
	if err != nil {
		return nil
	}
	if userAccount == nil {
		return nil
	}
	return userAccount
}

// FromHttpRequest creates a new ReqCtx from an HTTP request.
func FromHttpRequest(appCtx *appCtx.AppCtx, r *http.Request) ReqCtx {
	sessionIDInst := sessionID.FromSessionIDCookie(r)

	traceIDInst := traceID.FromHttpRequest(r)

	baseURL := httpRequest.GetRequestBaseURL(r)

	logger := slog.Default().With(
		slog.String("traceID", string(traceIDInst)),
	)

	userSessionInst := getUserSession(appCtx, sessionIDInst)

	userAccountInst := getUserAccount(appCtx, userSessionInst)

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
