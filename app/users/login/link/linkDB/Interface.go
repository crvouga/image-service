package linkDB

import (
	"imageresizerservice/app/users/login/link"
	"imageresizerservice/app/users/login/link/linkID"
	"imageresizerservice/library/sessionID"
	"imageresizerservice/library/uow"
)

type LinkDB interface {
	GetByLinkID(id linkID.LinkID) (*link.Link, error)
	GetBySessionID(sessionID sessionID.SessionID) ([]*link.Link, error)
	Upsert(uow *uow.Uow, link link.Link) error
}
