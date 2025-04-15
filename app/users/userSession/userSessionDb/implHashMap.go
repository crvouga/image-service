package userSessionDb

import (
	"imageresizerservice/users/userSession"
	"time"
)

type ImplHashMap struct {
	UserSessions map[string]userSession.UserSession
}

func NewImplHashMap() ImplHashMap {
	return ImplHashMap{
		UserSessions: make(map[string]userSession.UserSession),
	}
}

var _ UserSessionDb = ImplHashMap{}

func (db ImplHashMap) GetById(id string) (*userSession.UserSession, error) {
	time.Sleep(time.Second)

	found, ok := db.UserSessions[id]
	if !ok {
		return nil, nil
	}
	return &found, nil
}

func (db ImplHashMap) Upsert(l userSession.UserSession) error {
	time.Sleep(time.Second)

	db.UserSessions[l.Id] = l
	return nil
}
