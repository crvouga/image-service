package userAccountDb

import (
	"testing"
	"time"

	"imageresizerservice/app/users/userAccount"
	"imageresizerservice/app/users/userID"
	"imageresizerservice/library/email/emailAddress"
	"imageresizerservice/library/keyValueDb"
	"imageresizerservice/library/sqlite"
	"imageresizerservice/library/uow"
)

type Fixture struct {
	UowFactory uow.UowFactory
	Db         UserAccountDb
}

func newFixture() *Fixture {
	db := sqlite.New()

	return &Fixture{
		Db:         ImplKeyValueDb{Db: &keyValueDb.ImplHashMap{}},
		UowFactory: uow.UowFactory{Db: db},
	}
}

func Test_GetByID(t *testing.T) {
	f := newFixture()
	uow, _ := f.UowFactory.Begin()

	// Create a session
	account := userAccount.UserAccount{
		ID:           userID.Gen(),
		EmailAddress: emailAddress.NewElsePanic("test@test.com"),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// Insert the session
	err := f.Db.Upsert(uow, account)
	if err != nil {
		t.Errorf("Expected no error on insert, got %v", err)
	}

	// Get the session
	retrieved, err := f.Db.GetByUserID(account.ID)
	if err != nil {
		t.Errorf("Expected no error on retrieval, got %v", err)
	}

	if retrieved == nil {
		t.Fatal("Expected to retrieve session, got nil")
	}

	if retrieved.ID != account.ID {
		t.Errorf("Expected ID to be %s, got %s", account.ID, retrieved.ID)
	}

	if retrieved.EmailAddress != account.EmailAddress {
		t.Errorf("Expected EmailAddress to be %s, got %s", account.EmailAddress, retrieved.EmailAddress)
	}

	uow.Commit()
}

func Test_GetByIDNonExistent(t *testing.T) {
	f := newFixture()

	// Try to get a session that doesn't exist
	retrieved, err := f.Db.GetByUserID("nonexistent")

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
	account := userAccount.UserAccount{
		ID:           userID.Gen(),
		EmailAddress: emailAddress.NewElsePanic("test@test.com"),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// Insert the session
	err := f.Db.Upsert(uow, account)
	if err != nil {
		t.Errorf("Expected no error on insert, got %v", err)
	}

	// Verify it exists
	retrieved, err := f.Db.GetByUserID(account.ID)
	if err != nil {
		t.Errorf("Expected no error on retrieval, got %v", err)
	}

	if retrieved == nil {
		t.Fatal("Expected to retrieve session, got nil")
	}

	if retrieved.ID != account.ID {
		t.Errorf("Expected ID to be %s, got %s", account.ID, retrieved.ID)
	}

	uow.Commit()
}

func Test_UpsertUpdateSession(t *testing.T) {
	f := newFixture()
	uow, _ := f.UowFactory.Begin()

	// Create initial session
	account := userAccount.UserAccount{
		ID:           userID.Gen(),
		EmailAddress: emailAddress.NewElsePanic("test@test.com"),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// Insert the session
	err := f.Db.Upsert(uow, account)
	if err != nil {
		t.Errorf("Expected no error on insert, got %v", err)
	}

	// Update the session
	updatedAccount := userAccount.UserAccount{
		ID:           "update-session",
		EmailAddress: "test@test.com",
		CreatedAt:    account.CreatedAt,
		UpdatedAt:    time.Now(),
	}

	// Update the session
	err = f.Db.Upsert(uow, updatedAccount)
	if err != nil {
		t.Errorf("Expected no error on update, got %v", err)
	}

	// Verify it was updated
	retrieved, err := f.Db.GetByUserID("update-session")
	if err != nil {
		t.Errorf("Expected no error on retrieval, got %v", err)
	}

	if retrieved == nil {
		t.Fatal("Expected to retrieve session, got nil")
	}

	if retrieved.UpdatedAt.IsZero() {
		t.Error("Expected UpdatedAt to be set, but it's zero")
	}

	uow.Commit()
}

func Test_GetByEmailAddress(t *testing.T) {
	f := newFixture()
	uow, _ := f.UowFactory.Begin()

	// Create a session
	account := userAccount.UserAccount{
		ID:           userID.Gen(),
		EmailAddress: emailAddress.NewElsePanic("test@test.com"),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// Insert the session
	err := f.Db.Upsert(uow, account)
	if err != nil {
		t.Errorf("Expected no error on insert, got %v", err)
	}

	// Get the session
	retrieved, err := f.Db.GetByEmailAddress(emailAddress.NewElsePanic("test@test.com"))
	if err != nil {
		t.Errorf("Expected no error on retrieval, got %v", err)
	}

	if retrieved == nil {
		t.Fatal("Expected to retrieve session, got nil")
	}

	if retrieved.ID != account.ID {
		t.Errorf("Expected ID to be %s, got %s", account.ID, retrieved.ID)
	}

	uow.Commit()
}
