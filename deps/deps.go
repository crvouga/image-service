package deps

import (
	"image-resizer-service/email/send_email"
	"image-resizer-service/login/login_link/login_link_db"
)

type Deps struct {
	SendEmail   send_email.SendEmail
	LoginLinkDb login_link_db.LoginLinkDb
}
