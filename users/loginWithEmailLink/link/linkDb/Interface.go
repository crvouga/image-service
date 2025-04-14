package linkDb

import (
	"imageresizerservice/uow"

	"imageresizerservice/users/loginWithEmailLink/link"
)

type LinkDb interface {
	GetById(id string) (*link.Link, error)
	Upsert(uow *uow.Uow, link link.Link) error
}
