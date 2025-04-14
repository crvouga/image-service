package deps

import (
	"imageresizerservice/email/emailOutbox"
	"imageresizerservice/email/sendEmail"
	"imageresizerservice/keyValueDb"
	"imageresizerservice/uow"
	"imageresizerservice/users/loginWithEmailLink/link/linkDb"
)

type Deps struct {
	SendEmail   sendEmail.SendEmail
	LinkDb      linkDb.LinkDb
	UowFactory  uow.UowFactory
	EmailOutbox emailOutbox.EmailOutbox
	KeyValueDb  keyValueDb.KeyValueDb
}
