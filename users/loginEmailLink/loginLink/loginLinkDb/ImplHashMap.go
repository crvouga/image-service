package loginLinkDb

import (
	"imageresizerservice/users/loginEmailLink/loginLink"
	"time"
)

type ImplHashMap struct {
	LoginLinks map[string]loginLink.LoginLink
}

func NewImplHashMap() ImplHashMap {
	return ImplHashMap{
		LoginLinks: make(map[string]loginLink.LoginLink),
	}
}

var _ LoginLinkDb = ImplHashMap{}

func (db ImplHashMap) GetById(id string) (*loginLink.LoginLink, error) {
	time.Sleep(time.Second)

	found, ok := db.LoginLinks[id]
	if !ok {
		return nil, nil
	}
	return &found, nil
}

func (db ImplHashMap) Upsert(l loginLink.LoginLink) error {
	time.Sleep(time.Second)

	db.LoginLinks[l.Id] = l
	return nil
}
