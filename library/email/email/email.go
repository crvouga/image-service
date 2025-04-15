package email

import (
	"imageresizerservice/library/email/emailAddress"
)

type Email struct {
	To      emailAddress.EmailAddress
	From    emailAddress.EmailAddress
	Subject string
	Body    string
}
