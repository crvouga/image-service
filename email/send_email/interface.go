package send_email

type SendEmail interface {
	SendEmail(to string, subject string, body string) error
}
