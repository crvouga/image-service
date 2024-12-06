package loginLinkDb

import "imageresizerservice/users/loginEmailLink/loginLink"

type LoginLinkDb interface {
	GetById(id string) (*loginLink.LoginLink, error)
	Upsert(loginLink.LoginLink) error
}
