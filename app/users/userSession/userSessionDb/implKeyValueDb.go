package userSessionDb

import (
	"encoding/json"
	"imageresizerservice/app/ctx/sessionID"
	"imageresizerservice/app/users/userSession"
	"imageresizerservice/library/keyValueDb"
	"imageresizerservice/library/uow"
)

type ImplKeyValueDb struct {
	db keyValueDb.KeyValueDb
}

var _ UserSessionDb = ImplKeyValueDb{}

func NewImplKeyValueDb(db keyValueDb.KeyValueDb) *ImplKeyValueDb {
	return &ImplKeyValueDb{
		db: keyValueDb.NewImplNamespaced(db, "userSession"),
	}
}

func (db ImplKeyValueDb) GetBySessionID(id sessionID.SessionID) (*userSession.UserSession, error) {
	value, err := db.db.Get(string(id))
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
	jsonData, err := json.Marshal(session)
	if err != nil {
		return err
	}

	return db.db.Put(uow, string(session.ID), string(jsonData))
}

var _ UserSessionDb = (*ImplKeyValueDb)(nil)
