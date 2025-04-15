package reqCtx

import (
	"imageresizerservice/library/id"
	"net/http"
)

type ReqCtx struct {
	SessionID string
}

func New(sessionID string) ReqCtx {
	return ReqCtx{
		SessionID: sessionID,
	}
}

// FromHttpRequest extracts the ReqCtx from an HTTP request
// by retrieving the session ID from the request's cookies.
func FromHttpRequest(r *http.Request) (ReqCtx, error) {
	cookie, err := r.Cookie("sessionID")
	if err != nil {
		return ReqCtx{}, err
	}

	return New(cookie.Value), nil
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
