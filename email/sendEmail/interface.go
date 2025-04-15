package sendEmail

import (
	"imageresizerservice/email/email"
	"imageresizerservice/uow"
)

type SendEmail interface {
	SendEmail(uow *uow.Uow, email email.Email) error
}
