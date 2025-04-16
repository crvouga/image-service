package linkDb

import (
	"imageresizerservice/app/ctx/sessionID"
	"imageresizerservice/app/users/login/link"
	"imageresizerservice/app/users/login/link/linkID"
	"imageresizerservice/library/uow"
)

type LinkDb interface {
	GetByLinkID(id linkID.LinkID) (*link.Link, error)
	GetBySessionID(sessionID sessionID.SessionID) ([]*link.Link, error)
	Upsert(uow *uow.Uow, link link.Link) error
}
