package sendEmail

import (
	"imageresizerservice/email/email"
	"log"
	"time"
)

type ImplFake struct{}

func (f *ImplFake) SendEmail(email email.Email) error {
	log.Printf("Sending email to %s with subject %s and body %s", email.To, email.Subject, email.Body)
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	return nil
}

var _ SendEmail = (*ImplFake)(nil)
