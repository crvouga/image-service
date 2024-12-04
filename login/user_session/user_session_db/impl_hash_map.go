package user_session_db

import (
	"image-resizer-service/login/user_session"
	"time"
)

type ImplHashMap struct {
	UserSessions map[string]user_session.UserSession
}

func NewImplHashMap() ImplHashMap {
	return ImplHashMap{
		UserSessions: make(map[string]user_session.UserSession),
	}
}

var _ UserSessionDb = ImplHashMap{}

func (db ImplHashMap) GetById(id string) (*user_session.UserSession, error) {
	time.Sleep(time.Second)

	found, ok := db.UserSessions[id]
	if !ok {
		return nil, nil
	}
	return &found, nil
}

func (db ImplHashMap) Upsert(l user_session.UserSession) error {
	time.Sleep(time.Second)

	db.UserSessions[l.Id] = l
	return nil
}
