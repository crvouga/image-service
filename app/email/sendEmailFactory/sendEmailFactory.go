package sendEmailFactory

import (
	"imageresizerservice/app/ctx/reqCtx"
	"imageresizerservice/library/email/sendEmail"
)

func IsConfigured() bool {
	return false
}

// FromReqCtx returns a SendEmail implementation based on the request context
// Currently it returns the default implementation, but this can be extended
// to provide different implementations based on context information
func FromReqCtx(ctx *reqCtx.ReqCtx) sendEmail.SendEmail {

	fake := sendEmail.NewFake()

	return fake
}
