package reqCtx

import (
	"imageresizerservice/library/httpRequest"
	"imageresizerservice/library/id"
	"log/slog"
	"net/http"
)

type ReqCtx struct {
	BaseURL   string
	SessionID string
	TraceID   string
	Logger    *slog.Logger
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

func getSessionID(r *http.Request) string {
	cookie, err := r.Cookie("sessionID")
	if err != nil {
		return id.Gen()
	}
	return cookie.Value
}

func FromHttpRequest(r *http.Request) ReqCtx {
	sessionID := getSessionID(r)

	traceID := getTraceID(r)

	baseURL := httpRequest.GetRequestBaseURL(r)

	logger := slog.Default().With(
		slog.String("traceID", traceID),
	)

	reqCtx := ReqCtx{
		BaseURL:   baseURL,
		SessionID: sessionID,
		TraceID:   traceID,
		Logger:    logger,
	}

	return reqCtx
}

// WithSessionID is a middleware that ensures a sessionID cookie exists.
// If the cookie is missing, it creates a new one with a random UUID.
func WithSessionID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if sessionID cookie exists
		_, err := r.Cookie("sessionID")
		if err == http.ErrNoCookie {
			// Generate a new session ID (using a UUID would be ideal)
			sessionID := id.Gen()

			// Create a new cookie
			cookie := &http.Cookie{
				Name:     "sessionID",
				Value:    sessionID,
				Path:     "/",
				HttpOnly: true,
				MaxAge:   86400 * 30, // 30 days
			}

			// Set the cookie in the response
			http.SetCookie(w, cookie)

			// Add the cookie to the current request so that handlers
			// using this request can access it
			r.AddCookie(cookie)
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
