package userSessionDB

import (
	"encoding/json"
	"imageresizerservice/app/ctx/sessionID"
	"imageresizerservice/app/users/userSession"
	"imageresizerservice/library/keyValueDB"
	"imageresizerservice/library/uow"
)

type ImplKeyValueDB struct {
	db             keyValueDB.KeyValueDB
	indexSessionID keyValueDB.KeyValueDB
}

var _ UserSessionDB = (*ImplKeyValueDB)(nil)

func NewImplKeyValueDB(db keyValueDB.KeyValueDB) *ImplKeyValueDB {
	return &ImplKeyValueDB{
		db:             keyValueDB.NewImplNamespaced(db, "userSession"),
		indexSessionID: keyValueDB.NewImplNamespaced(db, "userSessionIndexSessionID"),
	}
}

func (db *ImplKeyValueDB) GetBySessionID(sessionId sessionID.SessionID) (*userSession.UserSession, error) {
	userSessionId, err := db.indexSessionID.Get(string(sessionId))

	if err != nil {
		return nil, err
	}

	if userSessionId == nil {
		return nil, nil
	}

	value, err := db.db.Get(*userSessionId)

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

func (db *ImplKeyValueDB) Upsert(uow *uow.Uow, userSession userSession.UserSession) error {
	// Check if db is initialized to prevent nil pointer dereference
	if db.db == nil || db.indexSessionID == nil {
		return nil
	}

	jsonData, err := json.Marshal(userSession)

	if err != nil {
		return err
	}

	if err := db.db.Put(uow, string(userSession.ID), string(jsonData)); err != nil {
		return err
	}

	if err := db.indexSessionID.Put(uow, string(userSession.SessionID), string(userSession.ID)); err != nil {
		return err
	}

	return nil
}

func (db *ImplKeyValueDB) ZapBySessionID(uow *uow.Uow, sessionID sessionID.SessionID) error {
	return db.indexSessionID.Zap(uow, string(sessionID))
}
