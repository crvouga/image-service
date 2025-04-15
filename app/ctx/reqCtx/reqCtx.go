package reqCtx

import (
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/ctx/sessionID"
	"imageresizerservice/app/users/userID"
	"imageresizerservice/library/httpRequest"
	"imageresizerservice/library/id"
	"log/slog"
	"net/http"
)

type ReqCtx struct {
	BaseURL   string
	SessionID sessionID.SessionID
	TraceID   string
	Logger    *slog.Logger
	UserID    *userID.UserID
}

// FromHttpRequest extracts the ReqCtx from an HTTP request
// by retrieving the session ID from the request's cookies.
func getTraceID(r *http.Request) string {
	traceID := r.Header.Get("x-trace-id")
	if traceID == "" {
		return id.Gen()
	}
	return traceID
}

func getUserID(appCtx *appCtx.AppCtx, sessionID sessionID.SessionID) *userID.UserID {
	userSession, err := appCtx.UserSessionDb.GetById(sessionID)
	if err != nil {
		return nil
	}
	if userSession == nil {
		return nil
	}
	return &userSession.UserID
}

func FromHttpRequest(appCtx *appCtx.AppCtx, r *http.Request) ReqCtx {
	sessionID := sessionID.FromSessionIDCookie(r)

	traceID := getTraceID(r)

	baseURL := httpRequest.GetRequestBaseURL(r)

	logger := slog.Default().With(
		slog.String("traceID", traceID),
	)

	userID := getUserID(appCtx, sessionID)

	reqCtx := ReqCtx{
		BaseURL:   baseURL,
		SessionID: sessionID,
		TraceID:   traceID,
		Logger:    logger,
		UserID:    userID,
	}

	return reqCtx
}
