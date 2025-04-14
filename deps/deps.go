package deps

import (
	"imageresizerservice/email/sendEmail"
	"imageresizerservice/uow"
	"imageresizerservice/users/loginEmailLink/loginLink/loginLinkDb"
)

type Deps struct {
	SendEmail   sendEmail.SendEmail
	LoginLinkDb loginLinkDb.LoginLinkDb
	UowFactory  uow.UowFactory
}
