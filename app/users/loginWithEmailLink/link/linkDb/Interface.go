package linkDb

import (
	"imageresizerservice/app/users/loginWithEmailLink/link"
	"imageresizerservice/library/uow"
)

type LinkDb interface {
	GetById(id string) (*link.Link, error)
	Upsert(uow *uow.Uow, link link.Link) error
}
