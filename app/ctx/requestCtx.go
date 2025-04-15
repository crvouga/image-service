package ctx

type RequestCtx struct {
	SessionID string
	// Add other request-specific data here as needed
}

func NewRequestCtx(sessionID string) RequestCtx {
	return RequestCtx{
		SessionID: sessionID,
	}
}
