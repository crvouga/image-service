package emailOutbox

import (
	"imageresizerservice/email/email"
	"imageresizerservice/uow"
)

type EmailOutbox interface {
	Add(uow *uow.Uow, email email.Email) error
	GetUnsentEmails() ([]email.Email, error)
	MarkAsSent(uow *uow.Uow, email email.Email) error
}
