package sendEmail

import (
	"imageresizerservice/library/email/email"
	"imageresizerservice/library/uow"
)

type SendEmail interface {
	SendEmail(uow *uow.Uow, email email.Email) error
}
