package userAccountDB

import (
	"testing"
	"time"

	"imageService/app/users/userAccount"
	"imageService/app/users/userAccount/userRole"
	"imageService/app/users/userID"
	"imageService/library/email/emailAddress"
	"imageService/library/keyValueDB"
	"imageService/library/sqlite"
	"imageService/library/uow"
)

type Fixture struct {
	UowFactory uow.UowFactory
	DB         UserAccountDB
}

func newFixture() *Fixture {
	db := sqlite.New()
	keyValueDB := &keyValueDB.ImplHashMap{}

	return &Fixture{
		DB:         NewImplKeyValueDB(keyValueDB),
		UowFactory: *uow.NewFactory(db),
	}
}

func Test_GetByUserID(t *testing.T) {
	f := newFixture()
	uow, _ := f.UowFactory.Begin()

	// Create a session
	account := userAccount.UserAccount{
		UserID:       userID.Gen(),
		EmailAddress: emailAddress.NewElsePanic("test@test.com"),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// Insert the session
	err := f.DB.Upsert(uow, account)
	if err != nil {
		t.Errorf("Expected no error on insert, got %v", err)
	}

	// Get the session
	retrieved, err := f.DB.GetByUserID(account.UserID)
	if err != nil {
		t.Errorf("Expected no error on retrieval, got %v", err)
	}

	if retrieved == nil {
		t.Fatal("Expected to retrieve session, got nil")
	}

	if retrieved.UserID != account.UserID {
		t.Errorf("Expected ID to be %s, got %s", account.UserID, retrieved.UserID)
	}

	if retrieved.EmailAddress != account.EmailAddress {
		t.Errorf("Expected EmailAddress to be %s, got %s", account.EmailAddress, retrieved.EmailAddress)
	}

	uow.Commit()
}

func Test_GetByIDNonExistent(t *testing.T) {
	f := newFixture()

	// Try to get a session that doesn't exist
	retrieved, err := f.DB.GetByUserID("nonexistent")

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
		UserID:       userID.Gen(),
		EmailAddress: emailAddress.NewElsePanic("test@test.com"),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// Insert the session
	err := f.DB.Upsert(uow, account)
	if err != nil {
		t.Errorf("Expected no error on insert, got %v", err)
	}

	// Verify it exists
	retrieved, err := f.DB.GetByUserID(account.UserID)
	if err != nil {
		t.Errorf("Expected no error on retrieval, got %v", err)
	}

	if retrieved == nil {
		t.Fatal("Expected to retrieve session, got nil")
	}

	if retrieved.UserID != account.UserID {
		t.Errorf("Expected ID to be %s, got %s", account.UserID, retrieved.UserID)
	}

	uow.Commit()
}

func Test_UpsertUpdateSession(t *testing.T) {
	f := newFixture()
	uow, _ := f.UowFactory.Begin()

	// Create initial session
	account := userAccount.UserAccount{
		UserID:       userID.Gen(),
		EmailAddress: emailAddress.NewElsePanic("test@test.com"),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// Insert the session
	err := f.DB.Upsert(uow, account)
	if err != nil {
		t.Errorf("Expected no error on insert, got %v", err)
	}

	// Update the session
	updatedAccount := userAccount.UserAccount{
		UserID:       "update-session",
		EmailAddress: "test@test.com",
		CreatedAt:    account.CreatedAt,
		UpdatedAt:    time.Now(),
	}

	// Update the session
	err = f.DB.Upsert(uow, updatedAccount)
	if err != nil {
		t.Errorf("Expected no error on update, got %v", err)
	}

	// Verify it was updated
	retrieved, err := f.DB.GetByUserID("update-session")
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
		UserID:       userID.Gen(),
		EmailAddress: emailAddress.NewElsePanic("test@test.com"),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// Insert the session
	err := f.DB.Upsert(uow, account)
	if err != nil {
		t.Errorf("Expected no error on insert, got %v", err)
	}

	// Get the session
	retrieved, err := f.DB.GetByEmailAddress(emailAddress.NewElsePanic("test@test.com"))
	if err != nil {
		t.Errorf("Expected no error on retrieval, got %v", err)
	}

	if retrieved == nil {
		t.Fatal("Expected to retrieve session, got nil")
	}

	if retrieved.UserID != account.UserID {
		t.Errorf("Expected ID to be %s, got %s", account.UserID, retrieved.UserID)
	}

	uow.Commit()
}

func Test_GetByRole(t *testing.T) {
	f := newFixture()
	uow, _ := f.UowFactory.Begin()

	// Create accounts with different roles
	account1 := userAccount.UserAccount{
		UserID:       userID.Gen(),
		EmailAddress: emailAddress.NewElsePanic("admin@test.com"),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Role:         userRole.Admin,
	}

	account2 := userAccount.UserAccount{
		UserID:       userID.Gen(),
		EmailAddress: emailAddress.NewElsePanic("standard@test.com"),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Role:         userRole.Standard,
	}

	account3 := userAccount.UserAccount{
		UserID:       userID.Gen(),
		EmailAddress: emailAddress.NewElsePanic("admin2@test.com"),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Role:         userRole.Admin,
	}

	// Insert the accounts
	err := f.DB.Upsert(uow, account1)
	if err != nil {
		t.Errorf("Expected no error on insert account1, got %v", err)
	}

	err = f.DB.Upsert(uow, account2)
	if err != nil {
		t.Errorf("Expected no error on insert account2, got %v", err)
	}

	err = f.DB.Upsert(uow, account3)
	if err != nil {
		t.Errorf("Expected no error on insert account3, got %v", err)
	}

	// Get accounts by role
	adminAccounts, err := f.DB.GetByRole(userRole.Admin)
	if err != nil {
		t.Errorf("Expected no error on retrieval, got %v", err)
	}

	if len(adminAccounts) != 2 {
		t.Errorf("Expected to retrieve 2 admin accounts, got %d", len(adminAccounts))
	}

	standardAccounts, err := f.DB.GetByRole(userRole.Standard)
	if err != nil {
		t.Errorf("Expected no error on retrieval, got %v", err)
	}

	if len(standardAccounts) != 1 {
		t.Errorf("Expected to retrieve 1 standard account, got %d", len(standardAccounts))
	}

	// Verify the accounts have the correct role
	for _, account := range adminAccounts {
		if account.Role != userRole.Admin {
			t.Errorf("Expected role to be %s, got %s", userRole.Admin, account.Role)
		}
	}

	for _, account := range standardAccounts {
		if account.Role != userRole.Standard {
			t.Errorf("Expected role to be %s, got %s", userRole.Standard, account.Role)
		}
	}

	uow.Commit()
}
