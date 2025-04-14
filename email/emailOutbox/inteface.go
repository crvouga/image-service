package emailOutbox

import "imageresizerservice/email/email"

type EmailOutbox interface {
	Add(email email.Email) error
	GetUnsentEmails() ([]email.Email, error)
	MarkAsSent(email email.Email) error
}
