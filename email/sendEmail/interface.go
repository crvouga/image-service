package sendEmail

import "imageresizerservice/email/email"

type SendEmail interface {
	SendEmail(email email.Email) error
}
