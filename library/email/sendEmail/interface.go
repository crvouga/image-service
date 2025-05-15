package sendEmail

import (
	"imageService/library/email/email"
	"imageService/library/uow"
)

type SendEmail interface {
	SendEmail(uow *uow.Uow, email email.Email) error
}
