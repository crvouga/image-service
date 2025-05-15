package userSessionDB

import (
	"encoding/json"
	"imageService/app/users/userSession"
	"imageService/library/keyValueDB"
	"imageService/library/sessionID"
	"imageService/library/uow"
)

type ImplKeyValueDB struct {
	userSessions                  keyValueDB.KeyValueDB
	indexUserSessionIDBySessionID keyValueDB.KeyValueDB
}

var _ UserSessionDB = (*ImplKeyValueDB)(nil)

func NewImplKeyValueDB(db keyValueDB.KeyValueDB) *ImplKeyValueDB {
	return &ImplKeyValueDB{
		userSessions:                  keyValueDB.NewImplNamespaced(db, "userSession"),
		indexUserSessionIDBySessionID: keyValueDB.NewImplNamespaced(db, "userSession:index:userSessionIDBySessionID"),
	}
}

func (db *ImplKeyValueDB) GetBySessionID(sessionId sessionID.SessionID) (*userSession.UserSession, error) {
	userSessionId, err := db.indexUserSessionIDBySessionID.Get(string(sessionId))

	if err != nil {
		return nil, err
	}

	if userSessionId == nil {
		return nil, nil
	}

	value, err := db.userSessions.Get(*userSessionId)

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
	if db.userSessions == nil || db.indexUserSessionIDBySessionID == nil {
		return nil
	}

	jsonData, err := json.Marshal(userSession)

	if err != nil {
		return err
	}

	if err := db.userSessions.Put(uow, string(userSession.ID), string(jsonData)); err != nil {
		return err
	}

	if err := db.indexUserSessionIDBySessionID.Put(uow, string(userSession.SessionID), string(userSession.ID)); err != nil {
		return err
	}

	return nil
}

func (db *ImplKeyValueDB) ZapBySessionID(uow *uow.Uow, sessionID sessionID.SessionID) error {
	return db.indexUserSessionIDBySessionID.Zap(uow, string(sessionID))
}
