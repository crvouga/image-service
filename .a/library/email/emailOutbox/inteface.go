package emailOutbox

import (
	"imageresizerservice/library/email/email"
	"imageresizerservice/library/uow"
)

type EmailOutbox interface {
	Add(uow *uow.Uow, email email.Email) error
	GetUnsentEmails() ([]email.Email, error)
	MarkAsSent(uow *uow.Uow, email email.Email) error
}
