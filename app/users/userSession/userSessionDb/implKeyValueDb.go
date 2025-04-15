package userSessionDb

import (
	"encoding/json"
	"imageresizerservice/app/users/userSession"
	"imageresizerservice/library/keyValueDb"
	"imageresizerservice/library/uow"
	"time"
)

type ImplKeyValueDb struct {
	Db keyValueDb.KeyValueDb
}

var _ UserSessionDb = ImplKeyValueDb{}

func (db ImplKeyValueDb) GetById(id string) (*userSession.UserSession, error) {
	time.Sleep(time.Second)

	value, err := db.Db.Get(id)
	if err != nil {
		return nil, err
	}

	if value == nil {
		return nil, nil
	}

	var session userSession.UserSession
	if err := json.Unmarshal([]byte(*value), &session); err != nil {
		return nil, err
	}

	return &session, nil
}

func (db ImplKeyValueDb) Upsert(uow *uow.Uow, session userSession.UserSession) error {
	time.Sleep(time.Second)

	jsonData, err := json.Marshal(session)
	if err != nil {
		return err
	}

	return db.Db.Put(uow, session.Id, string(jsonData))
}

var _ UserSessionDb = (*ImplKeyValueDb)(nil)
