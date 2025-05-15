package emailOutbox

import (
	"imageService/library/email/email"
	"imageService/library/uow"
)

type EmailOutbox interface {
	Add(uow *uow.Uow, email email.Email) error
	GetUnsentEmails() ([]email.Email, error)
	MarkAsSent(uow *uow.Uow, email email.Email) error
}
