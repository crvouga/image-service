package linkDb

import (
	"imageresizerservice/app/users/loginWithEmailLink/link"
	"imageresizerservice/app/users/loginWithEmailLink/link/linkID"
	"imageresizerservice/library/uow"
)

type LinkDb interface {
	GetById(id linkID.LinkID) (*link.Link, error)
	Upsert(uow *uow.Uow, link link.Link) error
}
