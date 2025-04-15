package httpRequest

import "net/http"

func GetRequestBaseURL(r *http.Request) string {
	scheme := "http"
	if r.TLS != nil || r.Header.Get("X-Forwarded-Proto") == "https" {
		scheme = "https"
	}

	// r.Host includes both hostname and port if present
	// e.g., "localhost:8080" or "example.com"
	return scheme + "://" + r.Host
}
