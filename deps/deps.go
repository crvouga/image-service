package deps

import "image-resizer-service/email/send_email"

type Deps struct {
	SendEmail send_email.SendEmail
}
