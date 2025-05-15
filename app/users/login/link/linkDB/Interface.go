package linkDB

import (
	"imageService/app/users/login/link"
	"imageService/app/users/login/link/linkID"
	"imageService/library/sessionID"
	"imageService/library/uow"
)

type LinkDB interface {
	GetByLinkID(id linkID.LinkID) (*link.Link, error)
	GetBySessionID(sessionID sessionID.SessionID) ([]*link.Link, error)
	Upsert(uow *uow.Uow, link link.Link) error
}
