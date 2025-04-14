package deps

import (
	"imageresizerservice/email/emailOutbox"
	"imageresizerservice/email/sendEmail"
	"imageresizerservice/keyValueDb"
	"imageresizerservice/uow"
	"imageresizerservice/users/loginEmailLink/loginLink/loginLinkDb"
)

type Deps struct {
	SendEmail   sendEmail.SendEmail
	LoginLinkDb loginLinkDb.LoginLinkDb
	UowFactory  uow.UowFactory
	EmailOutbox emailOutbox.EmailOutbox
	KeyValueDb  keyValueDb.KeyValueDb
}
