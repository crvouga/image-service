package sendEmail

import (
	"imageresizerservice/library/email/email"
	"imageresizerservice/library/uow"
	"log"
)

type ImplFake struct{}

func (f *ImplFake) SendEmail(uow *uow.Uow, email email.Email) error {
	log.Printf("Sending email to %s with subject %s and body %s", email.To, email.Subject, email.Body)
	return nil
}

var _ SendEmail = (*ImplFake)(nil)
