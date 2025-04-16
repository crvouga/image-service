package sessionID

import (
	"imageresizerservice/library/id"
	"net/http"
)

type SessionID string

func Gen() SessionID {
	return SessionID(id.Gen())
}

func New(id string) SessionID {
	return SessionID(id)
}

// WithSessionIDCookie is a middleware that ensures a sessionID cookie exists.
// If the cookie is missing, it creates a new one with a random UUID.
func WithSessionIDCookie(next http.Handler) http.Handler {
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

func FromSessionIDCookie(r *http.Request) SessionID {
	cookie, err := r.Cookie("sessionID")
	if err != nil {
		return Gen()
	}
	return New(cookie.Value)
}
