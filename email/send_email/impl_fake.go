package send_email

import "log"

type FakeSendEmail struct{}

func (f *FakeSendEmail) SendEmail(to string, subject string, body string) error {
	log.Printf("Sending email to %s with subject %s and body %s", to, subject, body)
	return nil
}
