package loginLinkDb

import (
	"imageresizerservice/uow"
	"imageresizerservice/users/loginEmailLink/loginLink"
)

type LoginLinkDb interface {
	GetById(id string) (*loginLink.LoginLink, error)
	Upsert(uow *uow.Uow, link loginLink.LoginLink) error
}
