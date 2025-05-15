package userSessionDB

import (
	"testing"
	"time"

	"imageService/app/users/userID"
	"imageService/app/users/userSession"
	"imageService/app/users/userSession/userSessionID"
	"imageService/library/keyValueDB"
	"imageService/library/sessionID"
	"imageService/library/sqlite"
	"imageService/library/uow"
)

type Fixture struct {
	UowFactory uow.UowFactory
	SessionDB  UserSessionDB
}

func newFixture() *Fixture {
	db := sqlite.New()

	return &Fixture{
		SessionDB:  NewImplKeyValueDB(keyValueDB.NewImplHashMap()),
		UowFactory: *uow.NewFactory(db),
	}
}

func Test_GetBySessionID(t *testing.T) {
	f := newFixture()
	uow, _ := f.UowFactory.Begin()

	// Create a session
	session := userSession.UserSession{
		ID:        userSessionID.Gen(),
		UserID:    userID.Gen(),
		SessionID: sessionID.Gen(),
		CreatedAt: time.Now(),
	}

	// Insert the session
	err := f.SessionDB.Upsert(uow, session)
	if err != nil {
		t.Errorf("Expected no error on insert, got %v", err)
	}

	// Get the session
	retrieved, err := f.SessionDB.GetBySessionID(session.SessionID)
	if err != nil {
		t.Errorf("Expected no error on retrieval, got %v", err)
	}

	if retrieved == nil {
		t.Fatal("Expected to retrieve session, got nil")
	}

	if retrieved.ID != session.ID {
		t.Errorf("Expected ID to be %s, got %s", session.ID, retrieved.ID)
	}

	if retrieved.UserID != session.UserID {
		t.Errorf("Expected UserID to be %s, got %s", session.UserID, retrieved.UserID)
	}

	uow.Commit()
}

func Test_GetByIDNonExistent(t *testing.T) {
	f := newFixture()

	// Try to get a session that doesn't exist
	nonexistentSessionID := sessionID.Gen()

	retrieved, err := f.SessionDB.GetBySessionID(nonexistentSessionID)

	if err != nil {
		t.Errorf("Expected no error for nonexistent session, got %v", err)
	}

	if retrieved != nil {
		t.Errorf("Expected nil for nonexistent session, got %+v", retrieved)
	}
}

func Test_UpsertNewSession(t *testing.T) {
	f := newFixture()
	uow, _ := f.UowFactory.Begin()

	// Create a session
	session := userSession.UserSession{
		ID:        userSessionID.Gen(),
		UserID:    userID.Gen(),
		SessionID: sessionID.Gen(),
		CreatedAt: time.Now(),
	}

	// Insert the session
	err := f.SessionDB.Upsert(uow, session)
	if err != nil {
		t.Errorf("Expected no error on insert, got %v", err)
	}

	// Verify it exists
	retrieved, err := f.SessionDB.GetBySessionID(session.SessionID)
	if err != nil {
		t.Errorf("Expected no error on retrieval, got %v", err)
	}

	if retrieved == nil {
		t.Fatal("Expected to retrieve session, got nil")
	}

	if retrieved.ID != session.ID {
		t.Errorf("Expected ID to be %s, got %s", session.ID, retrieved.ID)
	}

	uow.Commit()
}

func Test_UpsertUpdateSession(t *testing.T) {
	f := newFixture()
	uow, _ := f.UowFactory.Begin()

	// Create initial session
	session := userSession.UserSession{
		ID:        userSessionID.Gen(),
		UserID:    userID.Gen(),
		SessionID: sessionID.Gen(),
		CreatedAt: time.Now(),
	}

	// Insert the session
	err := f.SessionDB.Upsert(uow, session)
	if err != nil {
		t.Errorf("Expected no error on insert, got %v", err)
	}

	// Update the session
	updatedSession := userSession.UserSession{
		ID:        session.ID,
		UserID:    session.UserID,
		SessionID: session.SessionID,
		CreatedAt: session.CreatedAt,
		EndedAt:   time.Now(),
	}

	// Update the session
	err = f.SessionDB.Upsert(uow, updatedSession)
	if err != nil {
		t.Errorf("Expected no error on update, got %v", err)
	}

	// Verify it was updated
	retrieved, err := f.SessionDB.GetBySessionID(session.SessionID)
	if err != nil {
		t.Errorf("Expected no error on retrieval, got %v", err)
	}

	if retrieved == nil {
		t.Fatal("Expected to retrieve session, got nil")
	}

	if retrieved.EndedAt.IsZero() {
		t.Error("Expected EndedAt to be set, but it's zero")
	}

	uow.Commit()
}

func TestZapBySessionID(t *testing.T) {
	f := newFixture()
	uow, _ := f.UowFactory.Begin()

	// Create initial session
	session := userSession.UserSession{
		ID:        userSessionID.Gen(),
		UserID:    userID.Gen(),
		SessionID: sessionID.Gen(),
		CreatedAt: time.Now(),
	}

	// Insert the session
	err := f.SessionDB.Upsert(uow, session)
	if err != nil {
		t.Errorf("Expected no error on insert, got %v", err)
	}

	// Commit the session
	uow.Commit()

	// Start a new UOW for the zap operation
	uow, _ = f.UowFactory.Begin()

	// Zap the session
	err = f.SessionDB.ZapBySessionID(uow, session.SessionID)
	if err != nil {
		t.Errorf("Expected no error on zap, got %v", err)
	}

	// Commit the zap
	uow.Commit()

	// Verify it was removed
	retrieved, err := f.SessionDB.GetBySessionID(session.SessionID)
	if err != nil {
		t.Errorf("Expected no error on retrieval after zap, got %v", err)
	}

	if retrieved != nil {
		t.Errorf("Expected session to be removed, but it was retrieved: %+v", retrieved)
	}
}
