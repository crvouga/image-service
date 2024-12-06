package sendEmail

import (
	"log"
	"time"
)

type FakeSendEmail struct{}

func (f *FakeSendEmail) SendEmail(to string, subject string, body string) error {
	log.Printf("Sending email to %s with subject %s and body %s", to, subject, body)
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	return nil
}

var _ SendEmail = (*FakeSendEmail)(nil)
