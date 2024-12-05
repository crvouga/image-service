package deps

import (
	"imageresizerservice.com/email/send_email"
	"imageresizerservice.com/login/login_link/login_link_db"
)

type Deps struct {
	SendEmail   send_email.SendEmail
	LoginLinkDb login_link_db.LoginLinkDb
}
