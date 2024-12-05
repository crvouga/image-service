package login_link_db

import (
	"time"

	"imageresizerservice.com/login/login_link"
)

type ImplHashMap struct {
	LoginLinks map[string]login_link.LoginLink
}

func NewImplHashMap() ImplHashMap {
	return ImplHashMap{
		LoginLinks: make(map[string]login_link.LoginLink),
	}
}

var _ LoginLinkDb = ImplHashMap{}

func (db ImplHashMap) GetById(id string) (*login_link.LoginLink, error) {
	time.Sleep(time.Second)

	found, ok := db.LoginLinks[id]
	if !ok {
		return nil, nil
	}
	return &found, nil
}

func (db ImplHashMap) Upsert(l login_link.LoginLink) error {
	time.Sleep(time.Second)

	db.LoginLinks[l.Id] = l
	return nil
}
