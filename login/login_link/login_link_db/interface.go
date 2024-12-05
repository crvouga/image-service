package login_link_db

import "imageresizerservice.com/login/login_link"

type LoginLinkDb interface {
	GetById(id string) (*login_link.LoginLink, error)
	Upsert(login_link.LoginLink) error
}
