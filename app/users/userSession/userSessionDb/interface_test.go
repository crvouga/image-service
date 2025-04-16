package userSessionDb

import (
	"testing"
	"time"

	"imageresizerservice/app/users/userSession"
	"imageresizerservice/library/keyValueDb"
	"imageresizerservice/library/sqlite"
	"imageresizerservice/library/uow"
)

type Fixture struct {
	UowFactory uow.UowFactory
	SessionDb  UserSessionDb
}

func newFixture() *Fixture {
	db := sqlite.New()

	return &Fixture{
		SessionDb:  ImplKeyValueDb{db: &keyValueDb.ImplHashMap{}},
		UowFactory: uow.UowFactory{Db: db},
	}
}

func Test_GetByID(t *testing.T) {
	f := newFixture()
	uow, _ := f.UowFactory.Begin()

	// Create a session
	session := userSession.UserSession{
		ID:        "test-id",
		UserID:    "user-123",
		SessionID: "session-456",
		CreatedAt: time.Now(),
	}

	// Insert the session
	err := f.SessionDb.Upsert(uow, session)
	if err != nil {
		t.Errorf("Expected no error on insert, got %v", err)
	}

	// Get the session
	retrieved, err := f.SessionDb.GetBySessionID("test-id")
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
	retrieved, err := f.SessionDb.GetBySessionID("nonexistent")

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
		ID:        "new-session",
		UserID:    "user-123",
		SessionID: "session-456",
		CreatedAt: time.Now(),
	}

	// Insert the session
	err := f.SessionDb.Upsert(uow, session)
	if err != nil {
		t.Errorf("Expected no error on insert, got %v", err)
	}

	// Verify it exists
	retrieved, err := f.SessionDb.GetBySessionID("new-session")
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
		ID:        "update-session",
		UserID:    "user-123",
		SessionID: "session-456",
		CreatedAt: time.Now(),
	}

	// Insert the session
	err := f.SessionDb.Upsert(uow, session)
	if err != nil {
		t.Errorf("Expected no error on insert, got %v", err)
	}

	// Update the session
	updatedSession := userSession.UserSession{
		ID:        "update-session",
		UserID:    "user-123",
		SessionID: "session-456",
		CreatedAt: session.CreatedAt,
		EndedAt:   time.Now(),
	}

	// Update the session
	err = f.SessionDb.Upsert(uow, updatedSession)
	if err != nil {
		t.Errorf("Expected no error on update, got %v", err)
	}

	// Verify it was updated
	retrieved, err := f.SessionDb.GetBySessionID("update-session")
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
