package sendEmail

import (
	"imageresizerservice/email/email"
	"log"
	"time"
)

type FakeSendEmail struct{}

func (f *FakeSendEmail) SendEmail(email email.Email) error {
	log.Printf("Sending email to %s with subject %s and body %s", email.To, email.Subject, email.Body)
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	return nil
}

var _ SendEmail = (*FakeSendEmail)(nil)
