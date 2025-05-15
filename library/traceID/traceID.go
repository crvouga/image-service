package traceID

import (
	"imageService/library/id"
	"net/http"
)

type TraceID string

func Gen() TraceID {
	return TraceID(id.Gen())
}

func New(id string) TraceID {
	return TraceID(id)
}

const headerName = "x-trace-id"

// WithTraceID is a middleware that ensures a trace ID exists in the request.
// If the header is missing, it creates a new one with a random ID.
// It also attaches the same trace ID to the response.
func WithTraceIDHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if trace ID header exists
		traceID := r.Header.Get(headerName)
		if traceID == "" {
			// Generate a new trace ID
			traceID = id.Gen()

			// Add the trace ID to the current request
			r.Header.Set(headerName, traceID)
		}

		// Set the trace ID in the response header
		w.Header().Set(headerName, traceID)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func FromHttpRequest(r *http.Request) TraceID {
	traceID := r.Header.Get(headerName)
	if traceID == "" {
		return Gen()
	}
	return New(traceID)
}
