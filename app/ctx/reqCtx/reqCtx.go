package reqCtx

import (
	"net/http"
)

type ReqCtx struct {
	SessionID string
}

func NewReqCtx(sessionID string) ReqCtx {
	return ReqCtx{
		SessionID: sessionID,
	}
}

// GetReqCtxFromRequest extracts the ReqCtx from an HTTP request
// by retrieving the session ID from the request's cookies.
func GetReqCtxFromRequest(r *http.Request) (ReqCtx, error) {
	cookie, err := r.Cookie("sessionID")
	if err != nil {
		return ReqCtx{}, err
	}

	return NewReqCtx(cookie.Value), nil
}
